// Package lxtypes provides reusable generic type definitions for common programming patterns.
//
// Functional Interfaces (inspired by Java's java.util.function):
//
// Basic Interfaces:
//   - Predicate[T]: Tests a condition (T -> bool)
//   - Consumer[T]: Performs an operation (T -> void)
//   - Function[T, U]: Transforms input (T -> U)
//   - Supplier[T]: Provides a value (() -> T)
//
// Binary Interfaces:
//   - BiPredicate[T, U]: Tests condition on two inputs
//   - BiConsumer[T, U]: Operates on two inputs
//   - BiFunction[T, U, R]: Transforms two inputs to output
//
// Specialized Operators:
//   - UnaryOperator[T]: Transforms T to T
//   - BinaryOperator[T]: Combines two T into T
//   - Comparator[T]: Compares two T values
//
// Tuple Types:
//   - Pair[T, U]: Two-element tuple with helper methods
//   - Triple[T, U, V]: Three-element tuple
//   - Quad[T, U, V, W]: Four-element tuple
//
// Utility Functions:
//   - Compose[T, U, V]: Composes two functions (f ∘ g)
//   - NewPair[T, U]: Creates a new Pair
//   - NewTriple[T, U, V]: Creates a new Triple
//   - NewQuad[T, U, V, W]: Creates a new Quad
//
// Example Usage:
//
//	// Filter with Predicate
//	isEven := lxtypes.Predicate[int](func(n int) bool {
//	    return n%2 == 0
//	})
//
//	// Transform with Function
//	toString := lxtypes.Function[int, string](func(n int) string {
//	    return strconv.Itoa(n)
//	})
//
//	// Combine with BinaryOperator
//	sum := lxtypes.BinaryOperator[int](func(a, b int) int {
//	    return a + b
//	})
//
//	// Compose functions
//	double := func(n int) int { return n * 2 }
//	addTen := func(n int) int { return n + 10 }
//	composed := lxtypes.Compose(double, addTen)
//
//	// Work with tuples
//	p := lxtypes.NewPair(42, "answer")
//	fmt.Println(p.First, p.Second)
//	swapped := p.Swap()  // Pair[string, int]{"answer", 42}
//
//	t := lxtypes.NewTriple(1, "hello", true)
//	x, y, z := t.Values()  // Unpack values
//
// This package is designed to be extended with additional type definitions
// such as Result[T, E], Option[T], Either[L, R], and other functional types.
package lxtypes
