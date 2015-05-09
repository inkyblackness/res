package data

import (
	"fmt"
)

type TileCoordinate uint16

func (coord TileCoordinate) String() (result string) {
	result += fmt.Sprintf("%d:%d", coord.Tile(), coord.Offset())

	return
}

func (coord TileCoordinate) Tile() byte {
	return byte(uint16(coord) >> 8)
}

func (coord TileCoordinate) Offset() byte {
	return byte(coord & 0x00FF)
}
