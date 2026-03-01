package lxtypes_test

import (
	"fmt"
	"strings"

	"github.com/nthanhhai2909/lx/lxtypes"
)

// ============================================================================
// Pair Examples
// ============================================================================

func ExamplePair() {
	// Creating a pair
	p := lxtypes.NewPair(42, "answer")
	fmt.Printf("First: %d, Second: %s\n", p.First, p.Second)

	// Unpacking values
	x, y := p.Values()
	fmt.Printf("Unpacked: %d, %s\n", x, y)
	// Output:
	// First: 42, Second: answer
	// Unpacked: 42, answer
}

func ExamplePair_Swap() {
	p := lxtypes.NewPair(42, "answer")
	swapped := p.Swap()
	fmt.Printf("Original: (%d, %s)\n", p.First, p.Second)
	fmt.Printf("Swapped: (%s, %d)\n", swapped.First, swapped.Second)
	// Output:
	// Original: (42, answer)
	// Swapped: (answer, 42)
}

func ExamplePair_MapFirst() {
	p := lxtypes.NewPair(5, "test")
	doubled := p.MapFirst(func(n int) int { return n * 2 })
	fmt.Printf("Original: (%d, %s)\n", p.First, p.Second)
	fmt.Printf("After MapFirst: (%d, %s)\n", doubled.First, doubled.Second)
	// Output:
	// Original: (5, test)
	// After MapFirst: (10, test)
}

func ExamplePair_MapSecond() {
	p := lxtypes.NewPair(5, "test")
	upper := p.MapSecond(func(s string) string { return strings.ToUpper(s) })
	fmt.Printf("Original: (%d, %s)\n", p.First, p.Second)
	fmt.Printf("After MapSecond: (%d, %s)\n", upper.First, upper.Second)
	// Output:
	// Original: (5, test)
	// After MapSecond: (5, TEST)
}

// ============================================================================
// Triple Examples
// ============================================================================

func ExampleTriple() {
	// Creating a triple
	t := lxtypes.NewTriple(1, "hello", true)
	fmt.Printf("First: %d, Second: %s, Third: %t\n", t.First, t.Second, t.Third)

	// Unpacking values
	x, y, z := t.Values()
	fmt.Printf("Unpacked: %d, %s, %t\n", x, y, z)
	// Output:
	// First: 1, Second: hello, Third: true
	// Unpacked: 1, hello, true
}

func ExampleTriple_ToPair() {
	t := lxtypes.NewTriple(1, "hello", true)
	p := t.ToPair()
	fmt.Printf("Triple: (%d, %s, %t)\n", t.First, t.Second, t.Third)
	fmt.Printf("Pair: (%d, %s)\n", p.First, p.Second)
	// Output:
	// Triple: (1, hello, true)
	// Pair: (1, hello)
}

// ============================================================================
// Quad Examples
// ============================================================================

func ExampleQuad() {
	// Creating a quad
	q := lxtypes.NewQuad(1, "hello", true, 3.14)
	fmt.Printf("Values: %d, %s, %t, %.2f\n", q.First, q.Second, q.Third, q.Fourth)

	// Unpacking values
	w, x, y, z := q.Values()
	fmt.Printf("Unpacked: %d, %s, %t, %.2f\n", w, x, y, z)
	// Output:
	// Values: 1, hello, true, 3.14
	// Unpacked: 1, hello, true, 3.14
}

func ExampleQuad_ToTriple() {
	q := lxtypes.NewQuad(1, "hello", true, 3.14)
	t := q.ToTriple()
	fmt.Printf("Quad: (%d, %s, %t, %.2f)\n", q.First, q.Second, q.Third, q.Fourth)
	fmt.Printf("Triple: (%d, %s, %t)\n", t.First, t.Second, t.Third)
	// Output:
	// Quad: (1, hello, true, 3.14)
	// Triple: (1, hello, true)
}

func ExampleQuad_ToPair() {
	q := lxtypes.NewQuad(1, "hello", true, 3.14)
	p := q.ToPair()
	fmt.Printf("Quad: (%d, %s, %t, %.2f)\n", q.First, q.Second, q.Third, q.Fourth)
	fmt.Printf("Pair: (%d, %s)\n", p.First, p.Second)
	// Output:
	// Quad: (1, hello, true, 3.14)
	// Pair: (1, hello)
}
