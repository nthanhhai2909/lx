package lxtypes_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxtypes"
)

func TestEitherLeft(t *testing.T) {
	either := lxtypes.EitherLeft[string, int]("error")

	// Test IsLeft returns true
	if !either.IsLeft() {
		t.Error("Expected IsLeft() to return true for Left-sided Either")
	}

	// Test IsRight returns false
	if either.IsRight() {
		t.Error("Expected IsRight() to return false for Left-sided Either")
	}

	// Test Left() returns value without error
	left, err := either.Left()
	if err != nil {
		t.Errorf("Left() returned unexpected error: %v", err)
	}
	if left != "error" {
		t.Errorf("Left value = %v, want 'error'", left)
	}

	// Test Right() returns error
	_, err = either.Right()
	if err == nil {
		t.Error("Right() should return error on Left-sided Either")
	}
	if err != lxtypes.ErrRightOnLeft {
		t.Errorf("Right() error = %v, want ErrRightOnLeft", err)
	}
}

func TestEitherRight(t *testing.T) {
	either := lxtypes.EitherRight[string, int](42)

	// Test IsLeft returns false
	if either.IsLeft() {
		t.Error("Expected IsLeft() to return false for Right-sided Either")
	}

	// Test IsRight returns true
	if !either.IsRight() {
		t.Error("Expected IsRight() to return true for Right-sided Either")
	}

	// Test Right() returns value without error
	right, err := either.Right()
	if err != nil {
		t.Errorf("Right() returned unexpected error: %v", err)
	}
	if right != 42 {
		t.Errorf("Right value = %v, want 42", right)
	}

	// Test Left() returns error
	_, err = either.Left()
	if err == nil {
		t.Error("Left() should return error on Right-sided Either")
	}
	if err != lxtypes.ErrLeftOnRight {
		t.Errorf("Left() error = %v, want ErrLeftOnRight", err)
	}
}

func TestEitherIsLeft(t *testing.T) {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	if !left.IsLeft() {
		t.Error("EitherLeft.IsLeft() should return true")
	}

	if right.IsLeft() {
		t.Error("EitherRight.IsLeft() should return false")
	}
}

func TestEitherIsRight(t *testing.T) {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	if left.IsRight() {
		t.Error("EitherLeft.IsRight() should return false")
	}

	if !right.IsRight() {
		t.Error("EitherRight.IsRight() should return true")
	}
}

func TestEitherLeftOr(t *testing.T) {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	if got := left.LeftOr("default"); got != "error" {
		t.Errorf("LeftOr on Left = %v, want 'error'", got)
	}

	if got := right.LeftOr("default"); got != "default" {
		t.Errorf("LeftOr on Right = %v, want 'default'", got)
	}
}

func TestEitherRightOr(t *testing.T) {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	if got := left.RightOr(0); got != 0 {
		t.Errorf("RightOr on Left = %v, want 0", got)
	}

	if got := right.RightOr(0); got != 42 {
		t.Errorf("RightOr on Right = %v, want 42", got)
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

	validUser := lxtypes.EitherRight[ValidationError, User](User{Name: "Alice", Age: 30})
	user, err := validUser.Right()
	if err != nil {
		t.Errorf("Right() on valid user returned error: %v", err)
	}
	if user.Name != "Alice" {
		t.Errorf("User name = %v, want 'Alice'", user.Name)
	}

	invalidUser := lxtypes.EitherLeft[ValidationError, User](
		ValidationError{Field: "age", Message: "must be positive"},
	)
	validationErr, err := invalidUser.Left()
	if err != nil {
		t.Errorf("Left() on invalid user returned error: %v", err)
	}
	if validationErr.Field != "age" {
		t.Errorf("Validation error field = %v, want 'age'", validationErr.Field)
	}
}
