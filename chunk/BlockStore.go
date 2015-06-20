package chunk

import (
	"github.com/inkyblackness/res"
)

// BlockStore is a random access block container.
type BlockStore interface {
	// Type returns the type of the chunk.
	ChunkType() TypeID

	// ContentType returns the type of the data.
	ContentType() res.DataTypeID

	// BlockCount returns the number of blocks available in the chunk.
	// Flat chunks must contain exactly one block.
	BlockCount() uint16

	// Get returns the data for the requested block index.
	Get(block uint16) []byte

	// Put sets the data for the requested block index.
	Put(block uint16, data []byte)
}
