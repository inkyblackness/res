package data

import (
	"fmt"
)

type LevelObjectTable struct {
	Entries []*LevelObjectEntry
}

func DefaultLevelObjectTable(entryCount int) *LevelObjectTable {
	table := &LevelObjectTable{
		Entries: make([]*LevelObjectEntry, entryCount)}

	for i := range table.Entries {
		table.Entries[i] = DefaultLevelObjectEntry()
	}

	return table
}

func (table *LevelObjectTable) String() (result string) {
	for i, entry := range table.Entries {
		result += fmt.Sprintf("Index %d:\n", i)
		result += entry.String()
		result += "\n"
	}

	return
}
