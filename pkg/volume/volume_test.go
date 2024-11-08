package volume_test

import (
	"bytes"
	"testing"

	"github.com/zombocoder/melodica/pkg/volume"
)

// TestAdjustVolume tests the AdjustVolume function from the volume package.
// It tests the function with different volume levels to ensure that the audio
// samples are correctly scaled based on the volume level.
func TestAdjustVolume(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		volume   float64
		expected []byte
	}{
		{
			name:     "Volume at 1.0 (no change)",
			data:     []byte{0x01, 0x00, 0x02, 0x00}, // Two 16-bit samples: 1, 2
			volume:   1.0,
			expected: []byte{0x01, 0x00, 0x02, 0x00},
		},
		{
			name:     "Volume at 0.5 (reduce volume by half)",
			data:     []byte{0x04, 0x00, 0x08, 0x00}, // Two 16-bit samples: 4, 8
			volume:   0.5,
			expected: []byte{0x02, 0x00, 0x04, 0x00}, // Expected half: 2, 4
		},
		{
			name:     "Volume at 2.0 (double the volume)",
			data:     []byte{0x01, 0x00, 0x02, 0x00}, // Two 16-bit samples: 1, 2
			volume:   2.0,
			expected: []byte{0x02, 0x00, 0x04, 0x00}, // Expected double: 2, 4
		},
		{
			name:     "Volume at 0.0 (mute)",
			data:     []byte{0x01, 0x00, 0x02, 0x00}, // Two 16-bit samples: 1, 2
			volume:   0.0,
			expected: []byte{0x00, 0x00, 0x00, 0x00}, // Expected mute: 0, 0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]byte, len(tt.data))
			copy(data, tt.data) // Make a copy to avoid modifying the original test data
			volume.AdjustVolume(data, tt.volume)

			if !bytes.Equal(data, tt.expected) {
				t.Errorf("AdjustVolume() = %v, expected %v", data, tt.expected)
			}
		})
	}
}
