package playlist_test

import (
	"os"
	"testing"

	"github.com/zombocoder/melodica/pkg/playlist"
)

// TestLoadPlaylist tests the LoadPlaylist function from the playlist package.
// It tests the function by creating a temporary file with sample playlist data
// and then loading the playlist from the file. The test verifies that the
// playlist is correctly loaded, trimmed of whitespace, and empty lines are
// skipped.
func TestLoadPlaylist(t *testing.T) {
	// Set up temporary test file with sample playlist data
	tempFile, err := os.CreateTemp("", "playlist_test.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up the file afterward

	// Write sample playlist data to the temp file
	sampleData := "song1.mp3\nsong2.mp3\n  song3.mp3  \n\nsong4.mp3\n"
	if _, err := tempFile.WriteString(sampleData); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Expected output after trimming whitespace and skipping empty lines
	expectedPlaylist := []string{"song1.mp3", "song2.mp3", "song3.mp3", "song4.mp3"}

	// Call LoadPlaylist with the temp file
	playlist, err := playlist.LoadPlaylist(tempFile.Name())
	if err != nil {
		t.Fatalf("LoadPlaylist() returned an error: %v", err)
	}

	// Verify the playlist matches the expected output
	if len(playlist) != len(expectedPlaylist) {
		t.Errorf("LoadPlaylist() returned %d items, expected %d", len(playlist), len(expectedPlaylist))
		return
	}

	// Compare each item in the playlist
	for i, song := range playlist {
		if song != expectedPlaylist[i] {
			t.Errorf("LoadPlaylist() returned item %d = %q, expected %q", i, song, expectedPlaylist[i])
		}
	}
}
