package voc

func lengthFromBlockStart(blockStart []byte) int {
	return int(blockStart[3])<<16 + int(blockStart[2])<<8 + int(blockStart[1])
}

func divisorToSampleRate(divisor byte) float32 {
	return 1000000.0 / float32(256-int(divisor))
}
