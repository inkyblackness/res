package resfile

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReaderFromReturnsErrorForNilSource(t *testing.T) {
	reader, err := ReaderFrom(nil)

	assert.Nil(t, reader, "reader should be nil")
	assert.Equal(t, errSourceNil, err)
}

func TestReaderFromReturnsInstanceOnEmptySource(t *testing.T) {
	source := bytes.NewReader(emptyResourceFile())
	reader, err := ReaderFrom(source)

	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, reader)
}

func TestReaderFromReturnsErrorOnInvalidHeaderString(t *testing.T) {
	sourceData := emptyResourceFile()
	sourceData[10] = byte("A"[0])

	_, err := ReaderFrom(bytes.NewReader(sourceData))

	assert.Equal(t, errFormatMismatch, err)
}

func TestReaderFromReturnsErrorOnMissingCommentTerminator(t *testing.T) {
	sourceData := emptyResourceFile()
	sourceData[len(headerString)] = byte(0)

	_, err := ReaderFrom(bytes.NewReader(sourceData))

	assert.Equal(t, errFormatMismatch, err)
}

func TestReaderFromReturnsErrorOnInvalidDirectoryStart(t *testing.T) {
	sourceData := emptyResourceFile()
	sourceData[chunkDirectoryFileOffsetPos] = byte(0xFF)

	_, err := ReaderFrom(bytes.NewReader(sourceData))

	assert.NotNil(t, err)
}

func TestReaderFromCanDecodeExampleResourceFile(t *testing.T) {
	_, err := ReaderFrom(bytes.NewReader(exampleResourceFile()))

	assert.Nil(t, err, "no error expected")
}

func TestReaderIDsReturnsTheStoredChunkIDsInOrderFromFile(t *testing.T) {
	reader, _ := ReaderFrom(bytes.NewReader(exampleResourceFile()))

	assert.Equal(t, []ChunkID{exampleChunkIDSingleBlockChunk, exampleChunkIDSingleBlockChunkCompressed,
		exampleChunkIDFragmentedChunk, exampleChunkIDFragmentedChunkCompressed}, reader.IDs())
}

func TestReaderChunkReturnsNilForUnknownID(t *testing.T) {
	reader, _ := ReaderFrom(bytes.NewReader(emptyResourceFile()))
	chunkReader := reader.Chunk(ChunkID(0x1111))
	assert.Nil(t, chunkReader)
}

func TestReaderChunkReturnsAChunkReaderForKnownID(t *testing.T) {
	reader, _ := ReaderFrom(bytes.NewReader(exampleResourceFile()))
	chunkReader := reader.Chunk(exampleChunkIDSingleBlockChunk)
	assert.NotNil(t, chunkReader)
}
