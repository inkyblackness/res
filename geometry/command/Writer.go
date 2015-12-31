package command

import (
	"bytes"
	"encoding/binary"
)

// Writer implements helper functions to write the commands
type Writer struct {
	buf *bytes.Buffer
}

// NewWriter returns a writer instance.
func NewWriter() *Writer {
	return &Writer{
		buf: bytes.NewBuffer(nil)}
}

func (writer *Writer) write16(value uint16) {
	binary.Write(writer.buf, binary.LittleEndian, &value)
}

func (writer *Writer) write32(value uint32) {
	binary.Write(writer.buf, binary.LittleEndian, &value)
}

// Bytes returns the current byte buffer of the writer
func (writer *Writer) Bytes() []byte {
	return writer.buf.Bytes()
}

// WriteHeader writes the command header.
func (writer *Writer) WriteHeader(faceCount int) {
	writer.write16(0x0027)
	writer.write16(0x0008)
	writer.write16(0x0002)
	writer.write16(uint16(faceCount))
}

// WriteDefineVertex writes the command.
func (writer *Writer) WriteDefineVertex(vector Vector) {
	writer.write16(uint16(CmdDefineVertex))
	writer.write16(uint16(0))
	writer.write32(uint32(vector.X))
	writer.write32(uint32(vector.Y))
	writer.write32(uint32(vector.Z))
}

// WriteDefineVertices writes the command.
func (writer *Writer) WriteDefineVertices(vectors []Vector) {
	writer.write16(uint16(CmdDefineVertices))
	writer.write16(uint16(len(vectors)))
	writer.write16(uint16(0))
	for _, vector := range vectors {
		writer.write32(uint32(vector.X))
		writer.write32(uint32(vector.Y))
		writer.write32(uint32(vector.Z))
	}
}

// WriteDefineOneOffsetVertex writes the command.
func (writer *Writer) WriteDefineOneOffsetVertex(cmd ModelCommandID, newIndex int, referenceIndex int, offset float32) {
	writer.write16(uint16(cmd))
	writer.write16(uint16(newIndex))
	writer.write16(uint16(referenceIndex))
	writer.write32(uint32(ToFixed(offset)))
}

// WriteDefineTwoOffsetVertex writes the command.
func (writer *Writer) WriteDefineTwoOffsetVertex(cmd ModelCommandID, newIndex int, referenceIndex int, offset1 float32, offset2 float32) {
	writer.write16(uint16(cmd))
	writer.write16(uint16(newIndex))
	writer.write16(uint16(referenceIndex))
	writer.write32(uint32(ToFixed(offset1)))
	writer.write32(uint32(ToFixed(offset2)))
}

// WriteEndOfNode writes the command.
func (writer *Writer) WriteEndOfNode() {
	writer.write16(uint16(CmdEndOfNode))
}
