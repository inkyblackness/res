package data

// TileType describes the general type of a map tile.
type TileType byte

const (
	Solid TileType = 0x00
	Open           = 0x01

	DiagonalOpenSouthEast = 0x02
	DiagonalOpenSouthWest = 0x03
	DiagonalOpenNorthWest = 0x04
	DiagonalOpenNorthEast = 0x05

	SlopeSouthToNorth = 0x06
	SlopeWestToEast   = 0x07
	SlopeNorthToSouth = 0x08
	SlopeEastToWest   = 0x09

	ValleySouthEastToNorthWest = 0x0A
	ValleySouthWestToNorthEast = 0x0B
	ValleyNorthWestToSouthEast = 0x0C
	ValleyNorthEastToSouthWest = 0x0D

	RidgeNorthWestToSouthEast = 0x0E
	RidgeNorthEastToSouthWest = 0x0F
	RidgeSouthEastToNorthWest = 0x10
	RidgeSouthWestToNorthEast = 0x11
)
