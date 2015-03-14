package serial

// Encoder implements the Coder interface to write to a new byte array
type Encoder struct {
	offset int
	data   []byte
}

// NewEncoder creates and returns a fresh encoder
func NewEncoder() *Encoder {
	encoder := &Encoder{
		offset: 0,
		data:   make([]byte, 0)}

	return encoder
}

// Data returns the current data buffer
func (coder *Encoder) Data() []byte {
	return coder.data
}

// Len returns the length of the current data buffer
func (coder *Encoder) Len() int {
	return len(coder.data)
}

// CodeString encodes the provided string
func (coder *Encoder) CodeString(value *string) {
	coder.CodeBytes([]byte(*value))
}

// CodeBytes encodes the provided bytes
func (coder *Encoder) CodeBytes(value []byte) {
	size := len(value)

	coder.ensureAvailable(size)
	copy(coder.data[coder.offset:coder.offset+size], value)
	coder.offset += size
}

// CodeUint16 encodes an unsigned 16bit value
func (coder *Encoder) CodeUint16(value *uint16) {
	coder.ensureAvailable(2)
	coder.writeByte(byte((*value >> 0) & 0xFF))
	coder.writeByte(byte((*value >> 8) & 0xFF))
}

// CodeUint32 encodes an unsigned 32bit value
func (coder *Encoder) CodeUint32(value *uint32) {
	coder.ensureAvailable(4)
	coder.writeByte(byte((*value >> 0) & 0xFF))
	coder.writeByte(byte((*value >> 8) & 0xFF))
	coder.writeByte(byte((*value >> 16) & 0xFF))
	coder.writeByte(byte((*value >> 24) & 0xFF))
}

// CodeByte encodes a single byte
func (coder *Encoder) CodeByte(value *byte) {
	coder.ensureAvailable(1)
	coder.writeByte(*value)
}

func (coder *Encoder) writeByte(value byte) {
	coder.data[coder.offset] = value
	coder.offset++
}

func (coder *Encoder) ensureAvailable(size int) {
	expected := coder.offset + size
	available := len(coder.data)

	if expected > available {
		old := coder.data

		coder.data = make([]byte, expected)
		copy(coder.data, old)
	}
}
