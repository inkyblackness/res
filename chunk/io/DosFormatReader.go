package dos

import (
	"io"

	"github.com/inkyblackness/res/chunk"
)

type formatReader struct {
}

// NewReader wraps a random access reader and returns a chunk provider
func NewReader(source io.ReadSeeker) (chunk.Provider, error) {
	return nil, nil
}
