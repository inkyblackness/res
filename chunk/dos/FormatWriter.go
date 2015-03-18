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
	resourceIDs      []uint16
	chunkAddresses   map[uint16]*chunkAddress
}

// NewChunkConsumer creates a consumer which writes to a random access destination
// using the DOS format.
func NewChunkConsumer(dest io.WriteSeeker) chunk.Consumer {
	coder := serial.NewPositioningEncoder(dest)
	result := &formatWriter{coder: coder,
		resourceIDs:    nil,
		chunkAddresses: make(map[uint16]*chunkAddress)}

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
	writer.alignToBoundary()
	address := &chunkAddress{
		startOffset: writer.coder.CurPos(),
		chunkType:   byte(chunk.ChunkType()),
		contentType: byte(chunk.ContentType())}

	if chunk.ChunkType().HasDirectory() {
		blockCount := chunk.BlockCount()
		blockStart := uint32(2 + 4*blockCount + 4)

		writer.coder.CodeUint16(&blockCount)
		for blockIndex := uint16(0); blockIndex < blockCount; blockIndex++ {
			block := chunk.BlockData(blockIndex)
			writer.coder.CodeUint32(&blockStart)
			blockStart += uint32(len(block))
		}
		writer.coder.CodeUint32(&blockStart)
		address.uncompressedLength = writer.coder.CurPos() - address.startOffset
	}

	for blockIndex := uint16(0); blockIndex < chunk.BlockCount(); blockIndex++ {
		block := chunk.BlockData(blockIndex)
		writer.coder.CodeBytes(block)
		address.uncompressedLength += uint32(len(block))
	}
	address.chunkLength = writer.coder.CurPos() - address.startOffset

	writer.resourceIDs = append(writer.resourceIDs, uint16(id))
	writer.chunkAddresses[uint16(id)] = address
}

// Finish marks the end of consumption. After calling Finish, the consumer can't be used anymore.
func (writer *formatWriter) Finish() {
	writer.alignToBoundary()
	directoryStart := writer.coder.CurPos()

	writer.writeDirectoryOffset(directoryStart)
	writer.coder.SetCurPos(directoryStart)
	chunksWritten := uint16(len(writer.resourceIDs))
	writer.coder.CodeUint16(&chunksWritten)
	writer.coder.CodeUint32(&writer.firstChunkOffset)

	for _, resourceID := range writer.resourceIDs {
		address := writer.chunkAddresses[resourceID]
		writer.coder.CodeUint16(&resourceID)
		address.code(writer.coder)
	}
}
