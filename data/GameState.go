package data

import (
	"bytes"
	"fmt"
)

type GameState struct {
	HackerName [20]byte

	Unknown0014 [1]byte

	CombatRating  byte
	MissionRating byte
	PuzzleRating  byte
	CyberRating   byte

	Unknown0019 [0x20]byte

	CurrentLevel byte

	Unknown001B [0x4E6]byte

	HackerX     TileCoordinate
	Unknown0522 [2]byte
	HackerY     TileCoordinate
}

func DefaultGameState() *GameState {
	return &GameState{}
}

func (data *GameState) String() string {
	info := ""

	info += fmt.Sprintf("Hacker Name: <%s>\n", data.HackerNameString())
	info += fmt.Sprintf("Ratings: Co: %d, Mi: %d, Pu: %d, Cy: %d\n",
		data.CombatRating, data.MissionRating, data.PuzzleRating, data.CyberRating)
	info += fmt.Sprintf("Current Level: %d\n", data.CurrentLevel)
	info += fmt.Sprintf("Hacker Coord: X: %v, Y: %v\n", data.HackerX, data.HackerY)

	return info
}

func (data *GameState) HackerNameString() string {
	buffer := bytes.NewBuffer(data.HackerName[:])

	return buffer.String()
}
