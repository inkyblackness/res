package resfile

import (
	"io"
	"math"

	"github.com/inkyblackness/res/resfile/compression"
	"github.com/inkyblackness/res/serial"
)

func Write(target io.WriteSeeker, source *Reader) error {
	if target == nil {
		return errTargetNil
	}
	if source == nil {
		return errSourceNil
	}
	encoder := serial.NewPositioningEncoder(target)

	writeHeader(encoder)
	directoryOffsetPos := encoder.CurPos()
	encoder.Code(uint32(math.MaxUint32))
	firstChunkStartOffset := encoder.CurPos()

	chunkIDs := source.IDs()
	chunkDirectory := make([]chunkDirectoryEntry, len(chunkIDs))
	for index, chunkID := range chunkIDs {
		directoryEntry := &chunkDirectory[index]
		chunk, _ := source.Chunk(chunkID)
		var unpackedLength uint32
		var chunkType byte
		if chunk.Compressed() {
			chunkType |= chunkTypeFlagCompressed
		}

		chunkStartOffset := encoder.CurPos()

		if chunk.Fragmented() {
			chunkType |= chunkTypeFlagFragmented
			blockCount := chunk.BlockCount()
			encoder.Code(uint16(blockCount))
			entryStartOffset := encoder.CurPos()
			encoder.Code(make([]byte, 4*(blockCount+1)))

			unpackedLength = encoder.CurPos() - chunkStartOffset
			markBlockStart := func() {
				curPos := encoder.CurPos()
				encoder.SetCurPos(entryStartOffset)
				encoder.Code(unpackedLength)
				encoder.SetCurPos(curPos)
				entryStartOffset += 4
			}

			var blockWriter io.Writer = encoder
			blockFinish := func() {}
			if chunk.Compressed() {
				compressor := compression.NewCompressor(blockWriter)
				blockWriter = compressor
				blockFinish = func() { compressor.Close() }
			}

			for blockIndex := 0; blockIndex < blockCount; blockIndex++ {
				blockReader, _ := chunk.Block(blockIndex)

				markBlockStart()
				copied, _ := io.Copy(blockWriter, blockReader)
				unpackedLength += uint32(copied)
			}
			markBlockStart()
			blockFinish()
		} else {
			blockReader, _ := chunk.Block(0)
			var blockWriter io.Writer = encoder
			blockFinish := func() {}
			if chunk.Compressed() {
				compressor := compression.NewCompressor(blockWriter)
				blockWriter = compressor
				blockFinish = func() { compressor.Close() }
			}

			copied, _ := io.Copy(blockWriter, blockReader)
			blockFinish()
			unpackedLength = uint32(copied)
		}
		chunkEndOffset := encoder.CurPos()
		remainder := (boundarySize - chunkEndOffset%boundarySize) % boundarySize
		if remainder > 0 {
			encoder.Code(make([]byte, remainder))
		}

		directoryEntry.ID = chunkID.Value()
		directoryEntry.setChunkType(chunkType)
		directoryEntry.setContentType(byte(chunk.ContentType()))
		directoryEntry.setUnpackedLength(unpackedLength)
		directoryEntry.setPackedLength(chunkEndOffset - chunkStartOffset)
	}

	{
		directoryStartOffset := encoder.CurPos()
		encoder.SetCurPos(directoryOffsetPos)
		encoder.Code(&directoryStartOffset)
		encoder.SetCurPos(directoryStartOffset)
		encoder.Code(uint16(len(chunkDirectory)))
		encoder.Code(firstChunkStartOffset)
		for _, entry := range chunkDirectory {
			encoder.Code(&entry)
		}
	}
	return nil
}

func writeHeader(encoder serial.Coder) {
	header := make([]byte, chunkDirectoryFileOffsetPos)
	for index, r := range headerString {
		header[index] = byte(r)
	}
	header[len(headerString)] = commentTerminator
	encoder.Code(header)
}
