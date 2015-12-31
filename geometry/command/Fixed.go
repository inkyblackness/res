package command

import (
	"fmt"

	"github.com/inkyblackness/res/serial"
)

// Fixed is a value type for serialized coordinates.
type Fixed uint32

// ToFixed creates a Fixed out of a floating point value.
func ToFixed(value float32) Fixed {
	return Fixed(uint32(value * 256.0))
}

func CodeFixed(coder serial.Coder, coord *Fixed) {
	raw := uint32(*coord)
	coder.CodeUint32(&raw)
	*coord = Fixed(raw)
}

// Float returns the closest floating point value to the coordinate.
func (coord Fixed) Float() float32 {
	return float32(int32(coord)) / 256.0
}

// String returns the string presentation of the result of Float().
func (coord Fixed) String() string {
	return fmt.Sprintf("%v", coord.Float())
}
