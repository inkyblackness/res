package dos

import (
	"io"

	"github.com/inkyblackness/res/chunk"
)

type formatReader struct {
}

// NewReader returns a chunk provider reading from a random access reader
// over a DOS format resource file.
func NewReader(source io.ReadSeeker) (chunk.Provider, error) {
	return nil, nil
}
