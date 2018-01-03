package resfile

import (
	"math"
)

type chunkDirectoryEntry struct {
	id                         uint16
	unpackedLengthAndChunkType uint32
	packedLengthAndContentType uint32
}

/*
func maskBits(field uint32, bitOffset uint, bitCount int) uint32 {
	return (field >> bitOffset) & uint32(^(uint64(math.MaxUint64) << uint64(bitCount)))
}
*/

func setBits(field uint32, bitOffset uint, bitCount int, value uint32) uint32 {
	mask := uint32(^(uint64(math.MaxUint64) << uint64(bitCount)))
	return (field & (^mask << bitOffset)) | ((value & mask) << bitOffset)
}

func (entry *chunkDirectoryEntry) setUnpackedLength(value uint32) {
	entry.unpackedLengthAndChunkType = setBits(entry.unpackedLengthAndChunkType, 0, 24, value)
}

/*
func (entry *chunkDirectoryEntry) unpackedLength() uint32 {
	return maskBits(entry.unpackedLengthAndChunkType, 0, 24)
}
*/

func (entry *chunkDirectoryEntry) setChunkType(value byte) {
	entry.unpackedLengthAndChunkType = setBits(entry.unpackedLengthAndChunkType, 24, 8, uint32(value))
}

/*
func (entry *chunkDirectoryEntry) chunkType() byte {
	return byte(maskBits(entry.unpackedLengthAndChunkType, 24, 8))
}
*/

func (entry *chunkDirectoryEntry) setPackedLength(value uint32) {
	entry.packedLengthAndContentType = setBits(entry.packedLengthAndContentType, 0, 24, value)
}

/*
func (entry *chunkDirectoryEntry) packedLength() uint32 {
	return maskBits(entry.packedLengthAndContentType, 0, 24)
}
*/

func (entry *chunkDirectoryEntry) setContentType(value byte) {
	entry.packedLengthAndContentType = setBits(entry.packedLengthAndContentType, 24, 8, uint32(value))
}

/*
func (entry *chunkDirectoryEntry) contentType() byte {
	return byte(maskBits(entry.packedLengthAndContentType, 24, 8))
}
*/
