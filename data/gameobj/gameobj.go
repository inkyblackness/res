package gameobj

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/data/interpreters"
)

var genericDescriptions map[res.ObjectClass]*interpreters.Description
var specificDescriptions map[res.ObjectID]*interpreters.Description
var anyType = res.ObjectType(0xFF)

func init() {
	genericDescriptions = make(map[res.ObjectClass]*interpreters.Description)
	specificDescriptions = make(map[res.ObjectID]*interpreters.Description)

	initWeapons()
	initAmmoClips()
	initProjectiles()
	initExplosives()
	initItems()
}

func setSpecific(objClass res.ObjectClass, objSubclass int, desc *interpreters.Description) {
	specificDescriptions[res.MakeObjectID(objClass, res.ObjectSubclass(objSubclass), anyType)] = desc
}

func setSpecificByType(objClass res.ObjectClass, objSubclass int, objType int, desc *interpreters.Description) {
	specificDescriptions[res.MakeObjectID(objClass, res.ObjectSubclass(objSubclass), res.ObjectType(objType))] = desc
}
