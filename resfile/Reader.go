package resfile

import (
	"bytes"
	"errors"
	"io"

	"github.com/inkyblackness/res/serial"
)

// Reader provides methods to extract resource data from a serialized form.
// Chunks may be accessed only one at a time. The reader does not support
// reading from multiple chunks concurrently.
type Reader struct {
	decoder          *serial.PositioningDecoder
	firstChunkOffset uint32
	directory        []chunkDirectoryEntry
	keyedDirectory   map[uint16]*chunkDirectoryEntry
}

var errSourceNil = errors.New("decoder is nil")
var errFormatMismatch = errors.New("format mismatch")

// ReaderFrom accesses the provided decoder and creates a new Reader instance
// from it.
// Should the provided decoder not follow the resource file format, an error
// is returned.
func ReaderFrom(source io.ReadSeeker) (reader *Reader, err error) {
	if source == nil {
		return nil, errSourceNil
	}

	decoder := serial.NewPositioningDecoder(source)

	var dirOffset uint32
	dirOffset, err = readAndVerifyHeader(decoder)
	if err != nil {
		return nil, err
	}
	firstChunkOffset, directory := readDirectoryAt(dirOffset, decoder)
	if decoder.FirstError() != nil {
		return nil, decoder.FirstError()
	}

	reader = &Reader{
		decoder:          decoder,
		firstChunkOffset: firstChunkOffset,
		directory:        directory,
		keyedDirectory:   make(map[uint16]*chunkDirectoryEntry)}
	for index := 0; index < len(directory); index++ {
		entry := &reader.directory[index]
		reader.keyedDirectory[entry.ID] = entry
	}

	return
}

func readAndVerifyHeader(coder serial.Coder) (dirOffset uint32, err error) {
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

func readDirectoryAt(dirOffset uint32, coder serial.PositioningCoder) (firstChunkOffset uint32, directory []chunkDirectoryEntry) {
	coder.SetCurPos(dirOffset)
	var chunkCount uint16
	coder.Code(&chunkCount)
	coder.Code(&firstChunkOffset)
	directory = make([]chunkDirectoryEntry, chunkCount)
	coder.Code(directory)
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
func (reader *Reader) Chunk(id Identifier) ChunkReader {
	startOffset, entry := reader.findEntry(id.Value())
	if entry == nil {
		return nil
	}
	chunkType := entry.chunkType()
	compressed := (chunkType & chunkTypeFlagCompressed) != 0
	fragmented := (chunkType & chunkTypeFlagFragmented) != 0
	contentType := ContentType(entry.contentType())

	var chunkReader ChunkReader
	if fragmented {
		chunkReader = &fragmentedBlockChunkReader{
			contentType: contentType,
			compressed:  compressed}
	} else {
		reader.decoder.SetCurPos(startOffset)
		chunkReader = &singleBlockChunkReader{
			contentType: contentType,
			compressed:  compressed,
			source:      io.LimitReader(reader.decoder, int64(entry.packedLength()))}
	}

	return chunkReader
}

func (reader *Reader) findEntry(id uint16) (startOffset uint32, entry *chunkDirectoryEntry) {
	startOffset = reader.firstChunkOffset
	for index := 0; (index < len(reader.directory)) && (entry == nil); index++ {
		cur := &reader.directory[index]
		if cur.ID == id {
			entry = cur
		} else {
			startOffset += cur.packedLength()
		}
	}
	return
}
