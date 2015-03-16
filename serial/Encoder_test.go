package serial

import (
	check "gopkg.in/check.v1"
)

type EncoderSuite struct {
	coder PositioningCoder
	store *ByteStore
}

var _ = check.Suite(&EncoderSuite{})

func (suite *EncoderSuite) SetUpTest(c *check.C) {
	suite.store = NewByteStore()
	suite.coder = NewEncoder(suite.store)
}

func (suite *EncoderSuite) TestDataOnEmptyEncoderReturnsEmptyArray(c *check.C) {
	result := suite.store.Data()

	c.Assert(result, check.DeepEquals, make([]byte, 0))
}

func (suite *EncoderSuite) TestCodeUint32EncodesIntegerValue(c *check.C) {
	var value uint32 = 0x12345678
	suite.coder.CodeUint32(&value)

	result := suite.store.Data()

	c.Assert(result, check.DeepEquals, []byte{0x78, 0x56, 0x34, 0x12})
}

func (suite *EncoderSuite) TestCodeUint16EncodesIntegerValue(c *check.C) {
	var value uint16 = 0x3456
	suite.coder.CodeUint16(&value)

	result := suite.store.Data()

	c.Assert(result, check.DeepEquals, []byte{0x56, 0x34})
}

func (suite *EncoderSuite) TestCodeByteEncodesByteValue(c *check.C) {
	var value byte = 0x42
	suite.coder.CodeByte(&value)

	result := suite.store.Data()

	c.Assert(result, check.DeepEquals, []byte{0x42})
}

func (suite *EncoderSuite) TestSetCurPosRepositionsWritePointer(c *check.C) {
	value := uint32(0)
	suite.coder.CodeUint32(&value)
	value = 0x13243546

	suite.coder.SetCurPos(0)
	suite.coder.CodeUint32(&value)
	result := suite.store.Data()

	c.Assert(result, check.DeepEquals, []byte{0x46, 0x35, 0x24, 0x13})
}
