package voc

import (
	"bytes"
	"encoding/binary"

	check "gopkg.in/check.v1"
)

type LoadSuite struct {
}

var _ = check.Suite(&LoadSuite{})

func (suite *LoadSuite) TestLoadReturnsErrorOnNil(c *check.C) {
	_, err := Load(nil)

	c.Check(err, check.ErrorMatches, "source is nil")
}

func (suite *LoadSuite) newHeader() *bytes.Buffer {
	writer := bytes.NewBufferString(fileHeader)
	version := uint16(0x010A)
	headerSize := uint16(0x001A)

	binary.Write(writer, binary.LittleEndian, headerSize)
	binary.Write(writer, binary.LittleEndian, version)
	versionValidity := uint16(^version + uint16(0x1234))
	binary.Write(writer, binary.LittleEndian, versionValidity)

	return writer
}

func (suite *LoadSuite) TestLoadReturnsErrorOnInvalidVersion(c *check.C) {
	writer := suite.newHeader()

	writer.Write([]byte{0x00}) // Terminator

	data := writer.Bytes()
	data[24] = 0x00
	source := bytes.NewReader(data)
	_, err := Load(source)

	c.Check(err, check.ErrorMatches, "Version validity failed: 0x1129 != 0x1100")
}

func (suite *LoadSuite) TestLoadReturnsErrorOnValidButEmptyFile(c *check.C) {
	writer := suite.newHeader()

	writer.Write([]byte{0x00}) // Terminator

	source := bytes.NewReader(writer.Bytes())
	_, err := Load(source)

	c.Check(err, check.ErrorMatches, "No audio found")
}

func (suite *LoadSuite) TestLoadReturnsSoundDataOnSampleData(c *check.C) {
	writer := suite.newHeader()

	writer.Write([]byte{0x01})             // sound data
	writer.Write([]byte{0x03, 0x00, 0x00}) // block size
	writer.Write([]byte{0x64, 0x00})       // divisor, sound type
	writer.Write([]byte{0x80})             // one sample

	writer.Write([]byte{0x00}) // Terminator

	source := bytes.NewReader(writer.Bytes())
	data, err := Load(source)

	c.Assert(err, check.IsNil)
	c.Check(data, check.NotNil)
}

func (suite *LoadSuite) TestLoadReturnsSoundDataWithSampleRate(c *check.C) {
	writer := suite.newHeader()

	writer.Write([]byte{0x01})             // sound data
	writer.Write([]byte{0x03, 0x00, 0x00}) // block size
	writer.Write([]byte{0x9C, 0x00})       // divisor, sound type
	writer.Write([]byte{0x80})             // one sample

	writer.Write([]byte{0x00}) // Terminator

	source := bytes.NewReader(writer.Bytes())
	data, err := Load(source)

	c.Assert(err, check.IsNil)
	c.Check(data.SampleRate(), check.Equals, float32(10000.0))
}

func (suite *LoadSuite) TestLoadReturnsSoundDataWithSamples(c *check.C) {
	writer := suite.newHeader()

	writer.Write([]byte{0x01})                                     // sound data
	writer.Write([]byte{0x09, 0x00, 0x00})                         // block size
	writer.Write([]byte{0x9C, 0x00})                               // divisor, sound type
	writer.Write([]byte{0x80, 0xFF, 0x00, 0xC0, 0x40, 0x7F, 0x81}) // samples

	writer.Write([]byte{0x00}) // Terminator

	source := bytes.NewReader(writer.Bytes())
	data, err := Load(source)

	c.Assert(err, check.IsNil)
	expected := []int16{0x0000, 0x7FFF, -0x8000, 0x4081, -0x4080, -0x0102, 0x0103}
	samples := make([]int16, len(expected))
	data.Samples(samples)
	c.Check(samples, check.DeepEquals, expected)
}
