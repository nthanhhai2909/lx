package lxtypes_test

import (
	"errors"
	"testing"

	"github.com/nthanhhai2909/lx/lxtypes"
)

func TestEitherLeft(t *testing.T) {
	either := lxtypes.Left[string, int]("error")

	if !either.IsLeft() {
		t.Error("Expected Left to return true for IsLeft()")
	}
	if either.IsRight() {
		t.Error("Expected Left to return false for IsRight()")
	}
	if got := either.Left(); got != "error" {
		t.Errorf("Left() = %v, want 'error'", got)
	}
}

func TestEitherRight(t *testing.T) {
	either := lxtypes.Right[string, int](42)

	if either.IsLeft() {
		t.Error("Expected Right to return false for IsLeft()")
	}
	if !either.IsRight() {
		t.Error("Expected Right to return true for IsRight()")
	}
	if got := either.Right(); got != 42 {
		t.Errorf("Right() = %v, want 42", got)
	}
}

func TestEitherRightPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Right() on Left to panic")
		}
	}()
	either := lxtypes.Left[string, int]("error")
	either.Right()
}

func TestEitherLeftPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Left() on Right to panic")
		}
	}()
	either := lxtypes.Right[string, int](42)
	either.Left()
}

func TestEitherLeftOr(t *testing.T) {
	left := lxtypes.Left[string, int]("error")
	right := lxtypes.Right[string, int](42)

	if got := left.LeftOr("default"); got != "error" {
		t.Errorf("Left.LeftOr('default') = %v, want 'error'", got)
	}
	if got := right.LeftOr("default"); got != "default" {
		t.Errorf("Right.LeftOr('default') = %v, want 'default'", got)
	}
}

func TestEitherRightOr(t *testing.T) {
	left := lxtypes.Left[string, int]("error")
	right := lxtypes.Right[string, int](42)

	if got := left.RightOr(0); got != 0 {
		t.Errorf("Left.RightOr(0) = %v, want 0", got)
	}
	if got := right.RightOr(0); got != 42 {
		t.Errorf("Right.RightOr(0) = %v, want 42", got)
	}
}

func TestEitherSwap(t *testing.T) {
	left := lxtypes.Left[string, int]("error")
	swappedLeft := left.Swap()

	if !swappedLeft.IsRight() {
		t.Error("Expected swapped Left to be Right")
	}
	if got := swappedLeft.Right(); got != "error" {
		t.Errorf("Swapped value = %v, want 'error'", got)
	}

	right := lxtypes.Right[string, int](42)
	swappedRight := right.Swap()

	if !swappedRight.IsLeft() {
		t.Error("Expected swapped Right to be Left")
	}
	if got := swappedRight.Left(); got != 42 {
		t.Errorf("Swapped value = %v, want 42", got)
	}
}

func TestEitherMapLeft(t *testing.T) {
	toUpper := func(s string) string { return "ERROR: " + s }

	left := lxtypes.Left[string, int]("test")
	mapped := lxtypes.EitherMapLeft(left, toUpper)

	if !mapped.IsLeft() {
		t.Error("Expected mapped Left to be Left")
	}
	if got := mapped.Left(); got != "ERROR: test" {
		t.Errorf("Mapped value = %v, want 'ERROR: test'", got)
	}

	right := lxtypes.Right[string, int](42)
	mappedRight := lxtypes.EitherMapLeft(right, toUpper)

	if !mappedRight.IsRight() {
		t.Error("Expected mapped Right to remain Right")
	}
	if got := mappedRight.Right(); got != 42 {
		t.Errorf("Right value = %v, want 42", got)
	}
}

func TestEitherMapRight(t *testing.T) {
	double := func(n int) int { return n * 2 }

	right := lxtypes.Right[string, int](21)
	mapped := lxtypes.EitherMapRight(right, double)

	if !mapped.IsRight() {
		t.Error("Expected mapped Right to be Right")
	}
	if got := mapped.Right(); got != 42 {
		t.Errorf("Mapped value = %v, want 42", got)
	}

	left := lxtypes.Left[string, int]("error")
	mappedLeft := lxtypes.EitherMapRight(left, double)

	if !mappedLeft.IsLeft() {
		t.Error("Expected mapped Left to remain Left")
	}
	if got := mappedLeft.Left(); got != "error" {
		t.Errorf("Left value = %v, want 'error'", got)
	}
}

