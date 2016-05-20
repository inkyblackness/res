package logic

// TileMapReferencer is an interface for keeping cross-references in a tile map.
type TileMapReferencer interface {
	// ReferenceIndex returns the index of the first cross-reference for the given tile.
	ReferenceIndex(tileX, tileY int) int
	// SetReferenceIndex sets the index of the first cross-reference for the given tile.
	SetReferenceIndex(tileX, tileY int, index int)
}
