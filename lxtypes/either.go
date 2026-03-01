package lxtypes

// Either represents a value of one of two possible types (a disjoint union).
// An Either is either Left or Right, never both.
//
// This is a general-purpose type for representing binary choices where both alternatives
// are valid values (not necessarily errors). For error handling specifically, use Result[T].
//
// Common use cases:
//   - Different return types based on conditions
//   - Validation that can produce different types of results
//   - Polymorphic return values
//   - Union types before pattern matching
//
// Conventions:
//   - Left is often used for "failure" or "error" cases
//   - Right is often used for "success" or "normal" cases
//   - But both are equally valid values
//
// Example:
//
//	// Parse a value as either int or string
//	func ParseValue(s string) Either[int, string] {
//	    if n, err := strconv.Atoi(s); err == nil {
//	        return Right[int, string](n)
//	    }
//	    return Left[int, string](s)
//	}
//
//	either := ParseValue("42")
//	if either.IsRight() {
//	    fmt.Println("Number:", either.Right())
//	} else {
//	    fmt.Println("String:", either.Left())
//	}
type Either[L, R any] interface {
	// IsLeft returns true if this is a Left value.
	IsLeft() bool

	// IsRight returns true if this is a Right value.
	IsRight() bool

	// Left returns the Left value. Panics if this is a Right.
	Left() L

	// Right returns the Right value. Panics if this is a Left.
	Right() R

	// LeftOr returns the Left value or the provided default if Right.
	LeftOr(defaultValue L) L

	// RightOr returns the Right value or the provided default if Left.
	RightOr(defaultValue R) R

	// Swap returns an Either with Left and Right swapped.
	Swap() Either[R, L]
}

// leftValue represents a Left Either.
type leftValue[L, R any] struct {
	value L
}

// rightValue represents a Right Either.
type rightValue[L, R any] struct {
	value R
}

// Left creates an Either with a Left value.
func Left[L, R any](value L) Either[L, R] {
	return leftValue[L, R]{value: value}
}

// Right creates an Either with a Right value.
func Right[L, R any](value R) Either[L, R] {
	return rightValue[L, R]{value: value}
}

func (e leftValue[L, R]) IsLeft() bool {
	return true
}

func (e leftValue[L, R]) IsRight() bool {
	return false
}

func (e leftValue[L, R]) Left() L {
	return e.value
}

func (e leftValue[L, R]) Right() R {
	panic("called Right() on a Left value")
}

func (e leftValue[L, R]) LeftOr(defaultValue L) L {
	return e.value
}

func (e leftValue[L, R]) RightOr(defaultValue R) R {
	return defaultValue
}

func (e leftValue[L, R]) Swap() Either[R, L] {
	return Right[R, L](e.value)
}

func (e rightValue[L, R]) IsLeft() bool {
	return false
}

func (e rightValue[L, R]) IsRight() bool {
	return true
}

func (e rightValue[L, R]) Left() L {
	panic("called Left() on a Right value")
}

func (e rightValue[L, R]) Right() R {
	return e.value
}

func (e rightValue[L, R]) LeftOr(defaultValue L) L {
	return defaultValue
}

func (e rightValue[L, R]) RightOr(defaultValue R) R {
	return e.value
}

func (e rightValue[L, R]) Swap() Either[R, L] {
	return Left[R, L](e.value)
}

// EitherMapLeft transforms the Left value using the provided function.
// If this is a Right, returns the Right unchanged.
// This is a standalone function because Go doesn't support type parameters on interface methods.
func EitherMapLeft[L, L2, R any](either Either[L, R], fn func(L) L2) Either[L2, R] {
	if either.IsLeft() {
		return Left[L2, R](fn(either.Left()))
	}
	return Right[L2, R](either.Right())
}

// EitherMapRight transforms the Right value using the provided function.
// If this is a Left, returns the Left unchanged.
// This is a standalone function because Go doesn't support type parameters on interface methods.
func EitherMapRight[L, R, R2 any](either Either[L, R], fn func(R) R2) Either[L, R2] {
	if either.IsRight() {
		return Right[L, R2](fn(either.Right()))
	}
	return Left[L, R2](either.Left())
}

// EitherMap transforms both Left and Right values using the provided functions.
// This is a standalone function because Go doesn't support type parameters on interface methods.
func EitherMap[L, L2, R, R2 any](either Either[L, R], leftFn func(L) L2, rightFn func(R) R2) Either[L2, R2] {
	if either.IsLeft() {
		return Left[L2, R2](leftFn(either.Left()))
	}
	return Right[L2, R2](rightFn(either.Right()))
}

// EitherFold reduces an Either to a single value by applying one of two functions.
// This is a standalone function because Go doesn't support type parameters on interface methods.
func EitherFold[L, R, T any](either Either[L, R], leftFn func(L) T, rightFn func(R) T) T {
	if either.IsLeft() {
		return leftFn(either.Left())
	}
	return rightFn(either.Right())
}

// EitherFromResult converts a Result to an Either.
// Success becomes Right, Failure becomes Left with the error.
func EitherFromResult[T any](result Result[T]) Either[error, T] {
	if result.IsSuccess() {
		return Right[error, T](result.Value())
	}
	return Left[error, T](result.Error())
}

// EitherToResult converts an Either with error on the left to a Result.
// Only works when L is error type.
func EitherToResult[T any](either Either[error, T]) Result[T] {
	if either.IsRight() {
		return Success(either.Right())
	}
	return Failure[T](either.Left())
}
