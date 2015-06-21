package store

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
)

type blockRetriever func() []byte

type backedBlockStore struct {
	holder     chunk.BlockHolder
	onModified ModifiedCallback
	retriever  []blockRetriever
}

func newBackedBlockStore(holder chunk.BlockHolder, onModified ModifiedCallback) *backedBlockStore {
	blockCount := holder.BlockCount()
	backed := &backedBlockStore{
		holder:     holder,
		onModified: onModified,
		retriever:  make([]blockRetriever, int(blockCount))}

	for index := range backed.retriever {
		backed.retriever[index] = func() []byte { return holder.BlockData(uint16(index)) }
	}

	return backed
}

// Type returns the type of the chunk.
func (backed *backedBlockStore) ChunkType() chunk.TypeID {
	return backed.holder.ChunkType()
}

// ContentType returns the type of the data.
func (backed *backedBlockStore) ContentType() res.DataTypeID {
	return backed.holder.ContentType()
}

// BlockCount returns the number of blocks available in the chunk.
// Flat chunks must contain exactly one block.
func (backed *backedBlockStore) BlockCount() uint16 {
	return backed.holder.BlockCount()
}

// Get returns the data for the requested block index.
func (backed *backedBlockStore) Get(block uint16) []byte {
	return backed.retriever[int(block)]()
}

// Put sets the data for the requested block index.
func (backed *backedBlockStore) Put(block uint16, data []byte) {
	backed.retriever[int(block)] = func() []byte { return data }
	backed.onModified()
}