# lxtypes

Generic type definitions for common functional programming patterns in Go.

## Overview

The `lxtypes` package provides reusable, type-safe generic type definitions inspired by functional programming patterns. These types enable cleaner, more composable code when working with functions as values.

## Installation

```bash
go get github.com/nthanhhai2909/lx/lxtypes
```

## Type Definitions

### Basic Functional Types

#### Predicate[T]
Tests a condition on an input value.

```go
isPositive := lxtypes.Predicate[int](func(n int) bool {
    return n > 0
})

fmt.Println(isPositive(5))   // true
fmt.Println(isPositive(-3))  // false
```

**Methods:**
- `And(other Predicate[T]) Predicate[T]` - Logical AND composition
- `Or(other Predicate[T]) Predicate[T]` - Logical OR composition
- `Negate() Predicate[T]` - Logical negation

#### Consumer[T]
Performs an operation on an input without returning a result.

```go
printValue := lxtypes.Consumer[string](func(s string) {
    fmt.Println(s)
})

printValue("Hello, World!")
```

**Methods:**
- `AndThen(after Consumer[T]) Consumer[T]` - Sequential composition

#### Function[T, U]
Transforms an input of type T to an output of type U.

```go
toUpper := lxtypes.Function[string, string](func(s string) string {
    return strings.ToUpper(s)
})

result := toUpper("hello")  // "HELLO"
```

#### Supplier[T]
Provides a value with no input.

```go
counter := 0
getNext := lxtypes.Supplier[int](func() int {
    counter++
    return counter
})

fmt.Println(getNext())  // 1
fmt.Println(getNext())  // 2
```

### Binary Types

#### BiPredicate[T, U]
Tests a condition on two input values.

```go
equals := lxtypes.BiPredicate[int, int](func(a, b int) bool {
    return a == b
})

fmt.Println(equals(5, 5))  // true
```

**Methods:**
- `And(other BiPredicate[T, U]) BiPredicate[T, U]` - Logical AND composition
- `Or(other BiPredicate[T, U]) BiPredicate[T, U]` - Logical OR composition
- `Negate() BiPredicate[T, U]` - Logical negation

#### BiConsumer[T, U]
Performs an operation on two inputs without returning a result.

```go
printPair := lxtypes.BiConsumer[string, int](func(s string, n int) {
    fmt.Printf("%s: %d\n", s, n)
})

printPair("count", 42)  // count: 42
```

#### BiFunction[T, U, R]
Transforms two inputs to a single output.

```go
concat := lxtypes.BiFunction[string, string, string](func(a, b string) string {
    return a + b
})

result := concat("Hello, ", "World!")  // "Hello, World!"
```

### Operator Types

#### UnaryOperator[T]
Transforms a value to the same type.

```go
square := lxtypes.UnaryOperator[int](func(n int) int {
    return n * n
})

result := square(5)  // 25
```

#### BinaryOperator[T]
Combines two values of the same type.

```go
sum := lxtypes.BinaryOperator[int](func(a, b int) int {
    return a + b
})

result := sum(3, 7)  // 10
```

### Comparison

#### Comparator[T]
Compares two values for ordering.

```go
intComparator := lxtypes.Comparator[int](func(a, b int) int {
    if a < b {
        return -1
    }
    if a > b {
        return 1
    }
    return 0
})

result := intComparator(3, 5)  // -1 (3 < 5)
```

**Methods:**
- `Reversed() Comparator[T]` - Reverse the ordering
- `ThenComparing(other Comparator[T]) Comparator[T]` - Lexicographic comparison

## Utility Functions

### Compose[T, U, V]
Composes two functions, applying the first then the second.

```go
double := func(n int) int { return n * 2 }
addTen := func(n int) int { return n + 10 }

// Compose: addTen first, then double
// Compose(addTen, double)(5) = double(addTen(5)) = double(15) = 30
addTenThenDouble := lxtypes.Compose(addTen, double)

result := addTenThenDouble(5)  // 30
```

## Advanced Examples

### Combining Predicates

```go
isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
isPositive := lxtypes.Predicate[int](func(n int) bool { return n > 0 })

// Combine with And
isEvenAndPositive := isEven.And(isPositive)
fmt.Println(isEvenAndPositive(4))   // true
fmt.Println(isEvenAndPositive(-2))  // false

// Combine with Or
isEvenOrPositive := isEven.Or(isPositive)
fmt.Println(isEvenOrPositive(3))   // true
fmt.Println(isEvenOrPositive(-2))  // true

// Negate
isOdd := isEven.Negate()
fmt.Println(isOdd(3))  // true
```

### Combining BiPredicates

```go
inRange := lxtypes.BiPredicate[int, int](func(value, max int) bool {
    return value >= 0 && value <= max
})
lessThan := lxtypes.BiPredicate[int, int](func(a, b int) bool {
    return a < b
})

// Combine with And
validAndLess := inRange.And(lessThan)
fmt.Println(validAndLess(5, 10))  // true (in range AND less)
fmt.Println(validAndLess(10, 5))  // false (in range but NOT less)

// Combine with Or
equals := lxtypes.BiPredicate[int, int](func(a, b int) bool { return a == b })
validOrEqual := inRange.Or(equals)
fmt.Println(validOrEqual(5, 10))   // true (in range)
fmt.Println(validOrEqual(15, 15))  // true (equal)
fmt.Println(validOrEqual(20, 10))  // false (neither)

// Negate
notEquals := equals.Negate()
fmt.Println(notEquals(5, 3))  // true
```

### Sequential Consumers

```go
results := []string{}

append1 := lxtypes.Consumer[string](func(s string) {
    results = append(results, s)
})
append2 := lxtypes.Consumer[string](func(s string) {
    results = append(results, strings.ToUpper(s))
})

combined := append1.AndThen(append2)
combined("hello")

fmt.Println(results)  // [hello HELLO]
```

### Complex Comparators

```go
type Person struct {
    Name string
    Age  int
}

byName := lxtypes.Comparator[Person](func(a, b Person) int {
    if a.Name < b.Name {
        return -1
    }
    if a.Name > b.Name {
        return 1
    }
    return 0
})

byAge := lxtypes.Comparator[Person](func(a, b Person) int {
    return a.Age - b.Age
})

// Sort by name, then by age
byNameThenAge := byName.ThenComparing(byAge)

// Sort by age descending
byAgeDesc := byAge.Reversed()
```

## Design Philosophy

This package follows the lx project's core principles:

- **Type-Safe**: Full generic type support for compile-time safety
- **Composable**: Types support method chaining and composition
- **Zero Dependencies**: Uses only Go standard library
- **Idiomatic Go**: Follows Go conventions while providing functional patterns

## Use Cases

- **Higher-order functions**: Pass behavior as parameters
- **Filter/Map/Reduce operations**: Use with lxslices for functional transformations
- **Strategy pattern**: Inject different behaviors without interfaces
- **Composition**: Build complex operations from simple ones
- **Sorting**: Custom comparison logic with Comparator

## Related Packages

- **lxslices**: Functional slice operations (uses Predicate types)
- **lxptrs**: Pointer utilities
- **lxconstraints**: Generic type constraints

## Contributing

See the main [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.

## License

Apache 2.0 - See [LICENSE](../LICENSE) for details.

