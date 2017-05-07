package gameobj

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/data/interpreters"
)

var cyberProjectiles = interpreters.New().
	Refining("ColorScheme", 0, 6, cyberColorScheme, interpreters.Always)

func initProjectiles() {
	objClass := res.ObjectClass(2)

	setSpecificByType(objClass, 1, 9, cyberProjectiles)
	setSpecificByType(objClass, 1, 10, cyberProjectiles)
	setSpecificByType(objClass, 1, 11, cyberProjectiles)
	setSpecificByType(objClass, 1, 12, cyberProjectiles)
	setSpecificByType(objClass, 1, 13, cyberProjectiles)
}
