package dos

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/serial"
)

type blockReader struct {
	coder serial.Coder

	address chunkAddress
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
	return 0xFFFF
}

// BlockData returns the data for the requested block index.
func (reader *blockReader) BlockData(block uint16) []byte {
	return nil
}
