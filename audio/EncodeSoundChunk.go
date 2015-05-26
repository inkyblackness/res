package audio

import (
	"bytes"

	"github.com/inkyblackness/res/audio/voc"
)

// EncodeSoundChunk encodes the provided sound data into a byte array for a chunk.
func EncodeSoundChunk(soundData SoundData) []byte {
	writer := bytes.NewBuffer(nil)

	samples := make([]int16, soundData.SampleCount())
	soundData.Samples(samples)
	voc.Save(writer, soundData.SampleRate(), samples)

	return writer.Bytes()
}
