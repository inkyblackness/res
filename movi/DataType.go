package movi

// DataType identifies entries
type DataType byte

const (
	// endOfMedia marks the last entry
	endOfMedia = DataType(0)
	// Audio marks an audio entry.
	Audio = DataType(2)
)
