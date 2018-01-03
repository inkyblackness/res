package serial

import (
	"encoding/binary"
	"io"
)

// decoder is for decoding from a reader.
type decoder struct {
	source     io.Reader
	firstError error
	offset     uint32
}

// NewDecoder creates a new decoder from given source.
func NewDecoder(source io.Reader) Coder {
	return &decoder{source: source}
}

func (coder *decoder) FirstError() error {
	return coder.firstError
}

func (coder *decoder) Code(value interface{}) {
	if coder.firstError != nil {
		return
	}
	coder.firstError = binary.Read(coder.source, binary.LittleEndian, value)
	if coder.firstError != nil {
		return
	}
	coder.offset += uint32(binary.Size(value))
}
