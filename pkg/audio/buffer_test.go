package audio_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zombocoder/melodica/pkg/audio"
)

// TestBufferAudioData tests that BufferAudioData correctly downloads and buffers audio data.
// It creates a test server that serves mock audio data and verifies that the buffered data
// matches the mock audio data.
func TestBufferAudioData(t *testing.T) {
	// Mock audio data to be served by the test server
	mockAudioData := []byte("This is mock audio data")

	// Create a test server that serves the mock audio data
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(mockAudioData)
	}))
	defer ts.Close()

	// Call BufferAudioData with the test server's URL
	bufferedData, err := audio.BufferAudioData(ts.URL)
	if err != nil {
		t.Fatalf("BufferAudioData() returned an error: %v", err)
	}

	// Read the buffered data to verify it matches the mock audio data
	result := make([]byte, len(mockAudioData))
	n, err := bufferedData.Read(result)
	if err != nil && err != io.EOF {
		t.Fatalf("Failed to read from buffered data: %v", err)
	}

	// Ensure the amount of data read matches the mock audio data length
	if n != len(mockAudioData) {
		t.Errorf("Read %d bytes, expected %d bytes", n, len(mockAudioData))
	}

	// Verify that the data read matches the mock audio data
	if !bytes.Equal(result, mockAudioData) {
		t.Errorf("Buffered data = %q, expected %q", result, mockAudioData)
	}
}
