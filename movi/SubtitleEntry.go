package movi

// SubtitleEntrySize is the size, in bytes, of the entry structure
const SubtitleEntrySize = 16

// SubtitleEntry is the header structure of a subtitle data
type SubtitleEntry struct {
	// Control specifies how to interpret the string content
	Control SubtitleControl

	Unknown0004 byte
	Unknown0005 [11]byte
}
