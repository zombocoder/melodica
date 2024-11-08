package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/rivo/tview"
	"github.com/zombocoder/melodica/pkg/audio"
	"github.com/zombocoder/melodica/pkg/playlist"
	"github.com/zombocoder/melodica/pkg/volume"
)

// Initial volume level
var volumeLevel float64 = 1.0  // Ranges from 0.0 (mute) to 2.0 (double volume)
var paused bool                // Tracks if playback is paused
var pausedMutex sync.Mutex     // Protects access to paused variable
var currentPos int64           // Tracks the current position in the audio data
var spacePressed bool          // Debounce flag for Space key
const SampleRate = 44100       // Sample rate for audio playback
const ChannelNum = 2           // Number of channels (stereo),
const BitDepthInBytes = 2      // 16-bit audio
const BufferSizeInBytes = 4096 // Buffer size for audio playback
// Function to play buffered audio data with support for pause and resume
func playBufferedAudio(ctx context.Context, reader *bytes.Reader, otoCtx *oto.Context) error {
	decoder, err := mp3.NewDecoder(reader)
	if err != nil {
		return err
	}

	player := otoCtx.NewPlayer()
	if player == nil {
		return fmt.Errorf("failed to create audio player")
	}
	defer player.Close()

	buffer := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			currentPos, _ = reader.Seek(0, io.SeekCurrent) // Save current position
			return nil
		default:
			pausedMutex.Lock()
			if paused {
				pausedMutex.Unlock()
				time.Sleep(100 * time.Millisecond)
				continue
			}
			pausedMutex.Unlock()

			n, err := decoder.Read(buffer)
			if err == io.EOF {
				currentPos = 0 // Reset position for next playback
				return nil
			}
			if err != nil {
				return err
			}

			// Adjust volume
			volume.AdjustVolume(buffer[:n], volumeLevel)

			player.Write(buffer[:n])
		}
	}
}

// Function to play a specific track
func playTrack(app *tview.Application, playlist []string, otoCtx *oto.Context, playlistView *tview.List, nowPlayingView *tview.TextView, trackIndex int, playCancel *context.CancelFunc) {
	// Stop any existing playback
	if *playCancel != nil {
		(*playCancel)()
	}

	// Reset pause state
	pausedMutex.Lock()
	paused = false
	pausedMutex.Unlock()

	// Update the playlist view highlight to the current track
	playlistView.SetCurrentItem(trackIndex)

	// Start the specified track
	currentTrack := filepath.Base(playlist[trackIndex])
	nowPlayingView.SetText(fmt.Sprintf("Now Playing: %s | Volume: %.0f%%", currentTrack, volumeLevel*100))

	var ctx context.Context
	ctx, *playCancel = context.WithCancel(context.Background())
	go func() {
		audioReader, err := audio.BufferAudioData(playlist[trackIndex])
		if err != nil {
			log.Printf("could not buffer audio for %s: %v", currentTrack, err)
			return
		}
		audioReader.Seek(currentPos, io.SeekStart) // Resume from the last saved position
		if err := playBufferedAudio(ctx, audioReader, otoCtx); err != nil {
			log.Printf("could not play %s: %v", currentTrack, err)
		}
		// Automatically play the next track after the current one finishes
		if ctx.Err() == nil {
			app.QueueUpdateDraw(func() {
				nextIndex := (trackIndex + 1) % len(playlist) // Move to the next track, loop back if at the end
				currentPos = 0                                // Reset position for next track
				playTrack(app, playlist, otoCtx, playlistView, nowPlayingView, nextIndex, playCancel)
			})
		}
	}()
}

