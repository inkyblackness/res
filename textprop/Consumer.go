package textprop

// Consumer wraps methods to consume texture property data
type Consumer interface {
	// Consume takes the provided data and adds it to the existing ones
	Consume(data []byte)
}
