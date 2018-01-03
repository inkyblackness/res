package serial

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DecoderSuite struct {
	suite.Suite
	errorBuf errorBuffer
	coder    Coder
}

func TestDecoderSuite(t *testing.T) {
	suite.Run(t, new(DecoderSuite))
}

func (suite *DecoderSuite) TestCodeUint32() {
	suite.whenDecodingFrom([]byte{0x78, 0x56, 0x34, 0x12})

	var value uint32
	suite.coder.Code(&value)

	assert.Equal(suite.T(), uint32(0x12345678), value)
}

func (suite *DecoderSuite) TestCodeUint16() {
	suite.whenDecodingFrom([]byte{0x45, 0x23})

	var value uint16
	suite.coder.Code(&value)

	assert.Equal(suite.T(), uint16(0x2345), value)
}

func (suite *DecoderSuite) TestCodeByte() {
	suite.whenDecodingFrom([]byte{0xAB})

	var value byte
	suite.coder.Code(&value)

	assert.Equal(suite.T(), byte(0xAB), value)
}

func (suite *DecoderSuite) TestCodeByteSlice() {
	suite.whenDecodingFrom([]byte{0x78, 0x12, 0x34})

	value := make([]byte, 3)
	suite.coder.Code(value)

	assert.Equal(suite.T(), []byte{0x78, 0x12, 0x34}, value)
}

func (suite *DecoderSuite) TestFirstErrorReturnsFirstError() {
	suite.whenDecodingWithErrors()

	suite.coder.Code(uint32(0))
	suite.errorBuf.errorOnNextCall = true
	suite.coder.Code(uint16(0))

	assert.EqualError(suite.T(), suite.coder.FirstError(), "errorBuffer on call number 2")
}

func (suite *DecoderSuite) TestFirstErrorIgnoresFurtherErrors() {
	suite.whenDecodingWithErrors()

	suite.errorBuf.errorOnNextCall = true
	suite.coder.Code(uint32(0))
	suite.errorBuf.errorOnNextCall = true
	suite.coder.Code(uint32(0))

	assert.Equal(suite.T(), suite.errorBuf.callCounter, 1)
	assert.EqualError(suite.T(), suite.coder.FirstError(), "errorBuffer on call number 1")
}

func (suite *DecoderSuite) whenDecodingFrom(data []byte) {
	source := bytes.NewReader(data)
	suite.coder = NewDecoder(source)
}

func (suite *DecoderSuite) whenDecodingWithErrors() {
	suite.coder = NewDecoder(&suite.errorBuf)
}
