package chunk

import "io"

// Chunk provides meta information as well as reader access to its contained blocks.
type Chunk struct {
	Fragmented    bool
	ContentType   ContentType
	Compressed    bool
	BlockProvider BlockProvider
}

// BlockCount returns the number of available blocks in the chunk.
// Unfragmented chunks will always have exactly one block.
func (chunk *Chunk) BlockCount() int {
	return chunk.BlockProvider.BlockCount()
}

// Block returns the reader for the identified block.
// Each call returns a new reader instance.
// Data provided by this reader is always uncompressed.
func (chunk *Chunk) Block(index int) (io.Reader, error) {
	return chunk.BlockProvider.Block(index)
}
