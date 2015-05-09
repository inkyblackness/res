package data

import (
	"fmt"

	"github.com/inkyblackness/res"
)

const LevelObjectEntrySize int = 27

type LevelObjectEntry struct {
	InUse    byte
	Class    res.ObjectClass
	Subclass res.ObjectSubclass

	ClassTableIndex          uint16
	CrossReferenceTableIndex uint16
	Previous                 uint16
	Next                     uint16

	X    TileCoordinate
	Y    TileCoordinate
	Z    byte
	Rot1 byte
	Rot2 byte
	Rot3 byte

	Unknown0013 [1]byte

	Type res.ObjectType

	Unknown0015 [2]byte

	Unknown0018 [4]byte
}

func DefaultLevelObjectEntry() *LevelObjectEntry {
	return &LevelObjectEntry{}
}

func (entry *LevelObjectEntry) String() (result string) {
	result += fmt.Sprintf("In Use: %v\n", entry.IsInUse())
	result += fmt.Sprintf("ObjectID: %d/%d/%d\n", entry.Class, entry.Subclass, entry.Type)
	result += fmt.Sprintf("Coord: X: %v Y: %v Z: %d\n", entry.X, entry.Y, entry.Z)
	result += fmt.Sprintf("Rotation: %d, %d, %d\n", entry.Rot1, entry.Rot2, entry.Rot3)

	return
}

func (entry *LevelObjectEntry) IsInUse() bool {
	return entry.InUse != 0
}
