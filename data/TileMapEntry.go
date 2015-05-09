package data

import (
	"fmt"
)

type TileMapEntry struct {
	Type             TileType
	Floor            FloorInfo
	Ceiling          CeilingInfo
	SlopeHeight      byte
	FirstObjectIndex uint16
	Textures         TileTextureInfo
	Flags            TileFlag
	UnknownState     [4]byte
}

func DefaultTileMapEntry() *TileMapEntry {
	entry := &TileMapEntry{
		Type:         Solid,
		UnknownState: [4]byte{0xFF, 0x00, 0x00, 0x00}}

	return entry
}

func (entry *TileMapEntry) String() (result string) {
	result += fmt.Sprintf("Type: %v\n", entry.Type)
	result += fmt.Sprintf("First Object Index: %d\n", entry.FirstObjectIndex)
	result += fmt.Sprintf("Flags: %v\n", entry.Flags)

	return
}

type FloorInfo byte

type CeilingInfo byte
