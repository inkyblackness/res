package data

import (
	"fmt"
)

type Table struct {
	Entries []interface{}
}

func NewTable(entryCount int, factory func() interface{}) *Table {
	table := &Table{
		Entries: make([]interface{}, entryCount)}

	for i := range table.Entries {
		table.Entries[i] = factory()
	}

	return table
}

func (table *Table) String() (result string) {
	for i, entry := range table.Entries {
		result += fmt.Sprintf("Index %d:\n", i)
		result += fmt.Sprintf("%v\n", entry)
	}

	return
}
