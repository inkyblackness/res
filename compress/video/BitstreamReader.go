package video

// BitstreamReader is a utility to read big-endian integer values of arbitrary bit size from a bitstream.
type BitstreamReader struct {
	source          []byte
	nextSourceIndex int
	sourceBitLen    uint64

	currentBitPos uint64

	buffer       uint64
	bitsBuffered uint64
}

// NewBitsreamReader returns a new instance of a bistream reader using the provided byte slice as source.
func NewBitstreamReader(data []byte) *BitstreamReader {
	return &BitstreamReader{
		source:       data,
		sourceBitLen: uint64(len(data)) * 8}
}

func (reader *BitstreamReader) bufferNextByte() {
	reader.buffer = (reader.buffer << 8) | uint64(reader.source[reader.nextSourceIndex])
	reader.bitsBuffered += 8
	reader.nextSourceIndex++
}

// Read returns a value with the requested bit size, right aligned, as a uint32.
// Reading does not advance the current position. A successful read of a certain size will return the same
// value when called repeatedly with the same parameter.
// An error is returned if less data is available from the source than requested.
//
// The function panics when reading more than 32 bits.
func (reader *BitstreamReader) Read(bits int) (result uint32, err error) {
	if bits > 32 {
		panic("Limit of bit count: 32")
	}

	if (reader.currentBitPos + uint64(bits)) <= reader.sourceBitLen {
		for reader.bitsBuffered < uint64(bits) {
			reader.bufferNextByte()
		}
		result = uint32(reader.buffer >> (reader.bitsBuffered - uint64(bits)) & ^(uint64(0xFFFFFFFFFFFFFFFF) << uint64(bits)))
	} else {
		err = BitstreamEndError
	}

	return
}

// Advance skips the provided amount of bits and puts the current position there.
// Successive read operations will return values from the new position.
// An error is returned if advance would skip beyond the end.
//
// This function panics when advancing with negative values.
func (reader *BitstreamReader) Advance(bits int) (err error) {
	if bits < 0 {
		panic("Can only advance forward")
	}

	newIndex := reader.currentBitPos + uint64(bits)
	if newIndex <= reader.sourceBitLen {
		reader.currentBitPos = newIndex
		if uint64(bits) > reader.bitsBuffered {
			reader.bitsBuffered = 0
			reader.nextSourceIndex = int(reader.currentBitPos / 8)
			remainder := reader.currentBitPos % 8
			if remainder > 0 {
				reader.bufferNextByte()
				reader.bitsBuffered -= remainder
			}
		} else {
			reader.bitsBuffered -= uint64(bits)
		}
	} else {
		err = BitstreamEndError
	}

	return
}
