package resfile

import (
	"testing"

	"github.com/inkyblackness/res/serial"

	"github.com/stretchr/testify/assert"
)

func TestNewWriterReturnsErrorForNilTarget(t *testing.T) {
	writer, err := NewWriter(nil)

	assert.Nil(t, writer, "writer should be nil")
	assert.Equal(t, errTargetNil, err)
}

func TestWriterFinishWithoutAddingChunksCreatesValidFileWithoutChunks(t *testing.T) {
	emptyFileData := emptyResourceFile()
	store := serial.NewByteStore()
	writer, err := NewWriter(store)
	assert.Nil(t, err, "no error expected creating writer")

	err = writer.Finish()
	assert.Nil(t, err, "no error expected finishing writer")
	assert.Equal(t, emptyFileData, store.Data())
}

func TestWriterFinishReturnsErrorWhenAlreadyFinished(t *testing.T) {
	writer, _ := NewWriter(serial.NewByteStore())

	writer.Finish()

	err := writer.Finish()
	assert.Equal(t, errWriterFinished, err)
}

func TestWriterUncompressedSingleBlockChunkCanBeWritten(t *testing.T) {
	data := []byte{0xAB, 0x01, 0xCD, 0x02, 0xEF}
	store := serial.NewByteStore()
	writer, _ := NewWriter(store)
	chunkWriter, err := writer.CreateChunk(ChunkID(0x1234), ContentType(0x0A), false)
	assert.Nil(t, err, "no error expected")
	chunkWriter.Write(data)
	writer.Finish()

	result := store.Data()

	var expected []byte
	expected = append(expected, data...)
	expected = append(expected, 0x00, 0x00, 0x00)       // alignment for directory
	expected = append(expected, 0x01, 0x00)             // chunk count
	expected = append(expected, 0x80, 0x00, 0x00, 0x00) // offset to first chunk
	expected = append(expected, 0x34, 0x12)             // chunk ID
	expected = append(expected, 0x05, 0x00, 0x00)       // chunk length (uncompressed)
	expected = append(expected, 0x00)                   // chunk type (uncompressed, single-block)
	expected = append(expected, 0x05, 0x00, 0x00)       // chunk length in file
	expected = append(expected, 0x0A)                   // content type
	assert.Equal(t, expected, result[chunkDirectoryFileOffsetPos+4:])
}

func TestWriterUncompressedFragmentedChunkCanBeWritten(t *testing.T) {
	blockData1 := []byte{0xAB, 0x01, 0xCD}
	blockData2 := []byte{0x11, 0x22, 0x33, 0x44}
	store := serial.NewByteStore()
	writer, _ := NewWriter(store)
	chunkWriter, err := writer.CreateFragmentedChunk(ChunkID(0x5678), ContentType(0x0B), false)
	assert.Nil(t, err, "no error expected")
	chunkWriter.CreateBlock().Write(blockData1)
	chunkWriter.CreateBlock().Write(blockData2)
	writer.Finish()

	result := store.Data()

	var expected []byte
	expected = append(expected, 0x02, 0x00)             // number of blocks
	expected = append(expected, 0x0E, 0x00, 0x00, 0x00) // offset to first block
	expected = append(expected, 0x11, 0x00, 0x00, 0x00) // offset to second block
	expected = append(expected, 0x15, 0x00, 0x00, 0x00) // size of chunk
	expected = append(expected, blockData1...)
	expected = append(expected, blockData2...)
	expected = append(expected, 0x00, 0x00, 0x00)       // alignment for directory
	expected = append(expected, 0x01, 0x00)             // chunk count
	expected = append(expected, 0x80, 0x00, 0x00, 0x00) // offset to first chunk
	expected = append(expected, 0x78, 0x56)             // chunk ID
	expected = append(expected, 0x15, 0x00, 0x00)       // chunk length (uncompressed)
	expected = append(expected, 0x02)                   // chunk type
	expected = append(expected, 0x15, 0x00, 0x00)       // chunk length in file
	expected = append(expected, 0x0B)                   // content type
	assert.Equal(t, expected, result[chunkDirectoryFileOffsetPos+4:])
}
