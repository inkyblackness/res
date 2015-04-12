package dos

import (
	"github.com/inkyblackness/res/serial"
	"github.com/inkyblackness/res/textprop"
)

type formatWriter struct {
	dest  serial.SeekingWriteCloser
	coder serial.Coder
}

// NewConsumer wraps the provided Writer in a consumer for text properties.
func NewConsumer(dest serial.SeekingWriteCloser) textprop.Consumer {
	writer := &formatWriter{dest: dest, coder: serial.NewEncoder(dest)}

	return writer
}

// Consume takes the provided data and adds it to the stream
func (writer *formatWriter) Consume(data []byte) {
	writer.coder.CodeBytes(data)
}

func (writer *formatWriter) Finish() {
	writer.dest.Close()
}
