package serial

import "io"

// Encoder implements the Coder interface to write to a new byte array
type Encoder struct {
	offset uint32
	dest   io.WriteSeeker
}

// NewEncoder creates and returns a fresh encoder
func NewEncoder(dest io.WriteSeeker) *Encoder {
	encoder := &Encoder{
		offset: 0,
		dest:   dest}

	return encoder
}

// CurPos gets the current position in the data
func (coder *Encoder) CurPos() uint32 {
	return coder.offset
}

// SetCurPos sets the current position in the data
func (coder *Encoder) SetCurPos(offset uint32) {
	coder.dest.Seek(int64(offset), 0)
}

// CodeByte encodes a single byte
func (coder *Encoder) CodeByte(value *byte) {
	coder.writeBytes(*value)
}

// CodeBytes encodes the provided bytes
func (coder *Encoder) CodeBytes(value []byte) {
	coder.writeBytes(value...)
}

// CodeUint16 encodes an unsigned 16bit value
func (coder *Encoder) CodeUint16(value *uint16) {
	coder.writeBytes(byte((*value>>0)&0xFF), byte((*value>>8)&0xFF))
}

// CodeUint32 encodes an unsigned 32bit value
func (coder *Encoder) CodeUint32(value *uint32) {
	coder.writeBytes(byte((*value>>0)&0xFF), byte((*value>>8)&0xFF), byte((*value>>16)&0xFF), byte((*value>>24)&0xFF))
}

func (coder *Encoder) writeBytes(bytes ...byte) {
	written, err := coder.dest.Write(bytes)
	coder.offset += uint32(written)
	if err != nil {
		panic(err)
	}
}
