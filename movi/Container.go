package movi

import (
	"image/color"
)

// Container wraps the information and data of a MOVI container.
type Container interface {
	VideoWidth() uint16
	VideoHeight() uint16
	StartPalette() color.Palette

	AudioSampleRate() uint16

	EntryCount() int
	Entry(index int) Entry
}
