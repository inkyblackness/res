package objprop

import (
	check "gopkg.in/check.v1"
)

type ConstantsSuite struct {
}

var _ = check.Suite(&ConstantsSuite{})

func (suite *ConstantsSuite) TestStandardPropertiesReturnsProperLength(c *check.C) {
	descriptor := StandardProperties()
	totalLength := uint32(4)

	for _, classDesc := range descriptor {
		totalLength += classDesc.TotalDataLength()
	}

	c.Assert(totalLength, check.Equals, uint32(17951)) // as taken from original CD
}

func (suite *ConstantsSuite) TestStandardPropertiesReturnsProperAmount(c *check.C) {
	descriptor := StandardProperties()
	total := uint32(0)

	for _, classDesc := range descriptor {
		total += classDesc.TotalTypeCount()
	}

	c.Assert(total, check.Equals, uint32(476))
}
