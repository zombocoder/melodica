package volume

// Function to adjust volume by scaling audio samples
func AdjustVolume(data []byte, volume float64) {
	for i := 0; i < len(data); i += 2 {
		sample := int16(data[i]) | int16(data[i+1])<<8
		adjusted := int16(float64(sample) * volume)
		data[i] = byte(adjusted)
		data[i+1] = byte(adjusted >> 8)
	}
}
