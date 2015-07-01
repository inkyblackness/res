package text

// Codepage wraps the methods for serializing strings.
type Codepage interface {
	// Encode converts the provided string to a byte array.
	Encode(value string) []byte
	// Decode converts the provided byte slice to a string.
	Decode(data []byte) string
}
