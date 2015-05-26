package audio

import (
	"github.com/inkyblackness/res/audio/mem"

	check "gopkg.in/check.v1"
)

type CodeSoundChunkSuite struct {
}

var _ = check.Suite(&CodeSoundChunkSuite{})

func (suite *CodeSoundChunkSuite) TestChunkTypeReturnsProvidedValue(c *check.C) {
	rawSoundSamples := []int16{-0x8000, 0x0000, 0x2041, 0x4081, 0x7FFF}
	rawSound := mem.NewL16SoundData(20000.0, rawSoundSamples)
	encoded := EncodeSoundChunk(rawSound)
	decoded, err := DecodeSoundChunk(encoded)

	c.Assert(err, check.IsNil)

	c.Check(decoded.SampleRate(), check.Equals, rawSound.SampleRate())
	c.Check(decoded.SampleCount(), check.Equals, rawSound.SampleCount())
	samples := make([]int16, decoded.SampleCount())
	decoded.Samples(samples)
	c.Check(samples, check.DeepEquals, rawSoundSamples)
}
