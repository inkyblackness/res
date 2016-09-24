package levelobj

import (
	"github.com/inkyblackness/res/data/interpreters"
	"github.com/inkyblackness/res/data/levelobj/actions"
	"github.com/inkyblackness/res/data/levelobj/conditions"
)

var basePanel = interpreters.New()

var gameVariablePanel = basePanel.
	Refining("Condition", 2, 4, conditions.GameVariable(), interpreters.Always)

var buttonPanel = gameVariablePanel.
	Refining("Action", 0, 22, actions.Unconditional(), interpreters.Always).
	With("AccessMask", 22, 2)

var recepticlePanel = basePanel

var standardRecepticle = recepticlePanel.
	Refining("Action", 0, 22, actions.Unconditional(), interpreters.Always).
	Refining("Condition", 2, 4, conditions.ObjectType(), interpreters.Always)

var antennaRelayPanel = recepticlePanel.
	With("TriggerObjectIndex1", 6, 2).
	With("TriggerObjectIndex2", 10, 2).
	With("DestroyObjectIndex", 14, 2)

var retinalIDScanner = recepticlePanel.
	Refining("Action", 0, 22, actions.Unconditional(), interpreters.Always)

var cyberspaceTerminal = gameVariablePanel.
	With("State", 0, 1).
	With("TargetX", 6, 4).
	With("TargetY", 10, 4).
	With("TargetZ", 14, 4).
	With("TargetLevel", 18, 4)

var energyChargeStation = gameVariablePanel.
	With("EnergyDelta", 6, 4).
	With("RechargeTime", 10, 4).
	With("TriggerObjectIndex", 14, 4).
	With("RechargedTimestamp", 18, 4)

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

	standardRecepticles := newInterpreterLeaf(standardRecepticle)
	antennaRelays := newInterpreterLeaf(antennaRelayPanel)
	recepticles := newInterpreterEntry(recepticlePanel)
	recepticles.set(0, standardRecepticles)
	recepticles.set(1, standardRecepticles)
	recepticles.set(2, standardRecepticles)
	recepticles.set(3, antennaRelays) // standard panel
	recepticles.set(4, antennaRelays) // plastiqued
	recepticles.set(6, newInterpreterLeaf(retinalIDScanner))

	stations := newInterpreterEntry(basePanel)
	stations.set(0, newInterpreterLeaf(cyberspaceTerminal))
	stations.set(1, newInterpreterLeaf(energyChargeStation))

	inputPanels := newInterpreterEntry(inputPanel)

	cyberspaceSwitches := newInterpreterEntry(basePanel)
	cyberspaceSwitches.set(0, newInterpreterLeaf(inactiveCyberspaceSwitch))

	class := newInterpreterEntry(basePanel)
	class.set(0, newInterpreterLeaf(buttonPanel))
	class.set(1, recepticles)
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
