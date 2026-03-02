package lxtypes

import "time"

// Future represents a value that will be available in the future (async).
type Future[T any] interface {
	// Get blocks until the value is ready, returns value and error.
	Get() (T, error)

	// GetWithTimeout waits up to duration, returns an error if timeout.
	GetWithTimeout(timeout time.Duration) (T, error)

	// IsComplete returns true if computation finished (success or failure).
	IsComplete() bool

	// Then chains another operation (runs after this completes).
	Then(fn func(T) T) Future[T]

	// OnComplete registers a callback for when done.
	OnComplete(fn func(T, error))
}

// TODO: Implement a simple Future using goroutines and channels.
