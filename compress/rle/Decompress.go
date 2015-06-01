package rle

import (
	"bytes"
	"fmt"
	"io"
)

// Decompress decompresses the given array and returns the result.
func Decompress(reader io.Reader, expectedSize int) (uncompressed []byte, err error) {
	writer := bytes.NewBuffer(make([]byte, 0, expectedSize))
	endOfStream := false
	nextByte := func() byte {
		zz := []byte{0x00}
		_, err = reader.Read(zz)
		return zz[0]
	}

	for err == nil && !endOfStream {
		first := nextByte()

		if first == 0x00 {
			nn := nextByte()
			zz := nextByte()

			writeBytesOfValue(writer, int(nn), func() byte { return zz })
		} else if first < 0x80 {
			writeBytesOfValue(writer, int(first), nextByte)
		} else if first == 0x80 {
			control := uint16(nextByte()) + (uint16(nextByte()) << 8)
			if control == 0x0000 {
				writeBytesOfValue(writer, expectedSize-writer.Len(), func() byte { return 0x00 })
				endOfStream = true
			} else if control < 0x8000 {
				writeBytesOfValue(writer, int(control), func() byte { return 0x00 })
			} else if control < 0xC000 {
				writeBytesOfValue(writer, int(control&0x3FFF), nextByte)
			} else if (control & 0xFF00) == 0xC000 {
				err = fmt.Errorf("Undefined case 80 nn C0")
			} else {
				zz := nextByte()

				writeBytesOfValue(writer, int(control&0x3FFF), func() byte { return zz })
			}
		} else {
			writeBytesOfValue(writer, int(first&0x7F), func() byte { return 0x00 })
		}
	}
	uncompressed = writer.Bytes()

	return
}

func writeBytesOfValue(writer *bytes.Buffer, nn int, producer func() byte) {
	for i := 0; i < nn; i++ {
		writer.WriteByte(producer())
	}
}
