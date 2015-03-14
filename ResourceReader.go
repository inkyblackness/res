package res

// ResourceReader provides access to resources from an arbitrary source
type ResourceReader interface {
	RequestResource(id ResourceID)
}
