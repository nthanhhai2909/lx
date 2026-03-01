package lxtypes_test

import (
	"fmt"
	"strconv"

	"github.com/nthanhhai2909/lx/lxtypes"
)

func ExampleLeft() {
	either := lxtypes.Left[string, int]("error occurred")
	fmt.Println(either.IsLeft())
	fmt.Println(either.Left())
	// Output:
	// true
	// error occurred
}

func ExampleRight() {
	either := lxtypes.Right[string, int](42)
	fmt.Println(either.IsRight())
	fmt.Println(either.Right())
	// Output:
	// true
	// 42
}

func ExampleEither_RightOr() {
	left := lxtypes.Left[string, int]("error")
	right := lxtypes.Right[string, int](42)

	fmt.Println(left.RightOr(0))
	fmt.Println(right.RightOr(0))
	// Output:
	// 0
	// 42
}

func ExampleEither_Swap() {
	either := lxtypes.Left[string, int]("error")
	swapped := either.Swap()

	fmt.Println(swapped.IsRight())
	fmt.Println(swapped.Right())
	// Output:
	// true
	// error
}

func ExampleEitherMapRight() {
	either := lxtypes.Right[string, int](21)
	doubled := lxtypes.EitherMapRight(either, func(n int) int { return n * 2 })

	fmt.Println(doubled.Right())
	// Output:
	// 42
}

func ExampleEitherMapLeft() {
	either := lxtypes.Left[string, int]("error")
	mapped := lxtypes.EitherMapLeft(either, func(s string) string {
		return "Error: " + s
	})

	fmt.Println(mapped.Left())
	// Output:
	// Error: error
}

func ExampleEitherFold() {
	// Reduce either branch to a single string
	toString := func(either lxtypes.Either[string, int]) string {
		return lxtypes.EitherFold(either,
			func(s string) string { return "Error: " + s },
			func(n int) string { return fmt.Sprintf("Value: %d", n) },
		)
	}

	left := lxtypes.Left[string, int]("failed")
	right := lxtypes.Right[string, int](42)

	fmt.Println(toString(left))
	fmt.Println(toString(right))
	// Output:
	// Error: failed
	// Value: 42
}

func ExampleEitherFromResult() {
	success := lxtypes.Success(42)
	either := lxtypes.EitherFromResult(success)

	fmt.Println(either.IsRight())
	fmt.Println(either.Right())
	// Output:
	// true
	// 42
}

func ExampleEitherToResult() {
	right := lxtypes.Right[error, int](42)
	result := lxtypes.EitherToResult(right)

	fmt.Println(result.IsSuccess())
	fmt.Println(result.Value())
	// Output:
	// true
	// 42
}

func ExampleEither_parsing() {
	// Parse a value as either int or keep as string
	parseValue := func(s string) lxtypes.Either[string, int] {
		if n, err := strconv.Atoi(s); err == nil {
			return lxtypes.Right[string, int](n)
		}
		return lxtypes.Left[string, int](s)
	}

	// Process the either
	process := func(either lxtypes.Either[string, int]) string {
		return lxtypes.EitherFold(either,
			func(s string) string { return "String: " + s },
			func(n int) string { return fmt.Sprintf("Number: %d", n) },
		)
	}

	fmt.Println(process(parseValue("42")))
	fmt.Println(process(parseValue("hello")))
	// Output:
	// Number: 42
	// String: hello
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
			return lxtypes.Left[ValidationError, User](
				ValidationError{Field: "age", Message: "must be non-negative"},
			)
		}
		if name == "" {
			return lxtypes.Left[ValidationError, User](
				ValidationError{Field: "name", Message: "cannot be empty"},
			)
		}
		return lxtypes.Right[ValidationError, User](User{Name: name, Age: age})
	}

	result := validateUser("Alice", 30)
	if result.IsRight() {
		user := result.Right()
		fmt.Printf("Valid user: %s, age %d\n", user.Name, user.Age)
	} else {
		err := result.Left()
		fmt.Printf("Validation error in %s: %s\n", err.Field, err.Message)
	}
	// Output:
	// Valid user: Alice, age 30
}
