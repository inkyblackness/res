package objprop

import "github.com/inkyblackness/res"

// Provider wraps the Provide method.
type Provider interface {
	// Provide returns the general and specific data for the requested ObjectID.
	Provide(id res.ObjectID) (generic []byte, specific []byte)
}
