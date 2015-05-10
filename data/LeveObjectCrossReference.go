package data

import (
	"fmt"
)

const LevelObjectCrossReferenceSize int = 10

type LevelObjectCrossReference struct {
	TileX uint16
	TileY uint16

	LevelObjectTableIndex uint16

	NextObjectIndex uint16
	NextTileIndex   uint16
}

func DefaultLevelObjectCrossReference() *LevelObjectCrossReference {
	return &LevelObjectCrossReference{}
}

func (ref *LevelObjectCrossReference) String() (result string) {
	result += fmt.Sprintf("Coord: X: %v Y: %v\n", ref.TileX, ref.TileY)
	result += fmt.Sprintf("Level Object Table Index: %d\n", ref.LevelObjectTableIndex)
	result += fmt.Sprintf("Next Object Index: %d\n", ref.NextObjectIndex)
	result += fmt.Sprintf("Next Tile Index: %d\n", ref.NextTileIndex)

	return
}
