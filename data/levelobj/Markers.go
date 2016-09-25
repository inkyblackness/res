package levelobj

import (
	"github.com/inkyblackness/res/data/interpreters"
	"github.com/inkyblackness/res/data/levelobj/actions"
	"github.com/inkyblackness/res/data/levelobj/conditions"
)

var baseMarker = interpreters.New()

var repulsor = baseMarker.
	With("StartHeight", 10, 4).
	With("EndHeight", 14, 4).
	With("Flags", 18, 4)

var baseTrigger = baseMarker.
	Refining("Action", 0, 22, actions.Unconditional(), interpreters.Always)

var gameVariableTrigger = baseTrigger.
	Refining("Condition", 2, 4, conditions.GameVariable(), interpreters.Always)

var deathWatchTrigger = baseTrigger.
	With("ConditionType", 5, 1).
	Refining("TypeCondition", 2, 4, conditions.ObjectType(), func(inst *interpreters.Instance) bool {
		return inst.Get("ConditionType") == 0
	}).
	Refining("IndexCondition", 2, 4, conditions.ObjectIndex(), func(inst *interpreters.Instance) bool {
		return inst.Get("ConditionType") == 1
	})

var ecologyTrigger = baseTrigger.
	Refining("TypeCondition", 2, 4, conditions.ObjectType(), interpreters.Always).
	With("ConditionLimit", 5, 1)

var mapNote = baseMarker.
	With("EntryOffset", 18, 4)

func initMarkers() interpreterRetriever {

	gameVariableTriggers := newInterpreterLeaf(gameVariableTrigger)

	trigger := newInterpreterEntry(baseMarker)
	trigger.set(0, gameVariableTriggers) // tile entry trigger
	trigger.set(1, gameVariableTriggers) // null trigger
	trigger.set(2, gameVariableTriggers) // floor trigger
	trigger.set(3, gameVariableTriggers) // player death trigger
	trigger.set(4, newInterpreterLeaf(deathWatchTrigger))
	trigger.set(8, gameVariableTriggers) // level entry trigger
	trigger.set(10, newInterpreterLeaf(repulsor))
	trigger.set(11, newInterpreterLeaf(ecologyTrigger))
	trigger.set(12, gameVariableTriggers) // shodan trigger

	mapMarker := newInterpreterEntry(baseMarker)
	mapMarker.set(3, newInterpreterLeaf(mapNote))

	class := newInterpreterEntry(baseMarker)
	class.set(0, trigger)
	class.set(2, mapMarker)

	return class
}
