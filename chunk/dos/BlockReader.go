package dos

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/serial"
)

type blockReader struct {
	coder serial.Coder

	startOffset        uint32
	compressedLength   uint32
	uncompressedLength uint32

	typeID     chunk.TypeID
	dataTypeID res.DataTypeID
}
