package dos

import (
	check "gopkg.in/check.v1"

	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/serial"
)

type FormatWriterSuite struct {
	provider chunk.Provider
}

var _ = check.Suite(&FormatWriterSuite{})

func (suite *FormatWriterSuite) SetUpTest(c *check.C) {

}

func (suite *FormatWriterSuite) TestFinishWithoutAddingCreatesValidFileWithoutChunks(c *check.C) {
	expected := emptyResourceFile()
	dest := serial.NewByteStore()
	consumer := NewChunkConsumer(dest)

	consumer.Finish()
	result := dest.Data()

	c.Assert(result, check.DeepEquals, expected)
}
