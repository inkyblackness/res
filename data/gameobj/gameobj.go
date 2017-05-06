package gameobj

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/data/interpreters"
)

var genericDescriptions map[res.ObjectClass]*interpreters.Description
var specificDescriptions map[res.ObjectID]*interpreters.Description

func init() {
	genericDescriptions = make(map[res.ObjectClass]*interpreters.Description)
	specificDescriptions = make(map[res.ObjectID]*interpreters.Description)

	initWeapons()
}

func setSpecific(objClass res.ObjectClass, objSubclass int, desc *interpreters.Description) {
	specificDescriptions[res.MakeObjectID(objClass, res.ObjectSubclass(objSubclass), 0)] = desc
}
