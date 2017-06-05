package data

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/text"

	check "gopkg.in/check.v1"
)

type ElectronicMessageSuite struct {
	cp text.Codepage
}

var _ = check.Suite(&ElectronicMessageSuite{})

func (suite *ElectronicMessageSuite) SetUpTest(c *check.C) {
	suite.cp = text.DefaultCodepage()
}

func (suite *ElectronicMessageSuite) TestEncodeBasicMessage(c *check.C) {
	message := NewElectronicMessage()

	message.SetTitle("1")
	message.SetSender("2")
	message.SetSubject("3")
	message.SetVerboseText("4")
	message.SetTerseText("5")

	holder := message.Encode(suite.cp)

	c.Assert(holder, check.NotNil)
	c.Check(holder.ChunkType(), check.Equals, chunk.BasicChunkType.WithDirectory())
	c.Check(holder.ContentType(), check.Equals, res.Text)
	c.Assert(holder.BlockCount(), check.Equals, uint16(8))
	c.Check(holder.BlockData(0), check.DeepEquals, []byte{0x00})
	c.Check(holder.BlockData(1), check.DeepEquals, []byte{0x31, 0x00})
	c.Check(holder.BlockData(2), check.DeepEquals, []byte{0x32, 0x00})
	c.Check(holder.BlockData(3), check.DeepEquals, []byte{0x33, 0x00})
	c.Check(holder.BlockData(4), check.DeepEquals, []byte{0x34, 0x00})
	c.Check(holder.BlockData(5), check.DeepEquals, []byte{0x00})
	c.Check(holder.BlockData(6), check.DeepEquals, []byte{0x35, 0x00})
	c.Check(holder.BlockData(7), check.DeepEquals, []byte{0x00})
}
