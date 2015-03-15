package base

import (
	check "gopkg.in/check.v1"
)

type WordSuite struct{}

var _ = check.Suite(&WordSuite{})

func (suite *WordSuite) TestPartFrom0ReturnsFirst8Bits(c *check.C) {
	value := word(0x2040)

	c.Assert(value.partFrom(0, 8), check.Equals, byte(0x81))
}

func (suite *WordSuite) TestPartFrom8ReturnsLast6Bits(c *check.C) {
	value := word(0x003F)

	c.Assert(value.partFrom(8, 6), check.Equals, byte(0x3F))
}

func (suite *WordSuite) TestPartFrom0ReturnsFirst2Bits(c *check.C) {
	value := word(0x3000)

	c.Assert(value.partFrom(0, 2), check.Equals, byte(0x03))
}

func (suite *WordSuite) TestPartFrom4ReturnsFirst4Bits(c *check.C) {
	value := word(0xFE70)

	c.Assert(value.partFrom(4, 4), check.Equals, byte(0x09))
}
