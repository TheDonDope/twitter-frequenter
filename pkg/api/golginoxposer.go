package api

// Golginoxposer describes methods
type Golginoxposer interface {
	// ConcurrentExecute leverages Go routines and channels during runtime.
	ConcurrentExecute()
	// SerialExecute processes the data in a serial manner.
	SerialExecute()
}
