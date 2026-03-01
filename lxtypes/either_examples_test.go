package lxtypes_test

import (
	"fmt"

	"github.com/nthanhhai2909/lx/lxtypes"
)

func ExampleEitherLeft() {
	either := lxtypes.EitherLeft[string, int]("error occurred")

	fmt.Println(either.IsLeft())  // Check if Left
	fmt.Println(either.IsRight()) // Check if Right

	left, err := either.Left()
	if err == nil {
		fmt.Println(left)
	}
	// Output:
	// true
	// false
	// error occurred
}

func ExampleEitherRight() {
	either := lxtypes.EitherRight[string, int](42)

	fmt.Println(either.IsLeft())  // Check if Left
	fmt.Println(either.IsRight()) // Check if Right

	right, err := either.Right()
	if err == nil {
		fmt.Println(right)
	}
	// Output:
	// false
	// true
	// 42
}

func ExampleEither_IsLeft() {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	// Pattern matching with IsLeft/IsRight
	if left.IsLeft() {
		val, _ := left.Left()
		fmt.Println("Left:", val)
	}

	if right.IsRight() {
		val, _ := right.Right()
		fmt.Println("Right:", val)
	}
	// Output:
	// Left: error
	// Right: 42
}

func ExampleEither_LeftOr() {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	fmt.Println(left.LeftOr("default"))
	fmt.Println(right.LeftOr("default"))
	// Output:
	// error
	// default
}

func ExampleEither_RightOr() {
	left := lxtypes.EitherLeft[string, int]("error")
	right := lxtypes.EitherRight[string, int](42)

	fmt.Println(left.RightOr(0))
	fmt.Println(right.RightOr(0))
	// Output:
	// 0
	// 42
}

func ExampleEither_validation() {
	type ValidationError struct {
		Field   string
		Message string
	}

	type User struct {
		Name string
		Age  int
	}

	validateUser := func(name string, age int) lxtypes.Either[ValidationError, User] {
		if age < 0 {
			return lxtypes.EitherLeft[ValidationError, User](
				ValidationError{Field: "age", Message: "must be non-negative"},
			)
		}
		if name == "" {
			return lxtypes.EitherLeft[ValidationError, User](
				ValidationError{Field: "name", Message: "cannot be empty"},
			)
		}
		return lxtypes.EitherRight[ValidationError, User](User{Name: name, Age: age})
	}

	result := validateUser("Alice", 30)
	if result.IsRight() {
		user, _ := result.Right()
		fmt.Printf("Valid user: %s, age %d\n", user.Name, user.Age)
	} else if result.IsLeft() {
		err, _ := result.Left()
		fmt.Printf("Validation error in %s: %s\n", err.Field, err.Message)
	}
	// Output:
	// Valid user: Alice, age 30
}
