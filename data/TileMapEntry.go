package data

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

type FloorInfo byte

type CeilingInfo byte
