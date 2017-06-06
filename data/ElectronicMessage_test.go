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

func (suite *ElectronicMessageSuite) TestEncodeMeta_A(c *check.C) {
	message := NewElectronicMessage()

	message.SetNextMessage(0x20)
	message.SetColorIndex(0x13)
	message.SetLeftDisplay(30)
	message.SetRightDisplay(40)

	holder := message.Encode(suite.cp)

	c.Assert(holder, check.NotNil)
	c.Assert(holder.BlockCount() > 0, check.Equals, true)
	c.Check(holder.BlockData(0), check.DeepEquals, suite.cp.Encode("i20 c13 30,40"))
}

func (suite *ElectronicMessageSuite) TestEncodeMeta_B(c *check.C) {
	message := NewElectronicMessage()

	message.SetInterrupt(true)
	message.SetLeftDisplay(31)

	holder := message.Encode(suite.cp)

	c.Assert(holder, check.NotNil)
	c.Assert(holder.BlockCount() > 0, check.Equals, true)
	c.Check(holder.BlockData(0), check.DeepEquals, suite.cp.Encode("t 31"))
}

func (suite *ElectronicMessageSuite) TestEncodeMeta_C(c *check.C) {
	message := NewElectronicMessage()

	message.SetInterrupt(true)

	holder := message.Encode(suite.cp)

	c.Assert(holder, check.NotNil)
	c.Assert(holder.BlockCount() > 0, check.Equals, true)
	c.Check(holder.BlockData(0), check.DeepEquals, suite.cp.Encode("t"))
}

func (suite *ElectronicMessageSuite) TestEncodeFillsEmptyTextLinesWithABlank(c *check.C) {
	message := NewElectronicMessage()

	message.SetVerboseText("line1\n\n\nline2")
	message.SetTerseText("terse1\n\n\nterse2")

	holder := message.Encode(suite.cp)

	c.Assert(holder, check.NotNil)
	c.Assert(holder.BlockCount() > 0, check.Equals, true)
	c.Check(holder.BlockData(4), check.DeepEquals, suite.cp.Encode("line1\n"))
	c.Check(holder.BlockData(5), check.DeepEquals, suite.cp.Encode(" \n"))
	c.Check(holder.BlockData(6), check.DeepEquals, suite.cp.Encode(" \n"))
	c.Check(holder.BlockData(7), check.DeepEquals, suite.cp.Encode("line2"))
	c.Check(holder.BlockData(9), check.DeepEquals, suite.cp.Encode("terse1\n"))
	c.Check(holder.BlockData(10), check.DeepEquals, suite.cp.Encode(" \n"))
	c.Check(holder.BlockData(11), check.DeepEquals, suite.cp.Encode(" \n"))
	c.Check(holder.BlockData(12), check.DeepEquals, suite.cp.Encode("terse2"))
}

func (suite *ElectronicMessageSuite) TestEncodeBreaksUpLinesAfterLimitCharacters(c *check.C) {
	message := NewElectronicMessage()

	message.SetVerboseText("aaaaaaaaa bbbbbbbbb ccccccccc ddddddddd eeeeeeeee fffffffff ggggggggg hhhhhhhhh iiiiiiiii jjjjjjjjj kkkkk")

	holder := message.Encode(suite.cp)

	c.Assert(holder, check.NotNil)
	c.Assert(holder.BlockCount() > 0, check.Equals, true)
	c.Check(holder.BlockData(4), check.DeepEquals,
		suite.cp.Encode("aaaaaaaaa bbbbbbbbb ccccccccc ddddddddd eeeeeeeee fffffffff ggggggggg hhhhhhhhh iiiiiiiii jjjjjjjjj "))
	c.Check(holder.BlockData(5), check.DeepEquals,
		suite.cp.Encode("kkkkk"))
}
