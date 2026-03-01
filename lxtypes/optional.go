package lxtypes

// Optional represents a single optional value that may or may not be present.
// This is inspired by Java's Optional<T> type and is useful for representing values
// that may or may not exist, as an alternative to using pointers or special sentinel values.
//
// An Optional is either:
//   - Present: Contains a value (created with Of or OfNullable)
//   - Empty: Contains no value (created with Empty)
//
// Common use cases:
//   - Return values from functions that might not find a result
//   - Optional configuration values
//   - Safe dictionary/map lookups
//   - Eliminating null pointer exceptions
//
// Example:
//
//	func FindUser(id int) Optional[User] {
//	    if user, exists := users[id]; exists {
//	        return Of(user)
//	    }
//	    return Empty[User]()
//	}
//
//	user := FindUser(42)
//	if user.IsPresent() {
//	    fmt.Println("Found:", user.Get())
//	} else {
//	    fmt.Println("User not found")
//	}
//
//	// Or use OrElse for a default
//	user := FindUser(42).OrElse(defaultUser)
type Optional[T any] interface {
	// IsPresent returns true if the optional contains a value.
	IsPresent() bool

	// IsEmpty returns true if the optional is empty.
	IsEmpty() bool

	// Get returns the contained value. Panics if empty.
	// Use IsPresent() to check before calling Get(), or use OrElse() for a safe alternative.
	Get() T

	// OrElse returns the contained value or the provided default if empty.
	OrElse(defaultValue T) T

	// OrElseGet returns the contained value or computes it from a function if empty.
	OrElseGet(fn func() T) T

	// Or returns this Optional if it contains a value, otherwise returns other.
	Or(other Optional[T]) Optional[T]

	// OrElseSupply returns this Optional if it contains a value, otherwise calls fn.
	OrElseSupply(fn func() Optional[T]) Optional[T]
}

type presentOption[T any] struct {
	value T
}

type emptyOption[T any] struct{}

// Of creates an Optional containing the given value.
// Panics if value is nil for pointer types (use OfNullable for nil-safe creation).
func Of[T any](value T) Optional[T] {
	return presentOption[T]{value: value}
}

// OfNullable creates an Optional from a pointer.
// Returns Empty if the pointer is nil, otherwise returns an Optional with the dereferenced value.
func OfNullable[T any](ptr *T) Optional[T] {
	if ptr == nil {
		return Empty[T]()
	}
	return Of(*ptr)
}

// Empty creates an empty Optional.
func Empty[T any]() Optional[T] {
	return emptyOption[T]{}
}

func (o presentOption[T]) IsPresent() bool {
	return true
}

func (o presentOption[T]) IsEmpty() bool {
	return false
}

func (o presentOption[T]) Get() T {
	return o.value
}

func (o presentOption[T]) OrElse(defaultValue T) T {
	return o.value
}

func (o presentOption[T]) OrElseGet(fn func() T) T {
	return o.value
}

func (o presentOption[T]) Or(other Optional[T]) Optional[T] {
	return o
}

func (o presentOption[T]) OrElseSupply(fn func() Optional[T]) Optional[T] {
	return o
}

func (o emptyOption[T]) IsPresent() bool {
	return false
}

func (o emptyOption[T]) IsEmpty() bool {
	return true
}

func (o emptyOption[T]) Get() T {
	panic("called Get() on an empty Optional")
}

func (o emptyOption[T]) OrElse(defaultValue T) T {
	return defaultValue
}

func (o emptyOption[T]) OrElseGet(fn func() T) T {
	return fn()
}

func (o emptyOption[T]) Or(other Optional[T]) Optional[T] {
	return other
}

func (o emptyOption[T]) OrElseSupply(fn func() Optional[T]) Optional[T] {
	return fn()
}
