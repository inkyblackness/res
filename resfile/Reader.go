package resfile

import (
	"bytes"
	"errors"
	"io"

	"encoding/binary"
	"fmt"
	"github.com/inkyblackness/res/resfile/compression"
	"github.com/inkyblackness/res/serial"
)

// Reader provides methods to extract resource data from a serialized form.
// Chunks may be accessed concurrently due to the nature of the underlying io.ReaderAt.
type Reader struct {
	source           io.ReaderAt
	firstChunkOffset uint32
	directory        []chunkDirectoryEntry
	keyedDirectory   map[uint16]*chunkDirectoryEntry
}

var errSourceNil = errors.New("decoder is nil")
var errFormatMismatch = errors.New("format mismatch")

// ReaderFrom accesses the provided source and creates a new Reader instance
// from it.
// Should the provided decoder not follow the resource file format, an error
// is returned.
func ReaderFrom(source io.ReaderAt) (reader *Reader, err error) {
	if source == nil {
		return nil, errSourceNil
	}

	var dirOffset uint32
	dirOffset, err = readAndVerifyHeader(io.NewSectionReader(source, 0, chunkDirectoryFileOffsetPos+4))
	if err != nil {
		return nil, err
	}
	firstChunkOffset, directory, err := readDirectoryAt(dirOffset, source)
	if err != nil {
		return nil, err
	}

	reader = &Reader{
		source:           source,
		firstChunkOffset: firstChunkOffset,
		directory:        directory,
		keyedDirectory:   make(map[uint16]*chunkDirectoryEntry)}
	for index := 0; index < len(directory); index++ {
		entry := &reader.directory[index]
		reader.keyedDirectory[entry.ID] = entry
	}

	return
}

func readAndVerifyHeader(source io.ReadSeeker) (dirOffset uint32, err error) {
	coder := serial.NewPositioningDecoder(source)
	data := make([]byte, chunkDirectoryFileOffsetPos)
	coder.Code(data)
	coder.Code(&dirOffset)

	expected := make([]byte, len(headerString)+1)
	for index, r := range headerString {
		expected[index] = byte(r)
	}
	expected[len(headerString)] = commentTerminator
	if !bytes.Equal(data[:len(expected)], expected) {
		return 0, errFormatMismatch
	}

	return dirOffset, coder.FirstError()
}

func readDirectoryAt(dirOffset uint32, source io.ReaderAt) (firstChunkOffset uint32, directory []chunkDirectoryEntry, err error) {
	var header chunkDirectoryHeader
	headerSize := int64(binary.Size(&header))
	{
		headerCoder := serial.NewDecoder(io.NewSectionReader(source, int64(dirOffset), headerSize))
		headerCoder.Code(&header)
		if headerCoder.FirstError() != nil {
			return 0, nil, headerCoder.FirstError()
		}
	}

	firstChunkOffset = header.FirstChunkOffset
	directory = make([]chunkDirectoryEntry, header.ChunkCount)
	if header.ChunkCount > 0 {
		listCoder := serial.NewDecoder(io.NewSectionReader(source, int64(dirOffset)+headerSize, int64(binary.Size(directory))))
		listCoder.Code(directory)
		err = listCoder.FirstError()
	}
	return
}

// IDs returns the chunk identifier available via this reader.
// The order in the slice is the same as in the underlying serialized form.
func (reader *Reader) IDs() []ChunkID {
	ids := make([]ChunkID, len(reader.directory))
	for index, entry := range reader.directory {
		ids[index] = ChunkID(entry.ID)
	}
	return ids
}

// Chunk returns a reader for the specified chunk.
// If the ID is not known, nil is returned.
func (reader *Reader) Chunk(id Identifier) *ChunkReader {
	chunkStartOffset, entry := reader.findEntry(id.Value())
	if entry == nil {
		return nil
	}
	chunkType := entry.chunkType()
	compressed := (chunkType & chunkTypeFlagCompressed) != 0
	fragmented := (chunkType & chunkTypeFlagFragmented) != 0
	contentType := ContentType(entry.contentType())

	if fragmented {
		return reader.newFragmentedChunkReader(entry, contentType, compressed, chunkStartOffset)
	}
	return reader.newSingleBlockChunkReader(entry, contentType, compressed, chunkStartOffset)
}

func (reader *Reader) newFragmentedChunkReader(entry *chunkDirectoryEntry,
	contentType ContentType, compressed bool, chunkStartOffset uint32) *ChunkReader {
	baseDecoder := serial.NewDecoder(io.NewSectionReader(reader.source, int64(chunkStartOffset), int64(entry.packedLength())))
	var blockCount uint16
	type blockInfo struct {
		start uint32
		size  uint32
	}

	// TODO: unfinished

	baseDecoder.Code(&blockCount)
	var firstBlockOffset uint32
	baseDecoder.Code(&firstBlockOffset)
	lastBlockEndOffset := firstBlockOffset
	blocks := make([]blockInfo, blockCount)
	for blockIndex := uint16(0); blockIndex < blockCount; blockIndex++ {
		var endOffset uint32
		baseDecoder.Code(&endOffset)
		blocks[blockIndex].start = lastBlockEndOffset
		blocks[blockIndex].size = endOffset - lastBlockEndOffset
		lastBlockEndOffset = endOffset
	}

	return &ChunkReader{
		fragmented:  true,
		blockCount:  int(blockCount),
		contentType: contentType,
		compressed:  compressed}
}

func (reader *Reader) newSingleBlockChunkReader(entry *chunkDirectoryEntry,
	contentType ContentType, compressed bool, chunkStartOffset uint32) *ChunkReader {
	blockReader := func(index int) (io.Reader, error) {
		if index != 0 {
			return nil, fmt.Errorf("block index wrong: %v/%v", index, 1)
		}
		chunkSize := entry.packedLength()
		var chunkSource io.Reader = io.NewSectionReader(reader.source, int64(chunkStartOffset), int64(entry.packedLength()))
		if compressed {
			chunkSize = entry.unpackedLength()
			chunkSource = compression.NewDecompressor(chunkSource)
		}
		return io.LimitReader(chunkSource, int64(chunkSize)), nil
	}

	return &ChunkReader{
		fragmented:  false,
		blockCount:  1,
		contentType: contentType,
		compressed:  compressed,
		blockReader: blockReader}
}

func (reader *Reader) findEntry(id uint16) (startOffset uint32, entry *chunkDirectoryEntry) {
	startOffset = reader.firstChunkOffset
	for index := 0; (index < len(reader.directory)) && (entry == nil); index++ {
		cur := &reader.directory[index]
		if cur.ID == id {
			entry = cur
		} else {
			startOffset += cur.packedLength()
			startOffset += boundarySize - (startOffset % boundarySize)
		}
	}
	return
}
