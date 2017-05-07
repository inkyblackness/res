package gameobj

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/data/interpreters"
)

var explosiveGenerics = interpreters.New().
	Refining("BasicWeapon", 0, 8, basicWeapon, interpreters.Always)

func initExplosives() {
	objClass := res.ObjectClass(3)

	genericDescriptions[objClass] = explosiveGenerics
}
