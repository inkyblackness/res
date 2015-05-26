package mem

// L16SoundData is a in-memory sound buffer.
type L16SoundData struct {
	sampleRate float32
	samples    []int16
}

// NewL16SoundData returns a new sound data instance with the given data.
func NewL16SoundData(sampleRate float32, samples []int16) *L16SoundData {
	data := &L16SoundData{
		sampleRate: sampleRate,
		samples:    samples}

	return data
}

// SampleRate returns the amount of samples for one second.
func (data *L16SoundData) SampleRate() float32 {
	return data.sampleRate
}

// Samples copies the internal buffer into the given destination buffer.
func (data *L16SoundData) Samples(dest []int16) {
	copy(dest, data.samples)
}
