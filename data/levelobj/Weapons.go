package levelobj

import (
	"github.com/inkyblackness/res/data/interpreters"
)

var baseWeapon = interpreters.New()

var energyWeapon = baseWeapon.
	With("Charge", 0, 1).
	With("Temperature", 1, 1)

var projectileWeapon = baseWeapon.
	With("AmmoType", 0, 1).
	With("AmmoCount", 1, 1)

func initWeapons() interpreterRetriever {
	class := newInterpreterEntry(baseWeapon)
	class.set(1, newInterpreterLeaf(projectileWeapon))

	return class
}
