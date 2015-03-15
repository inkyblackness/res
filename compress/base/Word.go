package base

type word uint16

const (
	wordSize = uint(14)

	endOfStream = word(0x3FFF)
	reset       = word(0x3FFE)
)

func (value word) partFrom(fromBit uint, count uint) byte {
	remaining := wordSize - fromBit - count

	return byte(uint(value)>>remaining) & ((1 << count) - 1)
}
