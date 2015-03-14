package dos

import (
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
