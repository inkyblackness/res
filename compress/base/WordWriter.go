package base

import "github.com/inkyblackness/res/serial"

type wordWriter struct {
	coder         serial.Coder
	remainder     byte
	remainderFree uint
}

func newWordWriter(coder serial.Coder) *wordWriter {
	writer := &wordWriter{coder: coder, remainder: 0, remainderFree: 8}

	return writer
}

func (writer *wordWriter) close() {
	writer.write(endOfStream)
	if writer.remainderFree < 8 {
		writer.writeByte(writer.remainder)
	}
	writer.writeByte(byte(0x00))
}

func (writer *wordWriter) write(value word) {
	remaining := bitsPerWord
	for remaining >= writer.remainderFree {
		writer.remainder |= value.partFrom(bitsPerWord-remaining, writer.remainderFree)
		remaining -= writer.remainderFree

		writer.writeByte(writer.remainder)
		writer.remainder = byte(0x00)
		writer.remainderFree = 8
	}
	writer.remainderFree = 8 - remaining
	writer.remainder = value.partFrom(bitsPerWord-remaining, remaining) << writer.remainderFree
}

func (writer *wordWriter) writeByte(value byte) {
	writer.coder.CodeByte(&value)
}
