package geometry

// AnchorWalker implementations receive specific anchors from a node.
type AnchorWalker interface {
	// Nodes is called for a NodeAnchor.
	Nodes(surface NodeAnchor)
	// Faces is called for a FaceAnchor.
	Faces(surface FaceAnchor)
}
