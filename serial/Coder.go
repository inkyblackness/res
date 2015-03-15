package serial

// Coder represents an encoder/decoder for binary data
type Coder interface {
	// SetCurPos sets the current position in the data
	SetCurPos(offset uint32)
	// CodeUint16 serializes an unsigned 16bit integer value
	CodeUint16(value *uint16)
	// CodeUint32 serializes an unsigned 32bit integer value
	CodeUint32(value *uint32)
	// CodeBytes serializes a slice
	CodeBytes(value []byte)
	// CodeByte serializes a single byte
	CodeByte(value *byte)
}
