package audio

import (
	"bytes"
	"io"
	"net/http"
)

// Function to download and buffer the audio data for controlled playback
func BufferAudioData(url string) (*bytes.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return nil, err
	}
	return bytes.NewReader(buf.Bytes()), nil
}
