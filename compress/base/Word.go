package base

type word uint16

const (
	bitsPerWord = uint(14)

	endOfStream  = word(0x3FFF)
	reset        = word(0x3FFE)
	literalLimit = word(0x0100)
)

func (value word) partFrom(fromBit uint, count uint) byte {
	remaining := bitsPerWord - fromBit - count

	return byte(uint(value)>>remaining) & ((1 << count) - 1)
}
