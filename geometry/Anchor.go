package geometry

// Anchor describes a reference for further nodes or faces
type Anchor interface {
	// Normal returns the normal vector of the anchor.
	Normal() Vector
	// Reference returns the position of the anchor.
	Reference() Vector
}
