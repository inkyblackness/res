package dos

import (
	"io"

	"github.com/inkyblackness/res/serial"
	"github.com/inkyblackness/res/textprop"
)

type formatWriter struct {
	coder serial.Coder
}

// NewConsumer wraps the provided Writer in a consumer for text properties.
func NewConsumer(dest io.Writer) textprop.Consumer {
	writer := &formatWriter{coder: serial.NewEncoder(dest)}

	return writer
}

// Consume takes the provided data and adds it to the stream
func (writer *formatWriter) Consume(data []byte) {
	writer.coder.CodeBytes(data)
}
