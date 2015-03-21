package textprop

import "github.com/inkyblackness/res"

// Provider wraps the Provide method.
type Provider interface {
	// TextureCount returns the amount of textures available
	TextureCount() uint32
	// Provide returns the data for the requested TextureID.
	Provide(id res.TextureID) []byte
}
