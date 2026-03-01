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
//   - Optional[T] - Optional value (Java-style with comma-ok pattern: Get() returns (T, bool))
//   - Result[T] - Error handling with Go's (value, error) pattern (Value() returns (T, error))
//   - Either[L, R] - General binary choice between any two types (EitherLeft, EitherRight)
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
//	// Function composition
//	double := lxtypes.Function[int, int](func(n int) int { return n * 2 })
//	result := double.AndThen(func(n int) int { return n + 10 })(5)  // 20
//
//	// BiFunction composition
//	add := lxtypes.BiFunction[int, int, int](func(a, b int) int { return a + b })
//	result := add.AndThen(func(n int) int { return n * 2 })(3, 4)  // 14
//
//	// Optional values with comma-ok pattern (idiomatic Go)
//	opt := lxtypes.OptionalOf(42)
//	if value, ok := opt.Get(); ok {
//	    fmt.Println(value)  // 42
//	}
//
//	// Or use default values
//	value := opt.OrElse(0)  // 42
//
//	// Safe nil handling with OptionalOfNullable
//	var ptr *int
//	opt2 := lxtypes.OptionalOfNullable(ptr)  // Empty Optional
//	value2 := opt2.OrElse(99)                // 99
//
//	// Error handling with Result[T] using (value, error) pattern
//	result := lxtypes.ResultSuccess(42)
//	if value, err := result.Value(); err == nil {
//	    fmt.Println(value)  // 42
//	}
//
//	// Or use default value
//	failure := lxtypes.ResultFailure[int](errors.New("error"))
//	value := failure.ValueOr(99)  // 99
//
//	// General binary choice with Either[L, R]
//	either := lxtypes.EitherRight[string, int](42)
//	if right, ok := either.Right(); ok {
//	    fmt.Println(right)  // 42
//	}
//
//	// Tuples
//	p := lxtypes.NewPair(42, "answer")
//	fmt.Println(p.First, p.Second)  // 42 answer
//
// For comprehensive documentation, examples, and use cases, see:
// https://github.com/nthanhhai2909/lx/tree/main/lxtypes#readme
package lxtypes
