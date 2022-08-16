package node

const (
	RemoteKeeper = "remote"
)

// Source represents a generic source that allows to read the data of a specific SDK module
type Source interface {
	// Type returns the type of the keeper
	Type() string
}
