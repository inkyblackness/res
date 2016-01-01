package geometry

type simpleTextureMappedFace struct {
	abstractFace

	textureCoordinates []TextureCoordinate
}

// NewSimpleTextureMappedFace returns a new TextureMappedFace instance with given parameters.
func NewSimpleTextureMappedFace(vertices []int, textureCoordinates []TextureCoordinate) TextureMappedFace {
	return &simpleTextureMappedFace{
		abstractFace:       abstractFace{vertices: vertices},
		textureCoordinates: textureCoordinates}
}

func (face *simpleTextureMappedFace) Specialize(walker FaceWalker) {
	walker.TextureMapped(face)
}

func (face *simpleTextureMappedFace) TextureCoordinates() []TextureCoordinate {
	return face.textureCoordinates
}
