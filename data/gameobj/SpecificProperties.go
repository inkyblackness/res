package gameobj

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/data/interpreters"
)

// SpecificProperties returns an interpreter specific for the given object class and subclass.
func SpecificProperties(objClass res.ObjectClass, objSubclass res.ObjectSubclass, data []byte) *interpreters.Instance {
	objID := res.MakeObjectID(objClass, objSubclass, 0)
	desc := specificDescriptions[objID]
	if desc == nil {
		desc = interpreters.New()
	}

	return desc.For(data)
}
