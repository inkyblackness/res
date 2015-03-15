package base

import "github.com/inkyblackness/res/serial"

type wordWriter struct {
	coder         serial.Coder
	remainder     byte
	remainderFree uint
}

func newWordWriter(coder *serial.Encoder) *wordWriter {
	writer := &wordWriter{coder: coder, remainder: 0, remainderFree: 8}

	return writer
}

func (writer *wordWriter) close() {
	writer.write(endOfStream)
	writer.writeByte(writer.remainder)
	writer.writeByte(byte(0x00))
}

func (writer *wordWriter) write(value word) {
	remaining := wordSize
	for remaining >= writer.remainderFree {
		writer.remainder |= value.partFrom(wordSize-remaining, writer.remainderFree)
		remaining -= writer.remainderFree

		writer.writeByte(writer.remainder)
		writer.remainder = byte(0x00)
		writer.remainderFree = 8
	}
	writer.remainderFree = 8 - remaining
	writer.remainder = value.partFrom(wordSize-remaining, remaining) << writer.remainderFree
}

func (writer *wordWriter) writeByte(value byte) {
	writer.coder.CodeByte(&value)
}
