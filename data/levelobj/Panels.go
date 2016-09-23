package levelobj

import (
	"github.com/inkyblackness/res/data/interpreters"
	"github.com/inkyblackness/res/data/levelobj/actions"
	"github.com/inkyblackness/res/data/levelobj/conditions"
)

var basePanel = interpreters.New()

var gameVariablePanel = basePanel.
	//Refining(conditions.PlaceholderName, 2, 4, conditions.Placeholder(), interpreters.Never).
	Refining("Condition", 2, 4, conditions.GameVariable(), interpreters.Always)

var buttonPanel = gameVariablePanel.
	Refining("Action", 0, 22, actions.Unconditional(), interpreters.Always)

var cyberspaceTerminal = gameVariablePanel.
	With("State", 0, 1).
	With("TargetX", 6, 4).
	With("TargetY", 10, 4).
	With("TargetZ", 14, 4).
	With("TargetLevel", 18, 4)

var energyChargeStation = gameVariablePanel

var inputPanel = gameVariablePanel

var wirePuzzleData = interpreters.New().
	With("TargetObjectIndex", 0, 4).
	With("Layout", 4, 1).
	With("TargetPowerLevel", 5, 1).
	With("CurrentPowerLevel", 6, 1).
	With("TargetState", 8, 4).
	With("CurrentState", 12, 4)

var blockPuzzleData = interpreters.New().
	With("TargetObjectIndex", 0, 4).
	With("StateStoreObjectIndex", 4, 2).
	With("Layout", 8, 4)

var puzzleSpecificData = interpreters.New().
	With("Type", 7, 1).
	Refining("Wire", 0, 18, wirePuzzleData, func(inst *interpreters.Instance) bool {
		return inst.Get("Type") == 0
	}).
	Refining("Block", 0, 18, blockPuzzleData, func(inst *interpreters.Instance) bool {
		return inst.Get("Type") == 0x10
	})

var puzzlePanel = inputPanel.
	Refining("Puzzle", 6, 18, puzzleSpecificData, interpreters.Always)

var inactiveCyberspaceSwitch = gameVariablePanel.
	Refining("Action", 0, 22, actions.Unconditional(), interpreters.Always)

func initPanels() interpreterRetriever {

	stations := newInterpreterEntry(basePanel)
	stations.set(0, newInterpreterLeaf(cyberspaceTerminal))
	stations.set(1, newInterpreterLeaf(energyChargeStation))

	inputPanels := newInterpreterEntry(inputPanel)

	cyberspaceSwitches := newInterpreterEntry(basePanel)
	cyberspaceSwitches.set(0, newInterpreterLeaf(inactiveCyberspaceSwitch))

	class := newInterpreterEntry(basePanel)
	class.set(0, newInterpreterLeaf(buttonPanel))
	class.set(2, stations)
	class.set(3, inputPanels)
	class.set(5, cyberspaceSwitches)

	return class
}

func initCyberspacePanels() interpreterRetriever {

	cyberspaceSwitches := newInterpreterEntry(basePanel)
	cyberspaceSwitches.set(0, newInterpreterLeaf(inactiveCyberspaceSwitch))

	class := newInterpreterEntry(basePanel)
	class.set(5, cyberspaceSwitches)

	return class
}
