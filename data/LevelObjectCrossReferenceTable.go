package data

import (
	"fmt"
)

type LevelObjectCrossReferenceTable struct {
	Entries []*LevelObjectCrossReference
}

func DefaultLevelObjectCrossReferenceTable(entryCount int) *LevelObjectCrossReferenceTable {
	table := &LevelObjectCrossReferenceTable{
		Entries: make([]*LevelObjectCrossReference, entryCount)}

	for i := range table.Entries {
		table.Entries[i] = DefaultLevelObjectCrossReference()
	}

	return table
}

func (table *LevelObjectCrossReferenceTable) String() (result string) {
	for i, entry := range table.Entries {
		result += fmt.Sprintf("Index %d:\n", i)
		result += entry.String()
		result += "\n"
	}

	return
}
