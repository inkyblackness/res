package dos

import (
	"io"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/serial"
)

type formatWriter struct {
	coder serial.PositioningCoder

	firstChunkOffset uint32
	chunksWritten    uint16
}

// NewChunkConsumer creates a consumer which writes to a random access destination
// using the DOS format.
func NewChunkConsumer(dest io.WriteSeeker) chunk.Consumer {
	coder := serial.NewPositioningEncoder(dest)
	result := &formatWriter{coder: coder, chunksWritten: 0}

	codeHeader(coder)
	result.writeDirectoryOffset(0xFFFFFFFF)
	result.firstChunkOffset = coder.CurPos()

	return result
}

func codeHeader(coder serial.PositioningCoder) {
	var blank byte = 0x00
	commentTerminator := CommentTerminator

	coder.CodeBytes([]byte(HeaderString))
	coder.CodeByte(&commentTerminator)
	for coder.CurPos() < ChunkDirectoryFileOffsetPos {
		coder.CodeByte(&blank)
	}
}

func (writer *formatWriter) writeDirectoryOffset(offset uint32) {
	writer.coder.SetCurPos(ChunkDirectoryFileOffsetPos)
	writer.coder.CodeUint32(&offset)
}

func (writer *formatWriter) alignToBoundary() {
	blank := byte(0)

	for writer.coder.CurPos()%BoundarySize != 0 {
		writer.coder.CodeByte(&blank)
	}
}

// Consume adds the given chunk to the consumer.
func (writer *formatWriter) Consume(id res.ResourceID, chunk chunk.BlockHolder) {

}

// Finish marks the end of consumption. After calling Finish, the consumer can't be used anymore.
func (writer *formatWriter) Finish() {
	writer.alignToBoundary()

	directoryStart := writer.coder.CurPos()
	writer.writeDirectoryOffset(directoryStart)
	writer.coder.SetCurPos(directoryStart)
	writer.coder.CodeUint16(&writer.chunksWritten)
	writer.coder.CodeUint32(&writer.firstChunkOffset)
}
