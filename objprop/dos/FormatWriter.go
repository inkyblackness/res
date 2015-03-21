package dos

import (
	"io"

	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/objprop"
	"github.com/inkyblackness/res/serial"
)

type formatWriter struct {
	coder serial.PositioningCoder
	dest  io.Writer

	entries map[res.ObjectID]*typeEntry
}

// NewConsumer wraps the provided WriteSeeker in a consumer for object properties.
func NewConsumer(dest io.WriteSeeker, descriptors []objprop.ClassDescriptor) objprop.Consumer {
	writer := &formatWriter{
		coder:   serial.NewPositioningEncoder(dest),
		dest:    dest,
		entries: calculateEntryValues(descriptors)}

	ensureCoderLength(writer.coder, expectedDataLength(descriptors))

	return writer
}

// Consume takes the provided data and associates it with the given ID.
func (writer *formatWriter) Consume(id res.ObjectID, data objprop.ObjectData) {
	entry := writer.entries[id]

	codeObjectData(writer.coder, entry, &data)
}

// Finish marks the end of consumption. After calling Finish, the consumer can't be used anymore.
func (writer *formatWriter) Finish() {
	header := MagicHeader

	writer.coder.SetCurPos(0)
	writer.coder.CodeUint32(&header)
}

func ensureCoderLength(coder serial.Coder, length uint32) {
	remaining := length
	zeroLen := uint32(1024)
	zero := make([]byte, zeroLen)

	for remaining > 0 {
		toCopy := zeroLen
		if toCopy > remaining {
			toCopy = remaining
		}
		coder.CodeBytes(zero[:toCopy])
		remaining -= toCopy
	}
}
