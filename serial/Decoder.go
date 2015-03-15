package serial

import "io"

// Decoder is for decoding from a random access stream
type Decoder struct {
	source io.ReadSeeker
}

// NewDecoder creates a new decoder from given source
func NewDecoder(source io.ReadSeeker) *Decoder {
	coder := &Decoder{source: source}

	return coder
}

// SetCurPos sets the current position in the data
func (coder *Decoder) SetCurPos(offset uint32) {
	coder.source.Seek(int64(offset), 0)
}

// CodeByte decodes a single byte
func (coder *Decoder) CodeByte(value *byte) {
	buf := coder.readBytes(1)
	*value = buf[0]
}

// CodeBytes decodes the provided bytes
func (coder *Decoder) CodeBytes(value []byte) {
	_, err := coder.source.Read(value)
	if err != nil {
		panic(err)
	}
}

// CodeUint16 decodes an unsigned 16bit value
func (coder *Decoder) CodeUint16(value *uint16) {
	buf := coder.readBytes(2)
	*value = (uint16(buf[0]) << 0) | (uint16(buf[1]) << 8)
}

// CodeUint32 decodes a 32bit unsigned integer
func (coder *Decoder) CodeUint32(value *uint32) {
	buf := coder.readBytes(4)
	*value = (uint32(buf[0]) << 0) | (uint32(buf[1]) << 8) | (uint32(buf[2]) << 16) | (uint32(buf[3]) << 24)
}

func (coder *Decoder) readBytes(size int) []byte {
	buf := make([]byte, size)
	coder.CodeBytes(buf)

	return buf
}
