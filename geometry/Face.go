package geometry

// Face describes a rendered plane.
type Face interface {
	// Vertices returns the list of indices of the vertices for the face.
	Vertices() []int
}