func main() {
	// Set up logging to a file
	logFile, err := os.OpenFile("melodica.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Check for playlist file argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: melodica <playlist.txt>")
		os.Exit(1)
	}
	playlistFile := os.Args[1]

	app := tview.NewApplication()

	// Load playlist from file provided in the argument
	playlist, err := playlist.LoadPlaylist(playlistFile)
	if err != nil {
		log.Fatalf("could not load playlist: %v", err)
	}

	// Setup audio context
	otoCtx, err := oto.NewContext(SampleRate, ChannelNum, BitDepthInBytes, BufferSizeInBytes)
	if err != nil {
		log.Fatalf("could not create audio context: %v", err)
	}
	defer otoCtx.Close()

	// Track the current playing song index and playback control
	currentIndex := 0
	var playCancel context.CancelFunc

	// Create the "Now Playing" panel below the playlist
	nowPlayingView := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Now Playing: None").
		SetChangedFunc(func() {
			app.Draw()
		})
	nowPlayingView.SetBorder(true).SetTitle("Now Playing")

	// Create the playlist view (top panel with 100% width)
	playlistView := tview.NewList().SetWrapAround(true)
	for i, url := range playlist {
		filename := filepath.Base(url)
		index := i // Capture the current index for each list item
		playlistView.AddItem(fmt.Sprintf("[%d] %s", i+1, filename), "", 0, func() {
			// Play selected track from the list
			currentIndex = index
			currentPos = 0 // Reset position when starting a new track
			playTrack(app, playlist, otoCtx, playlistView, nowPlayingView, currentIndex, &playCancel)
		})
	}
	playlistView.SetBorder(true).SetTitle("Playlist")

	// Create a vertical layout with the playlist at the top and "Now Playing" below
	verticalLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(playlistView, 0, 4, true).   // Playlist takes up most of the space
		AddItem(nowPlayingView, 3, 1, false) // Now Playing panel below the playlist

	// Bottom panel to display key bindings only
	bottomTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Enter: Play | n: Next | p: Previous | s: Stop | x: Pause/Resume | >: Volume Up | <: Volume Down | Esc: Quit").
		SetChangedFunc(func() {
			app.Draw()
		})
	bottomTextView.SetBorder(true)

	// Main layout with vertical top layout and bottom panel
	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(verticalLayout, 0, 1, true).
		AddItem(bottomTextView, 3, 1, false) // Bottom panel with controls info only

	// Set up global key bindings
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			app.Stop()
		case tcell.KeyEnter:
			// Play selected track
			currentIndex = playlistView.GetCurrentItem()
			currentPos = 0 // Reset position when starting a new track
			playTrack(app, playlist, otoCtx, playlistView, nowPlayingView, currentIndex, &playCancel)
		case tcell.KeyRune:
			switch event.Rune() {
			case 'x':
				if spacePressed {
					return event // Ignore if Space is already pressed
				}
				spacePressed = true                     // Set the debounce flag
				defer func() { spacePressed = false }() // Reset flag after function completes

				// Toggle pause/resume
				pausedMutex.Lock()
				paused = !paused
				pausedMutex.Unlock()
				if paused {
					nowPlayingView.SetText("Paused")
				} else {
					nowPlayingView.SetText(fmt.Sprintf("Now Playing: %s | Volume: %.0f%%", filepath.Base(playlist[currentIndex]), volumeLevel*100))
				}
			case 'n': // Next track
				currentIndex = (currentIndex + 1) % len(playlist) // Move to the next track, loop back if at the end
				currentPos = 0                                    // Reset position for next track
				playTrack(app, playlist, otoCtx, playlistView, nowPlayingView, currentIndex, &playCancel)
			case 'p': // Previous track
				currentIndex = (currentIndex - 1 + len(playlist)) % len(playlist) // Move to the previous track, loop back if at the start
				currentPos = 0                                                    // Reset position for previous track
				playTrack(app, playlist, otoCtx, playlistView, nowPlayingView, currentIndex, &playCancel)
			case 's': // Stop playback
				if playCancel != nil {
					playCancel()
					playCancel = nil
				}
				nowPlayingView.SetText("Now Playing: None")
			case '>': // Volume up
				volumeLevel = math.Min(volumeLevel+0.1, 2.0) // Cap volume at 200%
				nowPlayingView.SetText(fmt.Sprintf("Now Playing: %s | Volume: %.0f%%", filepath.Base(playlist[currentIndex]), volumeLevel*100))
			case '<': // Volume down
				volumeLevel = math.Max(volumeLevel-0.1, 0.0) // Minimum volume at 0%
				nowPlayingView.SetText(fmt.Sprintf("Now Playing: %s | Volume: %.0f%%", filepath.Base(playlist[currentIndex]), volumeLevel*100))
			}
		}
		return event
	})

	// Set the root and run the application
	if err := app.SetRoot(mainLayout, true).Run(); err != nil {
		panic(err)
	}
}
