package serial

// Coder represents an Encoder/decoder for binary data.
type Coder interface {
	// FirstError returns the first error this coder encountered.
	FirstError() error

	// Code serializes the given value.
	// Any error state will be provided via FirstError(). Code will do nothing
	// if the coder is already in error state.
	Code(value interface{})
}
