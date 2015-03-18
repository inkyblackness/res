package dos

import (
	"fmt"
	"io"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"
	"github.com/inkyblackness/res/serial"
)

type formatReader struct {
	resourceIDs    []res.ResourceID
	chunkAddresses map[res.ResourceID]*chunkAddress
}

var errFormatMismatch = fmt.Errorf("Format mismatch")

// NewChunkProvider returns a chunk provider reading from a random access reader
// over a DOS format resource file.
func NewChunkProvider(source io.ReadSeeker) (provider chunk.Provider, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	if source == nil {
		panic(fmt.Errorf("source is nil"))
	}
	coder := serial.NewPositioningDecoder(source)

	skipAndVerifyHeaderString(coder)
	skipAndVerifyComment(coder)
	ids, addresses := readAndVerifyDirectory(coder)

	provider = &formatReader{resourceIDs: ids,
		chunkAddresses: addresses}

	return
}

func (reader *formatReader) IDs() []res.ResourceID {
	return reader.resourceIDs
}

// Provide implements the chunk.Provider interface
func (reader *formatReader) Provide(id res.ResourceID) chunk.BlockHolder {
	return nil
}

func skipAndVerifyHeaderString(coder serial.Coder) {
	headerStringBuffer := make([]byte, len(HeaderString))
	coder.CodeBytes(headerStringBuffer)
	if string(headerStringBuffer) != HeaderString {
		panic(errFormatMismatch)
	}
}

func skipAndVerifyComment(coder serial.PositioningCoder) {
	terminatorFound := false

	for remaining := ChunkDirectoryFileOffsetPos - coder.CurPos(); remaining > 0; remaining-- {
		temp := byte(0x00)
		coder.CodeByte(&temp)
		if temp == CommentTerminator {
			terminatorFound = true
		}
	}
	if !terminatorFound {
		panic(errFormatMismatch)
	}
}

func readAndVerifyDirectory(coder serial.PositioningCoder) ([]res.ResourceID, map[res.ResourceID]*chunkAddress) {
	directoryFileOffset := uint32(0)
	directoryEntries := uint16(0)
	firstChunkFileOffset := uint32(0)

	coder.CodeUint32(&directoryFileOffset)
	coder.SetCurPos(directoryFileOffset)

	coder.CodeUint16(&directoryEntries)
	coder.CodeUint32(&firstChunkFileOffset)
	ids := make([]res.ResourceID, int(directoryEntries))
	addresses := make(map[res.ResourceID]*chunkAddress)

	for i := uint16(0); i < directoryEntries; i++ {
		resourceID := uint16(0xFFFF)
		address := &chunkAddress{}

		coder.CodeUint16(&resourceID)
		address.code(coder)

		ids[i] = res.ResourceID(resourceID)
		addresses[ids[i]] = address
	}

	return ids, addresses
}
