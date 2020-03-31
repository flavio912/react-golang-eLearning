package gentypes

// GraphType should be implemented by types that have loaders
type GraphType interface {
	Key() string
}
