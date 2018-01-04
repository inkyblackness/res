package resfile

import "io"

// ChunkReader provides meta information as well as reader access to its contained blocks.
type ChunkReader interface {
	// Fragmented describes how many blocks can be expected.
	// Unfragmented chunks have exactly one block, fragmented chunks zero, one, or more.
	Fragmented() bool

	// ContentType describes the nature of the data within the chunk - the format of the blocks.
	ContentType() ContentType

	// Compressed returns true if the data is to be serialized in compressed form
	// in the resource file. Data provided by this reader is already uncompressed.
	Compressed() bool

	// BlockCount returns the number of available blocks in this chunk.
	// Unfragmented chunks will always have exactly one block.
	BlockCount() int

	// Block returns the reader for the identified block.
	Block(index int) io.Reader
}
