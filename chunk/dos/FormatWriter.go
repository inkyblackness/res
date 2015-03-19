package dos

import (
	"io"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/compress/base"
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
	blockCoder := serial.Coder(writer.coder)
	chunkFinish := func() {}

	if chunk.ChunkType().HasDirectory() {
		writer.writeBlockDirectory(address, chunk, writer.getDirectoryPadding(id))
	}
	if chunk.ChunkType().IsCompressed() {
		compressor := base.NewCompressor(writer.coder)
		chunkFinish = func() { compressor.Close() }
		blockCoder = serial.NewEncoder(compressor)
	}

	for blockIndex := uint16(0); blockIndex < chunk.BlockCount(); blockIndex++ {
		block := chunk.BlockData(blockIndex)
		blockCoder.CodeBytes(block)
		address.uncompressedLength += uint32(len(block))
	}
	chunkFinish()
	address.chunkLength = writer.coder.CurPos() - address.startOffset

	writer.resourceIDs = append(writer.resourceIDs, uint16(id))
	writer.chunkAddresses[uint16(id)] = address
}

func (writer *formatWriter) getDirectoryPadding(id res.ResourceID) (padding uint32) {
	// Some directories have a 2byte padding before the actual data
	if id >= 0x08FC && id <= 0x094B { // all chunks in obj3d.res
		padding = uint32(2)
	}
	return
}

func (writer *formatWriter) writeBlockDirectory(address *chunkAddress, chunk chunk.BlockHolder, padding uint32) {
	blockCount := chunk.BlockCount()
	blockStart := uint32(2+4*blockCount+4) + padding

	writer.coder.CodeUint16(&blockCount)
	for blockIndex := uint16(0); blockIndex < blockCount; blockIndex++ {
		block := chunk.BlockData(blockIndex)
		writer.coder.CodeUint32(&blockStart)
		blockStart += uint32(len(block))
	}
	writer.coder.CodeUint32(&blockStart)
	for i := uint32(0); i < padding; i++ {
		zero := byte(0x00)
		writer.coder.CodeByte(&zero)
	}
	address.uncompressedLength = writer.coder.CurPos() - address.startOffset
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
