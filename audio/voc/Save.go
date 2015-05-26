package voc

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Save encodes the provided samples into the given writer
func Save(writer io.Writer, sampleRate float32, samples []int16) {
	version := baseVersion
	sampleType := byte(0)
	writer.Write(bytes.NewBufferString(fileHeader).Bytes())
	binary.Write(writer, binary.LittleEndian, standardHeaderSize)
	binary.Write(writer, binary.LittleEndian, baseVersion)
	binary.Write(writer, binary.LittleEndian, uint16(uint16(^version)+versionCheckValue))

	dataBytes := len(samples) + 2
	writer.Write([]byte{byte(soundData), byte(dataBytes), byte(dataBytes >> 8), byte(dataBytes >> 16)})
	writer.Write([]byte{sampleRateToDivisor(sampleRate), sampleType})

	for _, sample := range samples {
		writer.Write([]byte{byte(uint16(int(sample)+0x8000) >> 8)})
	}

	writer.Write([]byte{byte(terminator)})
}
