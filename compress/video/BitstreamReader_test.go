package video

import (
	check "gopkg.in/check.v1"
)

type BitstreamReaderSuite struct {
}

var _ = check.Suite(&BitstreamReaderSuite{})

func (suite *BitstreamReaderSuite) TestReadPanicsForMoreThan32Bits(c *check.C) {
	reader := NewBitstreamReader([]byte{0x11, 0x22, 0x33, 0x44})

	c.Check(func() { reader.Read(33) }, check.Panics, "Limit of bit count: 32")
}

func (suite *BitstreamReaderSuite) TestReadReturnsErrorIfGoingBeyondEnd(c *check.C) {
	reader := NewBitstreamReader([]byte{0xAF})

	_, err := reader.Read(9)

	c.Check(err, check.Equals, BitstreamEndError)
}

func (suite *BitstreamReaderSuite) TestReadReturnsValueOfRequestedBitSize(c *check.C) {
	reader := NewBitstreamReader([]byte{0xAF})

	result, _ := reader.Read(3)

	c.Check(result, check.Equals, uint32(5))
}

func (suite *BitstreamReaderSuite) TestRepeatedReadReturnsSameValue(c *check.C) {
	reader := NewBitstreamReader([]byte{0xAF})

	result1, _ := reader.Read(3)
	result2, _ := reader.Read(3)

	c.Check(result1, check.Equals, uint32(5))
	c.Check(result2, check.Equals, result1)
}

func (suite *BitstreamReaderSuite) TestAdvancePanicsForNegativeValues(c *check.C) {
	reader := NewBitstreamReader([]byte{0x11, 0x22, 0x33, 0x44})

	c.Check(func() { reader.Advance(-10) }, check.Panics, "Can only advance forward")
}

func (suite *BitstreamReaderSuite) TestAdvanceReturnsErrorIfGoingBeyondEnd(c *check.C) {
	reader := NewBitstreamReader([]byte{0xFF})

	err := reader.Advance(9)

	c.Check(err, check.Equals, BitstreamEndError)
}

func (suite *BitstreamReaderSuite) TestAdvanceLetsReadFurtherBits(c *check.C) {
	reader := NewBitstreamReader([]byte{0xAF})

	reader.Advance(2)
	result, _ := reader.Read(4)

	c.Check(result, check.Equals, uint32(0x0B))
}

func (suite *BitstreamReaderSuite) TestAdvanceToEndIsPossible(c *check.C) {
	reader := NewBitstreamReader([]byte{0xAF})

	err := reader.Advance(8)

	c.Check(err, check.IsNil)
}

func (suite *BitstreamReaderSuite) TestInternalBufferDoesntLoseData(c *check.C) {
	reader := NewBitstreamReader([]byte{0x7F, 0xFF, 0xFF, 0xFF, 0x80})

	reader.Advance(1)
	result, _ := reader.Read(32)

	c.Check(result, check.Equals, uint32(0xFFFFFFFF))
}

func (suite *BitstreamReaderSuite) TestAdvanceCanJumpToLastBit(c *check.C) {
	reader := NewBitstreamReader([]byte{0x00, 0x00, 0x01})

	reader.Advance(23)
	result, _ := reader.Read(1)

	c.Check(result, check.Equals, uint32(1))
}

func (suite *BitstreamReaderSuite) TestReadAdvanceBeyondFirstRead(c *check.C) {
	reader := NewBitstreamReader([]byte{0xFF, 0x00, 0xFA})

	reader.Read(10)
	reader.Advance(20)
	result, _ := reader.Read(4)

	c.Check(result, check.Equals, uint32(0x0A))
}

func (suite *BitstreamReaderSuite) TestReadAdvanceWithinFirstRead(c *check.C) {
	reader := NewBitstreamReader([]byte{0xFF, 0x00, 0xFA})

	reader.Read(10)
	reader.Advance(4)
	result, _ := reader.Read(20)

	c.Check(result, check.Equals, uint32(0xF00FA))
}

func (suite *BitstreamReaderSuite) TestReadOfZeroBitsIsPossibleMidStream(c *check.C) {
	reader := NewBitstreamReader([]byte{0xFF})

	reader.Read(4)
	result, err := reader.Read(0)

	c.Assert(err, check.IsNil)
	c.Check(result, check.Equals, uint32(0))
}

func (suite *BitstreamReaderSuite) TestReadOfZeroBitsIsPossibleAtEnd(c *check.C) {
	reader := NewBitstreamReader([]byte{0xFF})

	reader.Advance(8)
	result, err := reader.Read(0)

	c.Assert(err, check.IsNil)
	c.Check(result, check.Equals, uint32(0))
}

func (suite *BitstreamReaderSuite) TestReadOfZeroBitsIsPossibleWithEmptySource(c *check.C) {
	reader := NewBitstreamReader([]byte{})

	result, err := reader.Read(0)

	c.Assert(err, check.IsNil)
	c.Check(result, check.Equals, uint32(0))
}
