package resfile

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	sourceData[chunkDirectoryFileOffsetPos+1] = byte(0xFF)

	_, err := ReaderFrom(bytes.NewReader(sourceData))

	assert.NotNil(t, err, "error expected")
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

func TestReaderChunkReturnsChunkWithMetaInformation(t *testing.T) {
	reader, _ := ReaderFrom(bytes.NewReader(exampleResourceFile()))
	info := func(chunkID ChunkID, name string, expected interface{}) string {
		return fmt.Sprintf("Chunk 0x%04X should have %v = %v", chunkID.Value(), name, expected)
	}
	verifyChunk := func(chunkID ChunkID, fragmented bool, contentType ContentType, compressed bool) {
		chunkReader := reader.Chunk(chunkID)
		assert.Equal(t, fragmented, chunkReader.Fragmented(), info(chunkID, "fragmented", fragmented))
		assert.Equal(t, contentType, chunkReader.ContentType(), info(chunkID, "contentType", contentType))
		assert.Equal(t, compressed, chunkReader.Compressed(), info(chunkID, "compressed", compressed))
	}
	verifyChunk(exampleChunkIDSingleBlockChunk, false, ContentType(0x01), false)
	verifyChunk(exampleChunkIDSingleBlockChunkCompressed, false, ContentType(0x02), true)
	verifyChunk(exampleChunkIDFragmentedChunk, true, ContentType(0x03), false)
	verifyChunk(exampleChunkIDFragmentedChunkCompressed, true, ContentType(0x04), true)
}

func TestReaderChunkWithUncompressedSingleBlockContent(t *testing.T) {
	reader, _ := ReaderFrom(bytes.NewReader(exampleResourceFile()))
	chunkReader := reader.Chunk(exampleChunkIDSingleBlockChunk)

	assert.Equal(t, 1, chunkReader.BlockCount())
	verifyBlockContent(t, chunkReader.Block(0), []byte{0x01, 0x01, 0x01})
}

func TestReaderChunkWithCompressedSingleBlockContent(t *testing.T) {
	reader, _ := ReaderFrom(bytes.NewReader(exampleResourceFile()))
	chunkReader := reader.Chunk(exampleChunkIDSingleBlockChunkCompressed)

	assert.Equal(t, 1, chunkReader.BlockCount())
	verifyBlockContent(t, chunkReader.Block(0), []byte{0x02, 0x02})
}

func verifyBlockContent(t *testing.T, reader io.Reader, expected []byte) {
	require.NotNil(t, reader, "reader is nil")
	data, err := ioutil.ReadAll(reader)
	assert.Nil(t, err, "no error expected")
	assert.Equal(t, expected, data)
}
