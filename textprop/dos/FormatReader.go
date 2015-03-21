package dos

import (
	"fmt"
	"io"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/serial"
	"github.com/inkyblackness/res/textprop"
)

type formatReader struct {
	coder        serial.PositioningCoder
	textureCount uint32
}

var errFormatMismatch = fmt.Errorf("Format mismatch")

// NewProvider wraps the provided ReadSeeker in a provider for texture properties.
func NewProvider(source io.ReadSeeker) (provider textprop.Provider, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	if source == nil {
		panic(fmt.Errorf("source is nil"))
	}
	count := readAndVerifyEntryCount(source)

	provider = &formatReader{coder: serial.NewPositioningDecoder(source), textureCount: count}

	return
}

func (provider *formatReader) TextureCount() uint32 {
	return provider.textureCount
}

func (provider *formatReader) Provide(id res.TextureID) []byte {
	data := make([]byte, textprop.TexturePropertiesLength)

	provider.coder.SetCurPos(textprop.TexturePropertiesLength * uint32(id))
	provider.coder.CodeBytes(data)

	return data
}

func readAndVerifyEntryCount(source io.Seeker) uint32 {
	sourceLength := getSeekerSize(source)
	count := sourceLength / textprop.TexturePropertiesLength

	if (count * textprop.TexturePropertiesLength) != sourceLength {
		panic(errFormatMismatch)
	}

	return count
}

func getSeekerSize(seeker io.Seeker) uint32 {
	length, err := seeker.Seek(0, 2)

	if err != nil {
		panic(err)
	}

	return uint32(length)
}
