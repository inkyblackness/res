package base

import (
	"bytes"
	"io"
	"math/rand"
	"time"

	"github.com/inkyblackness/res/serial"

	check "gopkg.in/check.v1"
)

type DecompressorSuite struct {
	store      *serial.ByteStore
	compressor io.WriteCloser
}

var _ = check.Suite(&DecompressorSuite{})

func (suite *DecompressorSuite) SetUpTest(c *check.C) {
}

func (suite *DecompressorSuite) TestDecompressTest1(c *check.C) {
	input := []byte{0x00, 0x01, 0x00, 0x01}

	suite.verify(c, input)
}

func (suite *DecompressorSuite) TestDecompressTest2(c *check.C) {
	input := []byte{0x00, 0x01, 0x00, 0x01, 0x00, 0x01}

	suite.verify(c, input)
}

func (suite *DecompressorSuite) TestDecompressTest3(c *check.C) {
	input := []byte{}

	suite.verify(c, input)
}

func (suite *DecompressorSuite) TestDecompressTest4(c *check.C) {
	input := []byte{0x00, 0x01, 0x00, 0x02, 0x01, 0x00, 0x01}

	suite.verify(c, input)
}

func (suite *DecompressorSuite) TestDecompressTest5(c *check.C) {
	input := []byte{0x00, 0x01, 0x00, 0x02, 0x01, 0x00, 0x01, 0x02, 0x01, 0x02, 0x01, 0x00, 0x01, 0x02}

	suite.verify(c, input)
}

func (suite *DecompressorSuite) TestDecompressTestRandom(c *check.C) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for testCase := 0; testCase < 1; testCase++ {
		input := make([]byte, r.Intn(1024))
		for i := 0; i < len(input); i++ {
			input[i] = byte(r.Intn(256))
		}
		suite.verify(c, input)
	}
}

func (suite *DecompressorSuite) verify(c *check.C, input []byte) {
	suite.store = serial.NewByteStore()
	suite.compressor = NewCompressor(serial.NewEncoder(suite.store))

	suite.compressor.Write(input)
	suite.compressor.Close()

	output := make([]byte, len(input))
	for i := range output {
		output[i] = 0xFF
	}
	source := bytes.NewReader(suite.store.Data())
	decompressor := NewDecompressor(serial.NewDecoder(source))
	decompressor.Read(output)

	c.Assert(output, check.DeepEquals, input)
}
