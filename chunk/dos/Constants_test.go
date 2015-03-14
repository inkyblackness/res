package dos

import (
	"bytes"
	"io"

	"github.com/inkyblackness/res/serial"
)

func emptyResourceFile() io.ReadSeeker {
	store := serial.NewByteStore()
	encoder := serial.NewEncoder(store)

	codeHeader(encoder, store)
	// write offset to dictionary - in this case right after header
	{
		dictionaryOffset := uint32(store.Len() + 4)
		encoder.CodeUint32(&dictionaryOffset)
	}
	{
		numberOfChunks := uint16(0)
		firstChunkOffset := uint32(0)

		encoder.CodeUint16(&numberOfChunks)
		encoder.CodeUint32(&firstChunkOffset)
	}

	return bytes.NewReader(store.Data())
}

func codeHeader(coder serial.Coder, store *serial.ByteStore) {
	var blank byte = 0x00
	commentTerminator := CommentTerminator

	coder.CodeBytes([]byte(HeaderString))
	coder.CodeByte(&commentTerminator)
	for store.Len() < ChunkDirectoryFileOffsetPos {
		coder.CodeByte(&blank)
	}
}
