package lxtypes_test

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/nthanhhai2909/lx/lxtypes"
)

func ExampleSuccess() {
	result := lxtypes.Success(42)
	fmt.Println(result.IsSuccess())
	fmt.Println(result.Value())
	// Output:
	// true
	// 42
}

func ExampleFailure() {
	result := lxtypes.Failure[int](errors.New("something went wrong"))
	fmt.Println(result.IsFailure())
	fmt.Println(result.Error())
	// Output:
	// true
	// something went wrong
}

func ExampleFromError() {
	// Success case
	value, err := strconv.Atoi("42")
	result := lxtypes.FromError(value, err)
	fmt.Println(result.IsSuccess())
	fmt.Println(result.Value())

	// Output:
	// true
	// 42
}

func ExampleResult_ValueOr() {
	success := lxtypes.Success(42)
	failure := lxtypes.Failure[int](errors.New("error"))

	fmt.Println(success.ValueOr(0))
	fmt.Println(failure.ValueOr(99))
	// Output:
	// 42
	// 99
}

func ExampleResultMap() {
	result := lxtypes.Success(21)
	doubled := lxtypes.ResultMap(result, func(n int) int { return n * 2 })
	fmt.Println(doubled.Value())
	// Output:
	// 42
}

func ExampleResultAndThen() {
	divide := func(a, b int) lxtypes.Result[int] {
		if b == 0 {
			return lxtypes.Failure[int](errors.New("division by zero"))
		}
		return lxtypes.Success(a / b)
	}

	result := divide(10, 2)
	chained := lxtypes.ResultAndThen(result, func(n int) lxtypes.Result[int] {
		return divide(n, 1)
	})

	fmt.Println(chained.Value())
	// Output:
	// 5
}

func ExampleResultRecover() {
	failure := lxtypes.Failure[int](errors.New("error"))

	recovered := lxtypes.ResultRecover(failure, func(e error) lxtypes.Result[int] {
		fmt.Println("Recovering from:", e)
		return lxtypes.Success(99)
	})

	fmt.Println(recovered.Value())
	// Output:
	// Recovering from: error
	// 99
}

func ExampleResult_chaining() {
	// Real-world example: parse and validate
	parseAge := func(s string) lxtypes.Result[int] {
		age, err := strconv.Atoi(s)
		return lxtypes.FromError(age, err)
	}

	validateAge := func(age int) lxtypes.Result[int] {
		if age >= 18 {
			return lxtypes.Success(age)
		}
		return lxtypes.Failure[int](errors.New("must be 18 or older"))
	}

	result := parseAge("25")
	validated := lxtypes.ResultAndThen(result, validateAge)
	doubled := lxtypes.ResultMap(validated, func(age int) int { return age * 2 })

	fmt.Println(doubled.ValueOr(0))
	// Output:
	// 50
}
