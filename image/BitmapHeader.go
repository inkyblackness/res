package image

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

	Unknown0018 int32
}
