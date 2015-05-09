package data

import (
	"fmt"
)

const (
	defaultMapWidth     uint32 = 64
	defaultMapHeight    uint32 = 64
	defaultHeightFactor uint32 = 3
)

type LevelInformation struct {
	MapWidth  uint32
	MapHeight uint32

	Unknown0008 [8]byte

	HeightFactor uint32

	IgnoredPlaceholder uint32

	CyberspaceFlag uint32

	Unknown001C [30]byte
}

func DefaultLevelInformation() *LevelInformation {
	info := &LevelInformation{
		MapWidth:     defaultMapWidth,
		MapHeight:    defaultMapHeight,
		Unknown0008:  [8]byte{0x06, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00},
		HeightFactor: defaultHeightFactor}

	return info
}

func (info *LevelInformation) String() string {
	result := ""

	result += fmt.Sprintf("Map Dimension: %d x %d\n", info.MapWidth, info.MapHeight)
	result += fmt.Sprintf("Cyberspace: %v\n", info.IsCyberspace())

	return result
}

func (info *LevelInformation) IsCyberspace() bool {
	return info.CyberspaceFlag != 0
}
