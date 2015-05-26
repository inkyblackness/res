package audio

// SoundData wraps the basic information about a collection of sound samples.
// The format is mono with signed 16bit PCM coding.
type SoundData interface {
	// SampleRate returns the amount of samples per second.
	SampleRate() float32
	// SampleCount returns the number of samples available from this data.
	SampleCount() int
	// Samples copies the samples into the provided destination buffer
	Samples(dest []int16)
}
