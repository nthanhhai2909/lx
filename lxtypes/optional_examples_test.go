package lxtypes_test

import (
	"fmt"
	"strconv"

	"github.com/nthanhhai2909/lx/lxtypes"
)

func ExampleOf() {
	opt := lxtypes.Of(42)
	fmt.Println(opt.IsPresent())
	fmt.Println(opt.Get())
	// Output:
	// true
	// 42
}

func ExampleEmpty() {
	opt := lxtypes.Empty[int]()
	fmt.Println(opt.IsEmpty())
	fmt.Println(opt.OrElse(99))
	// Output:
	// true
	// 99
}

func ExampleOfNullable() {
	// With non-nil pointer
	value := 42
	opt1 := lxtypes.OfNullable(&value)
	fmt.Println(opt1.IsPresent())
	fmt.Println(opt1.Get())

	// With nil pointer
	var nilPtr *int
	opt2 := lxtypes.OfNullable(nilPtr)
	fmt.Println(opt2.IsEmpty())
	// Output:
	// true
	// 42
	// true
}

func ExampleOptional_OrElse() {
	present := lxtypes.Of(42)
	empty := lxtypes.Empty[int]()

	fmt.Println(present.OrElse(0))
	fmt.Println(empty.OrElse(99))
	// Output:
	// 42
	// 99
}

func ExampleOptional_OrElseGet() {
	present := lxtypes.Of(42)
	empty := lxtypes.Empty[int]()

	fmt.Println(present.OrElseGet(func() int { return 0 }))
	fmt.Println(empty.OrElseGet(func() int { return 99 }))
	// Output:
	// 42
	// 99
}

func ExampleOptionalMap() {
	opt := lxtypes.Of(21)
	doubled := lxtypes.OptionalMap(opt, func(n int) int { return n * 2 })
	fmt.Println(doubled.Get())

	empty := lxtypes.Empty[int]()
	mappedEmpty := lxtypes.OptionalMap(empty, func(n int) int { return n * 2 })
	fmt.Println(mappedEmpty.IsEmpty())
	// Output:
	// 42
	// true
}

func ExampleOptionalAndThen() {
	safeDivide := func(n int) lxtypes.Optional[int] {
		if n == 0 {
			return lxtypes.Empty[int]()
		}
		return lxtypes.Of(100 / n)
	}

	opt10 := lxtypes.Of(10)
	result := lxtypes.OptionalAndThen(opt10, safeDivide)
	fmt.Println(result.Get())

	optZero := lxtypes.Of(0)
	resultEmpty := lxtypes.OptionalAndThen(optZero, safeDivide)
	fmt.Println(resultEmpty.IsEmpty())
	// Output:
	// 10
	// true
}

func ExampleOptional_Or() {
	opt1 := lxtypes.Of(42)
	opt2 := lxtypes.Of(99)
	empty := lxtypes.Empty[int]()

	fmt.Println(opt1.Or(opt2).Get())
	fmt.Println(empty.Or(opt2).Get())
	fmt.Println(empty.Or(lxtypes.Empty[int]()).IsEmpty())
	// Output:
	// 42
	// 99
	// true
}

func ExampleOptional_chaining() {
	// Real-world example: parse and validate
	parseNum := func(s string) lxtypes.Optional[int] {
		n, err := strconv.Atoi(s)
		if err != nil {
			return lxtypes.Empty[int]()
		}
		return lxtypes.Of(n)
	}

	// Validate that number is positive
	validatePositive := func(n int) lxtypes.Optional[int] {
		if n > 0 {
			return lxtypes.Of(n)
		}
		return lxtypes.Empty[int]()
	}

	result := parseNum("42")
	validated := lxtypes.OptionalAndThen(result, validatePositive)
	doubled := lxtypes.OptionalMap(validated, func(n int) int { return n * 2 })

	fmt.Println(doubled.OrElse(0))

	// Failed case
	invalid := parseNum("invalid")
	fmt.Println(invalid.OrElse(-1))
	// Output:
	// 84
	// -1
}
