package geometry

// TextureMappedFace is a face with a texture.
type TextureMappedFace interface {
	Face

	// TextureCoordinates returns the list of coordinates for the vertices.
	TextureCoordinates() []TextureCoordinate
}
