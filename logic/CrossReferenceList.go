package logic

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/inkyblackness/res/data"
)

// CrossReferenceListIndex is an index into a cross reference list
type CrossReferenceListIndex uint16

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

	list.resetEntry(list.entry(0))
	for index := size - 1; index > 0; index-- {
		list.addEntryToAvailablePool(CrossReferenceListIndex(index))
	}
}

// Entry returns a pointer to the entry of given index.
func (list *CrossReferenceList) entry(index CrossReferenceListIndex) *data.LevelObjectCrossReference {
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

// AddObjectToMap adds an object to the map, at the specified locations.
// The returned value is the first cross-reference index to be stored in the specified object.
func (list *CrossReferenceList) AddObjectToMap(objectIndex uint16, tileMap TileMapReferencer,
	locations []TileLocation) (entryIndex CrossReferenceListIndex, err error) {
	affectedIndices := []CrossReferenceListIndex{}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)

			for _, index := range affectedIndices {
				list.addEntryToAvailablePool(index)
			}
		}
	}()

	startEntry := list.entry(0)

	for _, location := range locations {
		if startEntry.NextObjectIndex == 0 {
			panic(fmt.Errorf("Cross-Reference list is exhausted. Can not add more objects."))
		}
		oldTileIndex := tileMap.ReferenceIndex(location)
		newReferenceIndex := CrossReferenceListIndex(startEntry.NextObjectIndex)
		newEntry := list.entry(newReferenceIndex)

		startEntry.NextObjectIndex = newEntry.NextObjectIndex

		newEntry.NextObjectIndex = uint16(oldTileIndex)
		newEntry.LevelObjectTableIndex = objectIndex
		newEntry.NextTileIndex = uint16(entryIndex)
		newEntry.TileX, newEntry.TileY = location.XY()

		entryIndex = newReferenceIndex
		affectedIndices = append(affectedIndices, newReferenceIndex)
	}
	list.entry(affectedIndices[0]).NextTileIndex = uint16(entryIndex)
	for locationIndex, location := range locations {
		tileMap.SetReferenceIndex(location, affectedIndices[locationIndex])
	}

	return
}

func (list *CrossReferenceList) addEntryToAvailablePool(index CrossReferenceListIndex) {
	startEntry := list.entry(0)
	entry := list.entry(index)

	list.resetEntry(entry)
	entry.NextObjectIndex = startEntry.NextObjectIndex
	startEntry.NextObjectIndex = uint16(index)
}
