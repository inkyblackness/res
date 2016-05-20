package logic

import (
	"bytes"
	"encoding/binary"

	"github.com/inkyblackness/res/data"
)

// CrossReferenceList provides the logic for handling the cross-reference table
// in level archives.
type CrossReferenceList struct {
	references []data.LevelObjectCrossReference
}

// NewCrossReferenceList returns a new instance of an uninitialized list.
// The size of the list defaults to the expected size of the cross-reference table.
func NewCrossReferenceList() *CrossReferenceList {
	references := new([1600]data.LevelObjectCrossReference)
	list := &CrossReferenceList{
		references: references[:]}

	return list
}

// Size returns the count of entries in the list.
func (list *CrossReferenceList) size() int {
	return len(list.references)
}

// Encode serializes the list into a bytestream.
func (list *CrossReferenceList) Encode() []byte {
	buf := bytes.NewBuffer(nil)

	binary.Write(buf, binary.LittleEndian, list.references)

	return buf.Bytes()
}

// Clear resets the list to an initial state, returning all references to the
// pool of available entries.
func (list *CrossReferenceList) Clear() {
	size := list.size()
	for index := 0; index < size; index++ {
		entry := list.entry(index)

		list.resetEntry(entry)
		entry.NextObjectIndex = uint16((index + 1) % size)
	}
}

// Entry returns a pointer to the entry of given index.
func (list *CrossReferenceList) entry(index int) *data.LevelObjectCrossReference {
	return &list.references[index]
}

// ResetEntry clears all fields of the given entry.
func (list *CrossReferenceList) resetEntry(entry *data.LevelObjectCrossReference) {
	entry.LevelObjectTableIndex = 0
	entry.NextTileIndex = 0
	entry.TileX = 0xFFFF
	entry.TileY = 0xFFFF
	entry.NextObjectIndex = 0
}

// AddObjectToMap adds an object to the map.
// The returned value is the first cross-reference index to be stored in the specified object.
func (list *CrossReferenceList) AddObjectToMap(objectIndex uint16, tileMap TileMapReferencer) uint16 {
	return 0
}
