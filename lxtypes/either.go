package lxtypes

import (
	"errors"
)

var (
	ErrLeftOnRight = errors.New("lxtypes.Either: cannot get Left value from Right-sided Either - use Right() or RightOr() instead")
	ErrRightOnLeft = errors.New("lxtypes.Either: cannot get Right value from Left-sided Either - use Left() or LeftOr() instead")
)

// Either represents a value that can be either Left or Right.
type Either[L, R any] interface {
	// IsLeft returns true if this is a Left value.
	IsLeft() bool

	// IsRight returns true if this is a Right value.
	IsRight() bool

	// Left returns the Left value. Return an error if this is Right.
	Left() (L, error)

	// Right returns the Right value. Return an error if this is Left.
	Right() (R, error)

	// LeftOr returns the Left value or the provided default if Right.
	LeftOr(defaultValue L) L

	// RightOr returns the Right value or the provided default if Left.
	RightOr(defaultValue R) R
}

// EitherLeft creates an Either with Left value.
func EitherLeft[L, R any](value L) Either[L, R] {
	return eitherLeft[L, R]{value: value}
}

// EitherRight creates an Either with Right value.
func EitherRight[L, R any](value R) Either[L, R] {
	return eitherRight[L, R]{value: value}
}

// ------------------------------------ Either Left implementation ------------------------------------

type eitherLeft[L, R any] struct {
	value L
}

func (e eitherLeft[L, R]) Left() (L, error) {
	return e.value, nil
}

func (e eitherLeft[L, R]) Right() (R, error) {
	var zero R
	return zero, ErrRightOnLeft
}

func (e eitherLeft[L, R]) IsLeft() bool {
	return true
}

func (e eitherLeft[L, R]) IsRight() bool {
	return false
}

func (e eitherLeft[L, R]) LeftOr(defaultValue L) L {
	return e.value
}

func (e eitherLeft[L, R]) RightOr(defaultValue R) R {
	return defaultValue
}

// ------------------------------------ Either Right implementation ------------------------------------

func (e eitherRight[L, R]) IsLeft() bool {
	return false
}

func (e eitherRight[L, R]) IsRight() bool {
	return true
}

type eitherRight[L, R any] struct {
	value R
}

func (e eitherRight[L, R]) Left() (L, error) {
	var zero L
	return zero, ErrLeftOnRight
}

func (e eitherRight[L, R]) Right() (R, error) {
	return e.value, nil
}

func (e eitherRight[L, R]) LeftOr(defaultValue L) L {
	return defaultValue
}

func (e eitherRight[L, R]) RightOr(defaultValue R) R {
	return e.value
}
