package data

import (
	"fmt"
)

const LevelWeaponEntrySize int = LevelObjectPrefixSize + 2

type LevelWeaponEntry struct {
	LevelObjectPrefix

	AmmoTypeOrCharge       byte
	AmmoCountOrTemperature byte
}

func NewLevelWeaponEntry() *LevelWeaponEntry {
	return &LevelWeaponEntry{}
}

func (entry *LevelWeaponEntry) String() (result string) {
	result += entry.LevelObjectPrefix.String()
	result += fmt.Sprintf("Ammo Type/Charge: %d\n", entry.AmmoTypeOrCharge)
	result += fmt.Sprintf("Ammo Count/Temp: %d\n", entry.AmmoCountOrTemperature)

	return
}
