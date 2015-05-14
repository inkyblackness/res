package data

// TileFlag describes simple properties of a map tile.
type TileFlag uint32

const (
	// TileVisited specifies whether the tile has been seen.
	TileVisited TileFlag = 0x80000000
)
