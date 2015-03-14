package dos

import (
	"bytes"
	"io"

	"github.com/inkyblackness/res/serial"
)

func emptyResourceFile() io.ReadSeeker {
	encoder := serial.NewEncoder()

	codeHeader(encoder)
	// write offset to dictionary - in this case right after header
	{
		dictionaryOffset := uint32(len(encoder.Data()) + 4)
		encoder.CodeUint32(&dictionaryOffset)
	}
	{
		numberOfChunks := uint16(0)
		firstChunkOffset := uint32(0)

		encoder.CodeUint16(&numberOfChunks)
		encoder.CodeUint32(&firstChunkOffset)
	}

	return bytes.NewReader(encoder.Data())
}

func codeHeader(coder serial.Coder) {
	var blank byte = 0x00
	headerString := HeaderString
	commentTerminator := CommentTerminator

	coder.CodeString(&headerString)
	coder.CodeByte(&commentTerminator)
	for coder.Len() < ChunkDirectoryFileOffsetPos {
		coder.CodeByte(&blank)
	}
}
