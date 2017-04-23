package data

import (
	"fmt"
)

// TileFlag describes simple properties of a map tile.
type TileFlag uint32

const (
	// TileVisited specifies whether the tile has been seen.
	TileVisited TileFlag = 0x80000000

	// MusicIndexTileFlagMask is the mask for the music identifier.
	MusicIndexTileFlagMask = 0x0000F000
)

// MusicIndex returns the music identifier for the tile.
func (flag TileFlag) MusicIndex() int {
	return int((flag & MusicIndexTileFlagMask) >> 12)
}

// WithMusicIndex returns a new flag combination having given music identifier.
func (flag TileFlag) WithMusicIndex(value int) TileFlag {
	cleared := uint32(flag) & ^uint32(MusicIndexTileFlagMask)
	newValue := (uint32(value) << 12) & MusicIndexTileFlagMask
	return TileFlag(cleared | newValue)
}

func (flag TileFlag) String() (result string) {
	texts := []string{}
	if (flag & TileVisited) != 0 {
		texts = append(texts, "TileVisited")
	}
	texts = append(texts, fmt.Sprintf("MusicIndex=%v", flag.MusicIndex()))

	for _, text := range texts {
		if len(result) > 0 {
			result += "|"
		}
		result += text
	}

	return
}
