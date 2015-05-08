package data

type TileMap struct {
	// The entries of the map. First columns (left to right), then rows (bottom to top).
	Entries []*TileMapEntry
}

func DefaultTileMap(width, height int) *TileMap {
	tiles := width * height
	tileMap := &TileMap{Entries: make([]*TileMapEntry, tiles)}

	for i := range tileMap.Entries {
		tileMap.Entries[i] = DefaultTileMapEntry()
	}

	return tileMap
}
