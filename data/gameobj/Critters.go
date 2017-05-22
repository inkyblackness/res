package gameobj

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/data/interpreters"
)

var critterAttackInfo = interpreters.New().
	With("Damage", 0x0004, 2).As(interpreters.RangedValue(0, 500)).
	With("OffenceValue", 0x0006, 1).
	With("HitChance", 0x000B, 1).
	With("AttackRange", 0x000C, 1).
	With("AttackDelay", 0x000E, 1).
	With("ProjectileType", 0x0011, 4).As(interpreters.SpecialValue("ObjectType"))

var critterGenerics = interpreters.New().
	Refining("PrimaryAttack", 0x0000, 21, critterAttackInfo, interpreters.Always).
	Refining("SecondaryAttack", 0x0015, 21, critterAttackInfo, interpreters.Always).
	With("ProjectileSourceHeightOffset", 0x002C, 1).As(interpreters.RangedValue(-128, 127)).
	With("Flags", 0x002D, 1).
	With("Unknown0031", 0x031, 1).As(interpreters.SpecialValue("Unknown")).
	With("FrameTime", 0x003A, 1).
	With("AttackSoundIndex", 0x003B, 1).
	With("IdleSoundIndex", 0x003C, 1).
	With("PainSoundIndex", 0x003D, 1).
	With("DeathSoundIndex", 0x003E, 1).
	With("HostileSoundIndex", 0x003F, 1).
	With("CorpseType", 0x0040, 4).As(interpreters.SpecialValue("ObjectType")).
	With("FrameCount", 0x0044, 1).
	With("SecondaryAttackProbability", 0x0045, 1).
	With("InterruptProbability", 0x0046, 1).
	With("RandomLootSelection", 0x0047, 1).As(interpreters.RangedValue(0, 14)).
	With("InjuryType", 0x0048, 1).As(interpreters.EnumValue(map[uint32]string{0: "meat", 1: "plant", 2: "metal", 3: "cyborg meat"}))

var cyberCritters = interpreters.New().
	Refining("ColorScheme", 0, 6, cyberColorScheme, interpreters.Always)

func initCritters() {
	objClass := res.ObjectClass(14)

	genericDescriptions[objClass] = critterGenerics

	setSpecific(objClass, 3, cyberCritters)
}
