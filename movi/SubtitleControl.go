package movi

// SubtitleControl specifies how to interpret a subtitle entry.
type SubtitleControl uint32

const (
	SubtitleArea    = uint32(0x41455241)
	SubtitleTextStd = uint32(0x20445453)
	SubtitleTextFrn = uint32(0x204E5246)
	SubtitleTextGer = uint32(0x20524547)
)

// String returns the string presentation of the control value.
func (ctrl SubtitleControl) String() string {
	return string([]rune{rune((ctrl >> 0) & 0xFF), rune((ctrl >> 8) & 0xFF), rune((ctrl >> 16) & 0xFF), rune((ctrl >> 24) & 0xFF)})
}
