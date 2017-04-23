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
	Refining("TypeCondition", 2, 4, conditions.ObjectType(), interpreters.Always)

var antennaRelayPanel = recepticlePanel.
	With("TriggerObjectIndex1", 6, 2).As(interpreters.ObjectIndex()).
	With("TriggerObjectIndex2", 10, 2).As(interpreters.ObjectIndex()).
	With("DestroyObjectIndex", 14, 2).As(interpreters.ObjectIndex())

var retinalIDScanner = recepticlePanel.
	Refining("Action", 0, 22, actions.Unconditional(), interpreters.Always)

var cyberspaceTerminal = gameVariablePanel.
	With("State", 0, 1).As(interpreters.EnumValue(map[uint32]string{0: "Off", 1: "Active", 2: "Locked"})).
	With("TargetX", 6, 4).As(interpreters.RangedValue(1, 63)).
	With("TargetY", 10, 4).As(interpreters.RangedValue(1, 63)).
	With("TargetZ", 14, 4).As(interpreters.RangedValue(0, 255)).
	With("TargetLevel", 18, 4).As(interpreters.EnumValue(map[uint32]string{10: "10", 14: "14", 15: "15"}))

var energyChargeStation = gameVariablePanel.
	With("EnergyDelta", 6, 4).As(interpreters.RangedValue(0, 255)).
	With("RechargeTime", 10, 4).As(interpreters.RangedValue(0, 3600)).
	With("TriggerObjectIndex", 14, 4).As(interpreters.ObjectIndex()).
	With("RechargedTimestamp", 18, 4)

var inputPanel = gameVariablePanel

var wirePuzzleData = interpreters.New().
	With("TargetObjectIndex", 0, 4).As(interpreters.ObjectIndex()).
	With("Layout", 4, 1).
	With("TargetPowerLevel", 5, 1).
	With("CurrentPowerLevel", 6, 1).
	With("TargetState", 8, 4).
	With("CurrentState", 12, 4)

var blockPuzzleData = interpreters.New().
	With("TargetObjectIndex", 0, 4).As(interpreters.ObjectIndex()).
	With("StateStoreObjectIndex", 4, 2).As(interpreters.ObjectIndex()).
	With("Layout", 8, 4)

var puzzleSpecificData = interpreters.New().
	With("Type", 7, 1).As(interpreters.EnumValue(map[uint32]string{0: "WirePuzzle", 0x10: "BlockPuzzle"})).
	Refining("Wire", 0, 18, wirePuzzleData, func(inst *interpreters.Instance) bool {
		return inst.Get("Type") == 0
	}).
	Refining("Block", 0, 18, blockPuzzleData, func(inst *interpreters.Instance) bool {
		return inst.Get("Type") == 0x10
	})

var puzzlePanel = inputPanel.
	Refining("Puzzle", 6, 18, puzzleSpecificData, interpreters.Always)

var elevatorPanel = inputPanel.
	With("DestinationObjectIndex2", 6, 2).As(interpreters.RangedValue(0, 871)).
	With("DestinationObjectIndex1", 8, 2).As(interpreters.RangedValue(0, 871)).
	With("DestinationObjectIndex4", 10, 2).As(interpreters.RangedValue(0, 871)).
	With("DestinationObjectIndex3", 12, 2).As(interpreters.RangedValue(0, 871)).
	With("DestinationObjectIndex6", 14, 2).As(interpreters.RangedValue(0, 871)).
	With("DestinationObjectIndex5", 16, 2).As(interpreters.RangedValue(0, 871)).
	With("AccessibleBitmask", 18, 2).
	With("ElevatorShaftBitmask", 20, 2)

var numberPad = inputPanel.
	With("Combination1", 6, 2).As(interpreters.SpecialValue("BinaryCodedDecimal")).
	With("TriggerObjectIndex1", 8, 2).As(interpreters.ObjectIndex()).
	With("Combination2", 10, 2).As(interpreters.SpecialValue("BinaryCodedDecimal")).
	With("TriggerObjectIndex2", 12, 2).As(interpreters.ObjectIndex()).
	With("Combination3", 14, 2).As(interpreters.SpecialValue("BinaryCodedDecimal")).
	With("TriggerObjectIndex3", 16, 2).As(interpreters.ObjectIndex()).
	With("FailObjectIndex", 18, 2).As(interpreters.ObjectIndex())

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

	puzzles := newInterpreterLeaf(puzzlePanel)
	elevatorPanels := newInterpreterLeaf(elevatorPanel)
	numberPads := newInterpreterLeaf(numberPad)
	inputPanels := newInterpreterEntry(inputPanel)
	inputPanels.set(0, puzzles)
	inputPanels.set(1, puzzles)
	inputPanels.set(2, puzzles)
	inputPanels.set(3, puzzles)
	inputPanels.set(4, elevatorPanels)
	inputPanels.set(5, elevatorPanels)
	inputPanels.set(6, elevatorPanels)
	inputPanels.set(7, numberPads)
	inputPanels.set(8, numberPads)
	inputPanels.set(9, puzzles)
	inputPanels.set(10, puzzles)

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