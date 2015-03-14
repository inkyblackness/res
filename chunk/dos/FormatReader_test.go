package dos

import (
	"bytes"

	check "gopkg.in/check.v1"

	"github.com/inkyblackness/res/chunk"
)

type FormatReaderSuite struct {
	provider chunk.Provider
}

var _ = check.Suite(&FormatReaderSuite{})

func (suite *FormatReaderSuite) SetUpTest(c *check.C) {

}

func (suite *FormatReaderSuite) TestNewChunkProviderReturnsErrorOnNil(c *check.C) {
	_, err := NewChunkProvider(nil)

	c.Assert(err, check.ErrorMatches, "source is nil")
}

func (suite *FormatReaderSuite) TestNewChunkProviderReturnsProviderOnEmptySource(c *check.C) {
	source := bytes.NewReader(emptyResourceFile())
	provider, _ := NewChunkProvider(source)

	c.Assert(provider, check.NotNil)
}

func (suite *FormatReaderSuite) TestNewChunkProviderReturnsErrorOnInvalidHeaderString(c *check.C) {
	sourceData := emptyResourceFile()
	sourceData[10] = byte("A"[0])

	_, err := NewChunkProvider(bytes.NewReader(sourceData))

	c.Assert(err, check.ErrorMatches, "Format mismatch")
}

func (suite *FormatReaderSuite) TestNewChunkProviderReturnsErrorOnMissingCommentTerminator(c *check.C) {
	sourceData := emptyResourceFile()
	sourceData[len(HeaderString)] = byte(0)

	_, err := NewChunkProvider(bytes.NewReader(sourceData))

	c.Assert(err, check.ErrorMatches, "Format mismatch")
}
