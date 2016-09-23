package levelobj

import (
	"github.com/inkyblackness/res/data/interpreters"
	"github.com/inkyblackness/res/data/levelobj/actions"
	"github.com/inkyblackness/res/data/levelobj/conditions"
)

var baseMarker = interpreters.New()

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

func initMarkers() interpreterRetriever {

	gameVariableTriggers := newInterpreterLeaf(gameVariableTrigger)
	baseTriggers := newInterpreterLeaf(baseTrigger)

	trigger := newInterpreterEntry(baseMarker)
	trigger.set(0, gameVariableTriggers) // tile entry trigger
	trigger.set(1, gameVariableTriggers) // null trigger
	trigger.set(2, baseTriggers)         // floor trigger
	trigger.set(3, gameVariableTriggers) // player death trigger
	trigger.set(4, newInterpreterLeaf(deathWatchTrigger))
	trigger.set(8, gameVariableTriggers)  // level entry trigger
	trigger.set(12, gameVariableTriggers) // shodan trigger

	class := newInterpreterEntry(baseMarker)
	class.set(0, trigger)

	return class
}
