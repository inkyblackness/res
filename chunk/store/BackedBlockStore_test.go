package store

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"

	check "gopkg.in/check.v1"
)

type BackedBlockStoreSuite struct {
}

var _ = check.Suite(&BackedBlockStoreSuite{})

func (suite *BackedBlockStoreSuite) TestChunkTypeReturnsValueFromHolder(c *check.C) {
	holder := chunk.NewBlockHolder(chunk.BasicChunkType, res.Palette, [][]byte{nil})

	backed := newBackedBlockStore(holder, func() {})

	c.Check(backed.ChunkType(), check.Equals, chunk.BasicChunkType)
}

func (suite *BackedBlockStoreSuite) TestContentTypeReturnsValueFromHolder(c *check.C) {
	holder := chunk.NewBlockHolder(chunk.BasicChunkType, res.Palette, [][]byte{nil})

	backed := newBackedBlockStore(holder, func() {})

	c.Check(backed.ContentType(), check.Equals, res.Palette)
}

func (suite *BackedBlockStoreSuite) TestBlockCountReturnsValueFromHolder(c *check.C) {
	holder := chunk.NewBlockHolder(chunk.BasicChunkType, res.Palette, [][]byte{nil, nil})

	backed := newBackedBlockStore(holder, func() {})

	c.Check(backed.BlockCount(), check.Equals, uint16(2))
}

func (suite *BackedBlockStoreSuite) TestGetReturnsBlockDataFromHolder_WhenUnchanged(c *check.C) {
	holder := chunk.NewBlockHolder(chunk.BasicChunkType, res.Palette, [][]byte{[]byte{0x01}, []byte{0x02}})

	backed := newBackedBlockStore(holder, func() {})
	result := backed.Get(uint16(1))

	c.Check(result, check.DeepEquals, []byte{0x02})
}

func (suite *BackedBlockStoreSuite) TestOnModifiedIsCalled_WhenBlockIsModified(c *check.C) {
	holder := chunk.NewBlockHolder(chunk.BasicChunkType, res.Palette, [][]byte{[]byte{0x01}, []byte{0x02}})
	called := false

	backed := newBackedBlockStore(holder, func() { called = true })
	backed.Put(uint16(1), []byte{0x03})

	c.Check(called, check.Equals, true)
}

func (suite *BackedBlockStoreSuite) TestGetReturnsNewBlockData_WhenBlockIsModified(c *check.C) {
	holder := chunk.NewBlockHolder(chunk.BasicChunkType, res.Palette, [][]byte{[]byte{0x01}, []byte{0x02}})

	backed := newBackedBlockStore(holder, func() {})
	backed.Put(uint16(1), []byte{0x04})
	result := backed.Get(uint16(1))

	c.Check(result, check.DeepEquals, []byte{0x04})
}