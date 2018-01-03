package resfile

import (
	"bytes"
	"encoding/binary"
)

func emptyResourceFile() []byte {
	buf := bytes.NewBufferString(headerString)
	headerTrailer := make([]byte, chunkDirectoryFileOffsetPos-buf.Len())
	headerTrailer[0] = commentTerminator

	binary.Write(buf, binary.LittleEndian, headerTrailer)
	dictionaryOffset := uint32(buf.Len() + 4)
	binary.Write(buf, binary.LittleEndian, &dictionaryOffset)

	numberOfChunks := uint16(0)
	firstChunkOffset := uint32(buf.Len())

	binary.Write(buf, binary.LittleEndian, &numberOfChunks)
	binary.Write(buf, binary.LittleEndian, &firstChunkOffset)

	return buf.Bytes()
}
