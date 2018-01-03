package resfile

import "github.com/inkyblackness/res/serial"

// FragmentedChunkWriter writes a chunk with zero, one, or more blocks.
type FragmentedChunkWriter struct {
	target *serial.PositioningEncoder

	blockStores []*serial.ByteStore
	blockWriter []*BlockWriter
}

func (writer *FragmentedChunkWriter) finish() (length uint32) {
	blockCount := len(writer.blockStores)
	writer.target.Code(uint16(blockCount))
	offset := 2 + (blockCount+1)*4
	for _, store := range writer.blockStores {
		writer.target.Code(uint32(offset))
		offset += len(store.Data())
	}
	writer.target.Code(uint32(offset))

	for _, store := range writer.blockStores {
		writer.target.Code(store.Data())
	}

	return writer.target.CurPos()
}

// CreateBlock provides a new, dedicated writer for a new block.
func (writer *FragmentedChunkWriter) CreateBlock() *BlockWriter {
	store := serial.NewByteStore()
	blockWriter := &BlockWriter{serial.NewPositioningEncoder(store)}
	writer.blockStores = append(writer.blockStores, store)
	writer.blockWriter = append(writer.blockWriter, blockWriter)
	return blockWriter
}
