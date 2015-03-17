package base

type dictEntry struct {
	prev  *dictEntry
	depth int

	value byte
	key   word

	next map[byte]*dictEntry
}

func rootDictEntry() *dictEntry {
	return &dictEntry{prev: nil, depth: 0, value: 0x00, key: reset, next: make(map[byte]*dictEntry)}
}

func (entry *dictEntry) Add(value byte, key word) *dictEntry {
	newEntry := &dictEntry{
		prev:  entry,
		depth: entry.depth + 1,
		value: value,
		key:   key,
		next:  make(map[byte]*dictEntry)}
	entry.next[value] = newEntry

	return newEntry
}

func (entry *dictEntry) Data() []byte {
	bytes := make([]byte, entry.depth, entry.depth)
	cur := entry
	for i := entry.depth - 1; i >= 0; i-- {
		bytes[i] = cur.value
		cur = cur.prev
	}

	return bytes
}

func (entry *dictEntry) FirstByte() byte {
	cur := entry
	for cur.depth != 1 {
		cur = entry.prev
	}

	return cur.value
}
