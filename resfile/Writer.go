package resfile

import (
	"errors"
	"io"
	"math"

	"github.com/inkyblackness/res/resfile/compression"
	"github.com/inkyblackness/res/serial"
)

// Writer provides methods to write a new resource file from scratch.
type Writer struct {
	encoder *serial.PositioningEncoder

	firstChunkOffset        uint32
	currentChunkStartOffset uint32
	currentChunk            chunkWriter

	directory []*chunkDirectoryEntry
}

var errTargetNil = errors.New("target is nil")

// NewWriter returns a new Writer instance prepared to add chunks.
// To finalize the created file, call Finish().
//
// This function will write initial information to the target and will return
// an error if the writer did. In such a case, the returned writer instance
// will produce invalid results and the state of the target is undefined.
func NewWriter(target io.WriteSeeker) (*Writer, error) {
	if target == nil {
		return nil, errTargetNil
	}

	encoder := serial.NewPositioningEncoder(target)
	writer := &Writer{encoder: encoder}
	writer.writeHeader()
	writer.firstChunkOffset = writer.encoder.CurPos()

	return writer, writer.encoder.FirstError()
}

var errWriterFinished = errors.New("writer is finished")

// CreateChunk adds a new single-block chunk to the current resource file.
// This chunk is closed by creating another chunk, or by finishing the writer.
func (writer *Writer) CreateChunk(id Identifier, contentType ContentType, compressed bool) (*BlockWriter, error) {
	if writer.encoder == nil {
		return nil, errWriterFinished
	}

	writer.finishLastChunk()
	if writer.encoder.FirstError() != nil {
		return nil, writer.encoder.FirstError()
	}

	writer.currentChunkStartOffset = writer.encoder.CurPos()
	var targetWriter io.Writer = serial.NewEncoder(writer.encoder)
	targetFinisher := func() {}
	chunkType := byte(0x00)
	if compressed {
		compressor := compression.NewCompressor(targetWriter)
		chunkType |= chunkTypeFlagCompressed
		targetWriter = compressor
		targetFinisher = func() { compressor.Close() } // nolint: errcheck
	}
	entry := &chunkDirectoryEntry{id: id.Value()}
	entry.setContentType(byte(contentType))
	entry.setChunkType(chunkType)
	writer.directory = append(writer.directory, entry)
	blockWriter := &BlockWriter{target: targetWriter, finisher: targetFinisher}
	writer.currentChunk = blockWriter

	return blockWriter, nil
}

// CreateFragmentedChunk adds a new fragmented chunk to the current resource file.
// This chunk is closed by creating another chunk, or by finishing the writer.
func (writer *Writer) CreateFragmentedChunk(id Identifier, contentType ContentType, compressed bool) (*FragmentedChunkWriter, error) {
	if writer.encoder == nil {
		return nil, errWriterFinished
	}

	writer.finishLastChunk()
	if writer.encoder.FirstError() != nil {
		return nil, writer.encoder.FirstError()
	}

	writer.currentChunkStartOffset = writer.encoder.CurPos()
	chunkType := chunkTypeFlagFragmented
	if compressed {
		chunkType |= chunkTypeFlagCompressed
	}
	entry := &chunkDirectoryEntry{id: id.Value()}
	entry.setContentType(byte(contentType))
	entry.setChunkType(chunkType)
	writer.directory = append(writer.directory, entry)
	chunkWriter := &FragmentedChunkWriter{target: serial.NewPositioningEncoder(writer.encoder), compressed: compressed}
	writer.currentChunk = chunkWriter

	return chunkWriter, nil
}

// Finish finalizes the resource file. After calling this function, the
// writer becomes unusable.
func (writer *Writer) Finish() (err error) {
	if writer.encoder == nil {
		return errWriterFinished
	}

	writer.finishLastChunk()

	directoryOffset := writer.encoder.CurPos()
	writer.encoder.SetCurPos(chunkDirectoryFileOffsetPos)
	writer.encoder.Code(directoryOffset)
	writer.encoder.SetCurPos(directoryOffset)
	writer.encoder.Code(uint16(len(writer.directory)))
	writer.encoder.Code(writer.firstChunkOffset)
	for _, entry := range writer.directory {
		writer.encoder.Code(entry)
	}

	err = writer.encoder.FirstError()
	writer.encoder = nil

	return
}

func (writer *Writer) writeHeader() {
	header := make([]byte, chunkDirectoryFileOffsetPos)
	for index, r := range headerString {
		header[index] = byte(r)
	}
	header[len(headerString)] = commentTerminator
	writer.encoder.Code(header)
	writer.encoder.Code(uint32(math.MaxUint32))
}

func (writer *Writer) finishLastChunk() {
	if writer.currentChunk != nil {
		currentEntry := writer.directory[len(writer.directory)-1]
		currentEntry.setUnpackedLength(writer.currentChunk.finish())
		currentEntry.setPackedLength(writer.encoder.CurPos() - writer.currentChunkStartOffset)

		writer.currentChunkStartOffset = 0
		writer.currentChunk = nil
	}
	writer.alignToBoundary()
}

func (writer *Writer) alignToBoundary() {
	extraBytes := writer.encoder.CurPos() % boundarySize
	if extraBytes > 0 {
		padding := make([]byte, boundarySize-extraBytes)
		writer.encoder.Code(padding)
	}
}
