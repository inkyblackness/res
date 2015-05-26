package mem

// L8ToL16Table returns a lookup table with 256 entries. For an L8 sample as the key,
// the value is the corresponding L16 sample value.
// This function returns a new array.
func L8ToL16Table() []int16 {
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
