package resfile

import (
	"github.com/inkyblackness/res/serial"
)

// BlockWriter is for writing the data of a single block in a chunk.
type BlockWriter struct {
	target *serial.PositioningEncoder
}

func (writer *BlockWriter) finish() (length uint32) {
	length = writer.target.CurPos()
	writer.target = nil
	return
}

// Write stores the given data in the block and follows the Writer interface.
func (writer *BlockWriter) Write(data []byte) (written int, err error) {
	return writer.target.Write(data)
}
