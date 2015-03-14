package dos

import (
	"fmt"
	"io"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/serial"
)

type formatReader struct {
}

var errFormatMismatch = fmt.Errorf("Format mismatch")

// NewChunkProvider returns a chunk provider reading from a random access reader
// over a DOS format resource file.
func NewChunkProvider(source io.ReadSeeker) (provider chunk.Provider, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	if source == nil {
		panic(fmt.Errorf("source is nil"))
	}
	coder := serial.NewDecoder(source)

	skipAndVerifyHeaderString(coder)
	skipAndVerifyComment(coder)

	provider = &formatReader{}

	return
}

func (reader *formatReader) IDs() []res.ResourceID {
	return nil
}

// Provide implements the chunk.Provider interface
func (reader *formatReader) Provide(id res.ResourceID) chunk.BlockHolder {
	return nil
}

func skipAndVerifyHeaderString(coder serial.Coder) {
	headerStringBuffer := make([]byte, len(HeaderString))
	coder.CodeBytes(headerStringBuffer)
	if string(headerStringBuffer) != HeaderString {
		panic(errFormatMismatch)
	}
}

func skipAndVerifyComment(coder serial.Coder) {
	terminatorFound := false

	for remaining := ChunkDirectoryFileOffsetPos - len(HeaderString); remaining > 0; remaining-- {
		temp := byte(0x00)
		coder.CodeByte(&temp)
		if temp == CommentTerminator {
			terminatorFound = true
		}
	}
	if !terminatorFound {
		panic(errFormatMismatch)
	}
}
