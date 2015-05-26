package voc

var byteLookupTable = makeByteLookupTable()

func makeByteLookupTable() []int16 {
	table := make([]int16, 256)
	index := 0

	for i := 0x80; i > 0; i-- {
		table[index] = int16(-(i << 8) - ((i * 2) & 0xFF))
		index++
	}

	table[index] = 0
	index++

	for i := 1; i < 0x80; i++ {
		table[index] = int16(i)<<8 + int16(i*2+1)
		index++
	}

	return table
}