func TestEitherMap(t *testing.T) {
	toUpper := func(s string) string { return "ERROR: " + s }
	double := func(n int) int { return n * 2 }

	left := lxtypes.Left[string, int]("test")
	mappedLeft := lxtypes.EitherMap(left, toUpper, double)

	if !mappedLeft.IsLeft() {
		t.Error("Expected mapped Left to be Left")
	}
	if got := mappedLeft.Left(); got != "ERROR: test" {
		t.Errorf("Mapped left = %v, want 'ERROR: test'", got)
	}

	right := lxtypes.Right[string, int](21)
	mappedRight := lxtypes.EitherMap(right, toUpper, double)

	if !mappedRight.IsRight() {
		t.Error("Expected mapped Right to be Right")
	}
	if got := mappedRight.Right(); got != 42 {
		t.Errorf("Mapped right = %v, want 42", got)
	}
}

func TestEitherFold(t *testing.T) {
	leftToString := func(s string) string { return "Error: " + s }
	rightToString := func(n int) string { return "Value: " + string(rune(n+'0')) }

	left := lxtypes.Left[string, int]("test")
	resultLeft := lxtypes.EitherFold(left, leftToString, rightToString)

	if resultLeft != "Error: test" {
		t.Errorf("Fold left = %v, want 'Error: test'", resultLeft)
	}

	right := lxtypes.Right[string, int](5)
	resultRight := lxtypes.EitherFold(right, leftToString, rightToString)

	if resultRight != "Value: 5" {
		t.Errorf("Fold right = %v, want 'Value: 5'", resultRight)
	}
}

func TestEitherFromResult(t *testing.T) {
	success := lxtypes.Success(42)
	either1 := lxtypes.EitherFromResult(success)

	if !either1.IsRight() {
		t.Error("Expected Success to convert to Right")
	}
	if got := either1.Right(); got != 42 {
		t.Errorf("Right value = %v, want 42", got)
	}

	failure := lxtypes.Failure[int](errors.New("test error"))
	either2 := lxtypes.EitherFromResult(failure)

	if !either2.IsLeft() {
		t.Error("Expected Failure to convert to Left")
	}
	if got := either2.Left().Error(); got != "test error" {
		t.Errorf("Left error = %v, want 'test error'", got)
	}
}

func TestEitherToResult(t *testing.T) {
	right := lxtypes.Right[error, int](42)
	result1 := lxtypes.EitherToResult(right)

	if !result1.IsSuccess() {
		t.Error("Expected Right to convert to Success")
	}
	if got := result1.Value(); got != 42 {
		t.Errorf("Success value = %v, want 42", got)
	}

	left := lxtypes.Left[error, int](errors.New("test error"))
	result2 := lxtypes.EitherToResult(left)

	if !result2.IsFailure() {
		t.Error("Expected Left to convert to Failure")
	}
	if got := result2.Error().Error(); got != "test error" {
		t.Errorf("Failure error = %v, want 'test error'", got)
	}
}

func TestEitherWithDifferentTypes(t *testing.T) {
	// Test with custom types
	type ValidationError struct {
		Field   string
		Message string
	}

	type User struct {
		Name string
		Age  int
	}

	// Valid case
	validUser := lxtypes.Right[ValidationError, User](User{Name: "Alice", Age: 30})
	if !validUser.IsRight() {
		t.Error("Expected valid user to be Right")
	}

	// Invalid case
	invalidUser := lxtypes.Left[ValidationError, User](
		ValidationError{Field: "age", Message: "must be positive"},
	)
	if !invalidUser.IsLeft() {
		t.Error("Expected invalid user to be Left")
	}
}

func TestEitherChaining(t *testing.T) {
	// Chain operations
	parseNum := func(s string) lxtypes.Either[string, int] {
		if s == "42" {
			return lxtypes.Right[string, int](42)
		}
		return lxtypes.Left[string, int]("parse error")
	}

	result := parseNum("42")
	doubled := lxtypes.EitherMapRight(result, func(n int) int { return n * 2 })
	final := lxtypes.EitherMapRight(doubled, func(n int) int { return n + 5 })

	if !final.IsRight() {
		t.Error("Expected chained result to be Right")
	}
	if got := final.Right(); got != 89 {
		t.Errorf("Chained result = %v, want 89", got)
	}

	// Chaining that fails
	result2 := parseNum("invalid")
	mapped := lxtypes.EitherMapRight(result2, func(n int) int { return n * 2 })

	if !mapped.IsLeft() {
		t.Error("Expected failed parse to remain Left")
	}
	if got := mapped.Left(); got != "parse error" {
		t.Errorf("Error = %v, want 'parse error'", got)
	}
}
