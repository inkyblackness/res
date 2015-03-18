package dos

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/serial"
)

type blockReader struct {
	coder   serial.PositioningCoder
	address *chunkAddress

	blocks [][]byte
}

// Type returns the type of the chunk.
func (reader *blockReader) ChunkType() chunk.TypeID {
	return chunk.TypeID(reader.address.chunkType)
}

// ContentType returns the type of the data.
func (reader *blockReader) ContentType() res.DataTypeID {
	return res.DataTypeID(reader.address.contentType)
}

// BlockCount returns the number of blocks available in the chunk.
// Flat chunks must contain exactly one block.
func (reader *blockReader) BlockCount() uint16 {
	reader.ensureBlocksBuffered()

	return uint16(len(reader.blocks))
}

// BlockData returns the data for the requested block index.
func (reader *blockReader) BlockData(block uint16) []byte {
	reader.ensureBlocksBuffered()

	return reader.blocks[block]
}

func (reader *blockReader) ensureBlocksBuffered() {
	if reader.blocks == nil {
		//blockOffsets := []uint32{0}

		reader.coder.SetCurPos(reader.address.startOffset)
		/*
			if ChunkType.HasDirectory() {

			}
		*/
		//reader.coder.SetCurPos(reader.address.startOffset + blockOffsets[0])
		data := make([]byte, reader.address.uncompressedLength)
		reader.coder.CodeBytes(data)
		reader.blocks = [][]byte{data}
	}
}
