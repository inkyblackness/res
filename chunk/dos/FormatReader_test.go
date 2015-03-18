package dos

import (
	"bytes"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/serial"

	check "gopkg.in/check.v1"
)

type FormatReaderSuite struct {
}

var _ = check.Suite(&FormatReaderSuite{})

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

func (suite *FormatReaderSuite) TestNewChunkProviderReturnsErrorOnInvalidDirectoryStart(c *check.C) {
	sourceData := emptyResourceFile()
	sourceData[ChunkDirectoryFileOffsetPos] = byte(0xFF)

	_, err := NewChunkProvider(bytes.NewReader(sourceData))

	c.Assert(err, check.ErrorMatches, "EOF")
}

func (suite *FormatReaderSuite) TestIDsReturnsTheStoredChunkIDsInOrder(c *check.C) {
	store := serial.NewByteStore()
	consumer := NewChunkConsumer(store)

	blockHolder1 := chunk.NewBlockHolder(chunk.BasicChunkType, res.Data, [][]byte{[]byte{}})
	consumer.Consume(res.ResourceID(0x5678), blockHolder1)
	blockHolder2 := chunk.NewBlockHolder(chunk.BasicChunkType, res.Data, [][]byte{[]byte{}})
	consumer.Consume(res.ResourceID(0x1234), blockHolder2)
	consumer.Finish()

	provider, _ := NewChunkProvider(bytes.NewReader(store.Data()))

	c.Assert(provider.IDs(), check.DeepEquals, []res.ResourceID{0x5678, 0x1234})
}

func (suite *FormatReaderSuite) TestProvideReturnsABlockProviderForKnownID(c *check.C) {
	store := serial.NewByteStore()
	consumer := NewChunkConsumer(store)

	blockHolder1 := chunk.NewBlockHolder(chunk.BasicChunkType, res.Data, [][]byte{[]byte{}})
	consumer.Consume(res.ResourceID(0x1122), blockHolder1)
	consumer.Finish()

	provider, _ := NewChunkProvider(bytes.NewReader(store.Data()))

	c.Assert(provider.Provide(0x1122), check.NotNil)
}

func (suite *FormatReaderSuite) TestProvideReturnsABlockProviderWithContent(c *check.C) {
	store := serial.NewByteStore()
	consumer := NewChunkConsumer(store)

	blockHolder1 := chunk.NewBlockHolder(chunk.BasicChunkType, res.Data, [][]byte{[]byte{0xAA, 0xBB, 0xCC}})
	consumer.Consume(res.ResourceID(0x3344), blockHolder1)
	consumer.Finish()

	provider, _ := NewChunkProvider(bytes.NewReader(store.Data()))

	c.Assert(provider.Provide(0x3344).BlockData(0), check.DeepEquals, []byte{0xAA, 0xBB, 0xCC})
}
