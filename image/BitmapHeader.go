package image

import (
	"fmt"
)

// BitmapHeader describes the header of an encoded bitmap
type BitmapHeader struct {
	Unknown0000 [4]byte
	Type        int16
	Unknown0006 int16

	Width        int16
	Height       int16
	Stride       int16
	WidthFactor  byte
	HeightFactor byte
	HotspotBox   [4]int16

	PaletteOffset int32
}

func (header *BitmapHeader) String() (result string) {
	result += fmt.Sprintf("Type: 0x%02X, %dx%d\n", header.Type, header.Width, header.Height)
	result += fmt.Sprintf("0006: 0x%04X\n", header.Unknown0006)
	result += fmt.Sprintf("%d,%d | %d,%d\n", header.HotspotBox[0], header.HotspotBox[1], header.HotspotBox[2], header.HotspotBox[3])
	return
}
