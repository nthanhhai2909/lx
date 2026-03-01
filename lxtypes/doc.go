// Package lxtypes provides reusable generic type definitions for functional programming,
// optional values, and error handling in Go.
//
// This package includes three main categories of types:
//
// 1. Functional Interfaces (inspired by Java's java.util.function):
//
//   - Predicate[T], BiPredicate[T, U] - Test conditions
//   - Consumer[T], BiConsumer[T, U] - Perform operations
//   - Function[T, U], BiFunction[T, U, R] - Transform values
//   - Supplier[T] - Provide values
//   - UnaryOperator[T], BinaryOperator[T] - Operate on same type
//   - Comparator[T] - Compare values for ordering
//
// 2. Optional and Error Handling:
//
//   - Optional[T] - Optional value (Java-style: Of, Empty, OfNullable, IsPresent, OrElse)
//   - Result[T] - Error handling with Go's error type (Success, Failure, FromError)
//   - Either[L, R] - General binary choice between any two types (Left, Right, Swap)
//
// 3. Tuple Types:
//
//   - Pair[T, U], Triple[T, U, V], Quad[T, U, V, W] - Multi-value tuples
//
// Quick Examples:
//
//	// Functional types
//	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
//	fmt.Println(isEven(4))  // true
//
//	// Optional values (Java-style)
//	opt := lxtypes.Of(42)
//	value := opt.OrElse(0)  // 42
//
//	// Safe nil handling
//	var ptr *int
//	opt2 := lxtypes.OfNullable(ptr)  // Empty Optional
//	value2 := opt2.OrElse(99)        // 99
//
//	// Error handling with Result[T] (specialized for Go's error)
//	result := lxtypes.Success(42)
//	if result.IsSuccess() {
//	    fmt.Println(result.Value())  // 42
//	}
//	// Convert from Go's (value, error) pattern
//	result2 := lxtypes.FromError(strconv.Atoi("42"))
//
//	// General binary choice with Either[L, R]
//	either := lxtypes.Right[string, int](42)
//	if either.IsRight() {
//	    fmt.Println(either.Right())  // 42
//	}
//
//	// Tuples
//	p := lxtypes.NewPair(42, "answer")
//	fmt.Println(p.First, p.Second)  // 42 answer
//
// For comprehensive documentation, examples, and use cases, see:
// https://github.com/nthanhhai2909/lx/tree/main/lxtypes#readme
package lxtypes
