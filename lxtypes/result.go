package lxtypes

// Result represents the result of an operation that may succeed with a value or fail with an error.
// This is specifically designed for error handling with Go's error type.
//
// A Result is either:
//   - Success: Contains a value (created with Success)
//   - Failure: Contains an error (created with Failure)
//
// This is inspired by Rust's Result<T, E> but specialized for Go's error handling patterns.
// For general binary choice between any two types, use Either[L, R].
//
// Example:
//
//	func Divide(a, b int) Result[int] {
//	    if b == 0 {
//	        return Failure[int](errors.New("division by zero"))
//	    }
//	    return Success(a / b)
//	}
//
//	result := Divide(10, 2)
//	if result.IsSuccess() {
//	    fmt.Println("Result:", result.Value())
//	} else {
//	    fmt.Println("Error:", result.Error())
//	}
//
//	// Or use ValueOr for a default
//	value := result.ValueOr(0)
type Result[T any] interface {
	// IsSuccess returns true if the result is successful.
	IsSuccess() bool

	// IsFailure returns true if the result is a failure.
	IsFailure() bool

	// Value returns the success value. Panics if the result is a failure.
	// Use IsSuccess() to check before calling Value(), or use ValueOr() for a safe alternative.
	Value() T

	// ValueOr returns the success value or the provided default if failure.
	ValueOr(defaultValue T) T

	// ValueOrElse returns the success value or computes it from the error if failure.
	ValueOrElse(fn func(error) T) T

	// Error returns the error. Panics if the result is successful.
	Error() error

	// OrElse returns this Result if successful, otherwise calls fn with the error.
	OrElse(fn func(error) Result[T]) Result[T]
}

// successResult represents a successful result.
type successResult[T any] struct {
	value T
}

// failureResult represents a failed result.
type failureResult[T any] struct {
	err error
}

// Success creates a successful Result with the given value.
func Success[T any](value T) Result[T] {
	return successResult[T]{value: value}
}

// Failure creates a failed Result with the given error.
func Failure[T any](err error) Result[T] {
	return failureResult[T]{err: err}
}

// FromError creates a Result from a (value, error) pair - common in Go.
// If error is nil, returns Success(value), otherwise returns Failure(error).
//
// Example:
//
//	value, err := strconv.Atoi("42")
//	result := FromError(value, err)  // Success(42)
func FromError[T any](value T, err error) Result[T] {
	if err != nil {
		return Failure[T](err)
	}
	return Success(value)
}

// successResult implementation
func (r successResult[T]) IsSuccess() bool {
	return true
}

func (r successResult[T]) IsFailure() bool {
	return false
}

func (r successResult[T]) Value() T {
	return r.value
}

func (r successResult[T]) ValueOr(defaultValue T) T {
	return r.value
}

func (r successResult[T]) ValueOrElse(fn func(error) T) T {
	return r.value
}

func (r successResult[T]) Error() error {
	panic("called Error() on a successful Result")
}

func (r successResult[T]) OrElse(fn func(error) Result[T]) Result[T] {
	return r
}

// failureResult implementation
func (r failureResult[T]) IsSuccess() bool {
	return false
}

func (r failureResult[T]) IsFailure() bool {
	return true
}

func (r failureResult[T]) Value() T {
	panic("called Value() on a failed Result")
}

func (r failureResult[T]) ValueOr(defaultValue T) T {
	return defaultValue
}

func (r failureResult[T]) ValueOrElse(fn func(error) T) T {
	return fn(r.err)
}

func (r failureResult[T]) Error() error {
	return r.err
}

func (r failureResult[T]) OrElse(fn func(error) Result[T]) Result[T] {
	return fn(r.err)
}

// ResultMap transforms the success value using the provided function.
// If the result is a failure, returns the failure unchanged.
// This is a standalone function because Go doesn't support type parameters on interface methods.
func ResultMap[T, U any](res Result[T], fn func(T) U) Result[U] {
	if res.IsFailure() {
		return Failure[U](res.Error())
	}
	return Success(fn(res.Value()))
}

// ResultAndThen chains another Result-returning operation.
// If the result is a failure, returns the failure unchanged.
// This is a standalone function because Go doesn't support type parameters on interface methods.
func ResultAndThen[T, U any](res Result[T], fn func(T) Result[U]) Result[U] {
	if res.IsFailure() {
		return Failure[U](res.Error())
	}
	return fn(res.Value())
}

// ResultRecover attempts to recover from a failure by calling fn with the error.
// If the result is successful, returns it unchanged.
func ResultRecover[T any](res Result[T], fn func(error) Result[T]) Result[T] {
	return res.OrElse(fn)
}
