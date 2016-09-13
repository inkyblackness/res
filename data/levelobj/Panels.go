package levelobj

import (
	"github.com/inkyblackness/res/data/interpreters"
)

var commonPanelData = interpreters.New()

var basePanel = interpreters.New().
	Refining("Common", 0, 6, commonPanelData, interpreters.Always)

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

var puzzlePanel = basePanel.
	Refining("Puzzle", 6, 18, puzzleSpecificData, interpreters.Always)
