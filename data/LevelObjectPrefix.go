package data

import (
	"fmt"
)

const LevelObjectPrefixSize int = 6

type LevelObjectPrefix struct {
	LevelObjectTableIndex uint16
	Previous              uint16
	Next                  uint16
}

func (prefix *LevelObjectPrefix) String() (result string) {
	result += fmt.Sprintf("Level object table index: %d\n", prefix.LevelObjectTableIndex)
	result += fmt.Sprintf("Links: <- %d | %d ->\n", prefix.Previous, prefix.Next)

	return
}
