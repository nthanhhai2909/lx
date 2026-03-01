# lxtypes

Generic type definitions for functional programming, optional values, and error handling in Go.

## Overview

The `lxtypes` package provides reusable, type-safe generic type definitions inspired by functional programming patterns from Java and Rust. These types enable cleaner, more composable code when working with functions, optional values, and error handling.

**Includes:**
- **Functional types**: Predicate, Consumer, Function, Supplier, and their binary variants
- **Optional values**: Optional[T] (Java-style Optional)
- **Error handling**: Result[T] (specialized for Go's error type)
- **Binary choice**: Either[L, R] (general-purpose union type)
- **Tuple types**: Pair, Triple, Quad for multi-value returns

## Installation

Install the lx module (includes all packages):

```bash
go get github.com/nthanhhai2909/lx
```

Then import only what you need:

```go
import "github.com/nthanhhai2909/lx/lxtypes"
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

## Optional and Error Handling Types

### Optional[T]

Represents a value that may or may not be present. Inspired by Java's `Optional<T>`.

An `Optional` is either:
- **Present**: Contains a value (created with `Of` or `OfNullable`)
- **Empty**: Contains no value (created with `Empty`)

#### Creating Optionals

```go
// Present value
opt := lxtypes.Of(42)

// Empty value
empty := lxtypes.Empty[int]()

// From nullable pointer (safe nil handling)
value := 42
opt1 := lxtypes.OfNullable(&value)  // Of(42)

var nilPtr *int
opt2 := lxtypes.OfNullable(nilPtr)  // Empty()
```

#### Checking Presence

```go
opt := lxtypes.Of(42)

if opt.IsPresent() {
    fmt.Println("Has value:", opt.Get())
}

if opt.IsEmpty() {
    fmt.Println("No value")
}
```

#### Safe Access with Defaults

```go
// Static default
value := opt.OrElse(0)

// Computed default
value := opt.OrElseGet(func() int {
    return computeDefault()
})

// Fallback to another Optional
result := opt.Or(lxtypes.Of(99))
```

**Methods:**
- `IsPresent() bool` - Check if value exists
- `IsEmpty() bool` - Check if empty
- `Get() T` - Get value (panics if empty)
- `OrElse(T) T` - Get value or default
- `OrElseGet(func() T) T` - Get value or computed default
- `Or(Optional[T]) Optional[T]` - Fallback to another Optional
- `OrElseSupply(func() Optional[T]) Optional[T]` - Computed fallback Optional


**Use Cases:**
- Safe dictionary/map lookups
- Database query results that might not exist
- Configuration values that are optional
- Eliminating nil pointer exceptions

### Result[T]

Represents the result of an operation that may succeed with a value or fail with an error. 
Specialized for Go's `error` type.

A `Result` is either:
- **Success**: Contains a value (created with `Success`)
- **Failure**: Contains an error (created with `Failure`)

#### Creating Results

```go
// Success
result := lxtypes.Success(42)

// Failure
result := lxtypes.Failure[int](errors.New("something went wrong"))

// Convert from Go's (value, error) pattern
value, err := strconv.Atoi("42")
result := lxtypes.FromError(value, err)  // Success(42)
```

#### Checking Success

```go
result := divide(10, 2)

if result.IsSuccess() {
    fmt.Println("Success:", result.Value())
} else {
    fmt.Println("Error:", result.Error())
}
```

#### Safe Access

```go
// With default
value := result.ValueOr(0)

// With computed default
value := result.ValueOrElse(func(err error) int {
    log.Println("Error:", err)
    return 0
})
```

**Methods:**
- `IsSuccess() bool` - Check if successful
- `IsFailure() bool` - Check if error
- `Value() T` - Get success value (panics if failure)
- `ValueOr(T) T` - Get value or default
- `ValueOrElse(func(error) T) T` - Get value or computed default
- `Error() error` - Get error value (panics if success)
- `OrElse(func(error) Result[T]) Result[T]` - Error recovery

**Standalone Functions:**
- `FromError[T](value T, err error) Result[T]` - Convert from Go's (value, error) pattern

**Use Cases:**
- Wrapping Go's standard library functions
- Error handling without exceptions
- Railway-oriented programming
- Chainable error propagation

### Either[L, R]

Represents a value of one of two possible types (a disjoint union). 
For general-purpose binary choice where both alternatives are valid values.

An `Either` is either:
- **Left**: Contains a left value (often used for errors/failures)
- **Right**: Contains a right value (often used for success/normal cases)

**Note**: Both Left and Right are equally valid - this is different from Result which is specifically for error handling with Go's `error` type.

#### Creating Eithers

```go
// Left
either := lxtypes.Left[string, int]("error")

// Right
either := lxtypes.Right[string, int](42)
```

#### Checking Which Side

```go
either := parseValue("42")

if either.IsRight() {
    fmt.Println("Number:", either.Right())
} else {
    fmt.Println("String:", either.Left())
}
```

#### Safe Access

```go
// With defaults
leftVal := either.LeftOr("default")
rightVal := either.RightOr(0)

// Swap sides
swapped := either.Swap()  // Either[int, string]
```

**Methods:**
- `IsLeft() bool` - Check if Left
- `IsRight() bool` - Check if Right
- `Left() L` - Get left value (panics if Right)
- `Right() R` - Get right value (panics if Left)
- `LeftOr(L) L` - Get left or default
- `RightOr(R) R` - Get right or default
- `Swap() Either[R, L]` - Swap Left and Right


**Use Cases:**
- Validation with custom error types
- Parsing that returns one of two types
- Union types before pattern matching
- Polymorphic return values

**Real-World Example**:
```go
type ValidationError struct {
    Field   string
    Message string
}

type User struct {
    Name string
    Age  int
}

func ValidateUser(name string, age int) lxtypes.Either[ValidationError, User] {
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
```

### When to Use Which?

| Type | Use When | Example |
|------|----------|---------|
| **Optional[T]** | Value might be absent | `FindUser(id) Optional[User]` |
| **Result[T]** | Operation might fail with Go's error | `ReadFile(path) Result[[]byte]` |
| **Either[L, R]** | Need custom error types or general binary choice | `Either[ValidationError, User]` |


## Tuple Types

### Pair[T, U]
A generic two-element tuple.

```go
// Create a pair
p := lxtypes.NewPair(42, "answer")
fmt.Printf("First: %d, Second: %s\n", p.First, p.Second)

// Unpack values
x, y := p.Values()

// Swap elements
swapped := p.Swap()  // Pair[string, int]

// Transform elements
doubled := p.MapFirst(func(n int) int { return n * 2 })
upper := p.MapSecond(func(s string) string { return strings.ToUpper(s) })
```

**Methods:**
- `Values() (T, U)` - Unpack the pair into separate values
- `Swap() Pair[U, T]` - Create a new pair with swapped elements
- `MapFirst(func(T) T) Pair[T, U]` - Transform the first element
- `MapSecond(func(U) U) Pair[T, U]` - Transform the second element

**Use Cases:**
- Return multiple values from functions
- Zip operations (combine two slices)
- Key-value pairs
- Coordinate pairs

### Triple[T, U, V]
A generic three-element tuple.

```go
// Create a triple
t := lxtypes.NewTriple(1, "hello", true)
fmt.Printf("Values: %d, %s, %t\n", t.First, t.Second, t.Third)

// Unpack values
x, y, z := t.Values()

// Convert to pair (drops third element)
p := t.ToPair()
```

**Methods:**
- `Values() (T, U, V)` - Unpack the triple into separate values
- `ToPair() Pair[T, U]` - Convert to a Pair, discarding the third element

**Use Cases:**
- RGB color values
- 3D coordinates
- Database query results with multiple columns
- Function results with status, value, and metadata

### Quad[T, U, V, W]
A generic four-element tuple.

```go
// Create a quad
q := lxtypes.NewQuad(1, "hello", true, 3.14)
fmt.Printf("Values: %d, %s, %t, %.2f\n", q.First, q.Second, q.Third, q.Fourth)

// Unpack values
w, x, y, z := q.Values()

// Convert to smaller tuples
p := q.ToPair()    // Pair with first two elements
t := q.ToTriple()  // Triple with first three elements
```

**Methods:**
- `Values() (T, U, V, W)` - Unpack the quad into separate values
- `ToPair() Pair[T, U]` - Convert to a Pair, discarding last two elements
- `ToTriple() Triple[T, U, V]` - Convert to a Triple, discarding the fourth element

**Use Cases:**
- RGBA color values
- Complex return values
- Database rows with multiple typed columns
- Configuration tuples

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
- **Industry Standards**: Inspired by Java's Optional and Rust's Result

## Use Cases

- **Higher-order functions**: Pass behavior as parameters
- **Filter/Map/Reduce operations**: Use with lxslices for functional transformations
- **Optional values**: Replace nil pointers and sentinel values with type-safe Option
- **Error handling**: Type-safe error handling with Result instead of exceptions
- **Safe dictionary lookups**: Return Option instead of value + bool
- **Strategy pattern**: Inject different behaviors without interfaces
- **Composition**: Build complex operations from simple ones
- **Sorting**: Custom comparison logic with Comparator
- **Null safety**: Eliminate null pointer exceptions with OfNullable

## Key Features

### Optional[T] Benefits
- ✅ Explicit presence/absence in type system
- ✅ Safe nil handling with OfNullable
- ✅ No more nil pointer panics
- ✅ Clear API: IsPresent, IsEmpty, Get, OrElse
- ✅ Flexible fallback options: Or, OrElseGet, OrElseSupply

### Result[T] Benefits
- ✅ Specialized for Go's error type (simpler API)
- ✅ Chainable error propagation with OrElse
- ✅ No exceptions - explicit error handling
- ✅ FromError helper for Go's (value, error) pattern
- ✅ Railway-oriented programming pattern
- ✅ Safe value access with ValueOr and ValueOrElse

### Either[L, R] Benefits
- ✅ General binary choice between any two types
- ✅ Custom error types with rich data
- ✅ Validation with detailed error information
- ✅ Polymorphic return values
- ✅ Union types representation
- ✅ Swap operation for flexibility

## Related Packages

- **lxslices**: Functional slice operations (uses Predicate types)
- **lxptrs**: Pointer utilities
- **lxconstraints**: Generic type constraints

## Contributing

See the main [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.

## License

Apache 2.0 - See [LICENSE](../LICENSE) for details.

