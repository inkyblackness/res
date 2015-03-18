package dos

import (
	//"bytes"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/serial"

	check "gopkg.in/check.v1"
)

type FormatWriterSuite struct {
	store    *serial.ByteStore
	consumer chunk.Consumer
}

var _ = check.Suite(&FormatWriterSuite{})

func (suite *FormatWriterSuite) SetUpTest(c *check.C) {
	suite.store = serial.NewByteStore()
	suite.consumer = NewChunkConsumer(suite.store)
}

func (suite *FormatWriterSuite) TestFinishWithoutAddingCreatesValidFileWithoutChunks(c *check.C) {
	expected := emptyResourceFile()

	suite.consumer.Finish()
	result := suite.store.Data()

	c.Assert(result, check.DeepEquals, expected)
}

func (suite *FormatWriterSuite) TestConsumeOfFlatUncompressedChunkCanBeWritten(c *check.C) {
	singleBlock := []byte{0xAB, 0x01, 0xCD, 0x02, 0xEF}
	blockHolder := chunk.NewBlockHolder(chunk.BasicChunkType, res.Data, [][]byte{singleBlock})

	suite.consumer.Consume(res.ResourceID(0x1234), blockHolder)
	suite.consumer.Finish()

	result := suite.store.Data()

	expected := []byte{}
	expected = append(expected, singleBlock...)
	expected = append(expected, 0x00, 0x00, 0x00)       // alignment for directory
	expected = append(expected, 0x01, 0x00)             // chunk count
	expected = append(expected, 0x80, 0x00, 0x00, 0x00) // offset to first chunk
	expected = append(expected, 0x34, 0x12)             // chunk ID
	expected = append(expected, 0x05, 0x00, 0x00)       // chunk length (uncompressed)
	expected = append(expected, 0x00)                   // chunk type
	expected = append(expected, 0x05, 0x00, 0x00)       // chunk length in file
	expected = append(expected, 0x00)                   // content type
	c.Assert(result[ChunkDirectoryFileOffsetPos+4:], check.DeepEquals, expected)
}
