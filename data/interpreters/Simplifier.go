package interpreters

import (
	"math"
)

// RawValueHandler is for a simple value range.
type RawValueHandler func(minValue, maxValue int64)

// EnumValueHandler is for enumerated (mapped) values.
type EnumValueHandler func(values map[uint32]string)

// EnumValue creates a field range describing enumerated values.
func EnumValue(values map[uint32]string) FieldRange {
	return func(simpl *Simplifier) bool {
		return simpl.enumValue(values)
	}
}

// Simplifier forwards descriptions in a way the requester can use.
type Simplifier struct {
	rawValueHandler  RawValueHandler
	enumValueHandler EnumValueHandler
}

// NewSimplifier returns a new instance of a simplifier, with the minimal
// handler set.
func NewSimplifier(rawValueHandler RawValueHandler) *Simplifier {
	return &Simplifier{rawValueHandler: rawValueHandler}
}

func (simpl *Simplifier) rawValue(e *entry) {
	max := int64(math.Pow(2, float64(e.count*8)))
	half := max / 2
	simpl.rawValueHandler(-half, half-1)
}

// SetEnumValueHandler registers the handler for enumerations.
func (simpl *Simplifier) SetEnumValueHandler(handler EnumValueHandler) {
	simpl.enumValueHandler = handler
}

func (simpl *Simplifier) enumValue(values map[uint32]string) (result bool) {
	if simpl.enumValueHandler != nil {
		simpl.enumValueHandler(values)
		result = true
	}
	return
}
