package data

const LevelItemEntrySize int = LevelObjectPrefixSize + 10

type LevelItemEntry struct {
	LevelObjectPrefix

	Unknown [10]byte
}

func NewLevelItemEntry() *LevelItemEntry {
	return &LevelItemEntry{}
}

func (entry *LevelItemEntry) String() (result string) {
	result += entry.LevelObjectPrefix.String()

	return
}
