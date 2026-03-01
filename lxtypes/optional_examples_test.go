package lxtypes_test

import (
	"fmt"

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
