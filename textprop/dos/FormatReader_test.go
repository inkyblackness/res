package dos

import (
	"bytes"

	"github.com/inkyblackness/res/textprop"

	check "gopkg.in/check.v1"
)

type FormatReaderSuite struct {
}

var _ = check.Suite(&FormatReaderSuite{})

func (suite *FormatReaderSuite) TestNewProviderReturnsErrorOnNil(c *check.C) {
	_, err := NewProvider(nil)

	c.Assert(err, check.ErrorMatches, "source is nil")
}

func (suite *FormatReaderSuite) TestNewProviderReturnsErrorOnFileWithWrongSize(c *check.C) {
	sourceData := make([]byte, textprop.TexturePropertiesLength-1)

	_, err := NewProvider(bytes.NewReader(sourceData))

	c.Assert(err, check.ErrorMatches, "Format mismatch")
}

func (suite *FormatReaderSuite) TestNewProviderReturnsProviderWithCountSet(c *check.C) {
	source := bytes.NewReader(make([]byte, textprop.TexturePropertiesLength*4))
	provider, _ := NewProvider(source)

	c.Assert(provider.TextureCount(), check.Equals, uint32(4))
}

func (suite *FormatReaderSuite) TestNewProviderReturnsProviderWithData(c *check.C) {
	sourceData := make([]byte, textprop.TexturePropertiesLength*3)
	for i := byte(0); i < byte(len(sourceData)); i++ {
		sourceData[i] = i
	}
	source := bytes.NewReader(sourceData)
	provider, _ := NewProvider(source)

	c.Assert(provider.Provide(1), check.DeepEquals, []byte{0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15})
}
