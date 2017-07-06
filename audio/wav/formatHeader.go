package wav

import (
	"encoding/binary"
)

type waveFormatType uint16

const (
	waveFormatTypePcm = 1
)

type waveFormat struct {
	formatType     waveFormatType
	channels       uint16
	samplesPerSec  uint32
	avgBytesPerSec uint32
	blockAlign     uint16
}

type waveFormatExtension struct {
	bitsPerSample uint16
	extensionSize uint16
}

type formatHeader struct {
	base      waveFormat
	extension waveFormatExtension
}

func (header formatHeader) size() uint32 {
	return uint32(binary.Size(&header.base) + binary.Size(&header.extension))
}
