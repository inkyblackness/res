package dos

import (
	"errors"
	"io"

	"github.com/inkyblackness/res/chunk"
)

type formatReader struct {
}

// NewChunkProvider returns a chunk provider reading from a random access reader
// over a DOS format resource file.
func NewChunkProvider(source io.ReadSeeker) (provider chunk.Provider, err error) {
	err = errors.New("source is nil")

	return
}
