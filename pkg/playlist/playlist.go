package playlist

import (
	"bufio"
	"os"
	"strings"
)

// Function to load playlist from a file
func LoadPlaylist(filename string) ([]string, error) {
	var playlist []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" { // Skip empty lines
			playlist = append(playlist, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return playlist, nil
}
