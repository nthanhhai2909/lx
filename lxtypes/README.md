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
- **Lazy evaluation**: Lazy[T] for deferred computation with caching

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

Represents a value that may or may not be present. Inspired by Java's `Optional<T>` but adapted to Go's idiomatic comma-ok pattern.

An `Optional` is either:
- **Present**: Contains a value (created with `OptionalOf` or `OptionalOfNullable`)
- **Empty**: Contains no value (created with `OptionalEmpty`)

#### Creating Optionals

```go
// Present value
opt := lxtypes.OptionalOf(42)

// Empty value
empty := lxtypes.OptionalEmpty[int]()

// From nullable pointer (safe nil handling)
value := 42
opt1 := lxtypes.OptionalOfNullable(&value)  // OptionalOf(42)

var nilPtr *int
opt2 := lxtypes.OptionalOfNullable(nilPtr)  // OptionalEmpty()
```

#### Checking Presence with Comma-Ok Pattern

```go
opt := lxtypes.OptionalOf(42)

// Use Go's idiomatic comma-ok pattern
if value, ok := opt.Get(); ok {
    fmt.Println("Has value:", value)  // Has value: 42
} else {
    fmt.Println("No value")
}
```

#### Working with Structs

```go
type User struct {
    Name  string
    Email string
    Age   int
}

// With struct values
user := User{Name: "Alice", Email: "alice@example.com", Age: 30}
opt := lxtypes.OptionalOf(user)

if u, ok := opt.Get(); ok {
    fmt.Printf("User: %s, Email: %s\n", u.Name, u.Email)
}

// With struct pointers
userPtr := &User{Name: "Bob", Email: "bob@example.com", Age: 25}
opt2 := lxtypes.OptionalOf(userPtr)

if u, ok := opt2.Get(); ok {
    fmt.Printf("User: %s\n", u.Name)
}

// Safe nil handling with OptionalOfNullable
var nilUser *User
opt3 := lxtypes.OptionalOfNullable(nilUser)  // Empty Optional

defaultUser := User{Name: "Guest", Email: "guest@example.com"}
user = opt3.OrElse(defaultUser)  // Returns defaultUser
```

#### Safe Access with Defaults

```go
// Static default
opt := lxtypes.OptionalOf(42)
value := opt.OrElse(0)  // 42

empty := lxtypes.OptionalEmpty[int]()
value2 := empty.OrElse(99)  // 99

// Computed default (lazy evaluation)
value3 := empty.OrElseGet(func() int {
    return expensiveComputation()
})
```

#### Practical Example: Database Lookup

```go
// Simulate a database lookup that might return nil
func FindUserByID(id int) *User {
    // Database lookup...
    if found {
        return &user
    }
    return nil
}

// Safe handling with Optional
userPtr := FindUserByID(123)
opt := lxtypes.OptionalOfNullable(userPtr)

// Use comma-ok pattern
if user, ok := opt.Get(); ok {
    fmt.Printf("Found: %s\n", user.Name)
} else {
    fmt.Println("User not found")
}

// Or use default
defaultUser := User{Name: "Guest", Email: "guest@example.com"}
user := opt.OrElse(defaultUser)
```

**Methods:**
- `Get() (T, bool)` - Returns (value, true) if present, or (zero, false) if empty (comma-ok pattern)
- `OrElse(T) T` - Get value or default
- `OrElseGet(func() T) T` - Get value or computed default (lazy evaluation)

**Use Cases:**
- Safe dictionary/map lookups
- Database query results that might not exist
- Configuration values that are optional
- Eliminating nil pointer panics
- API responses that may be absent
- File operations that may fail

### Result[T]

Represents the result of an operation that may succeed or fail. Adapted to Go's idiomatic (value, error) pattern.

A `Result` is either:
- **Success**: Contains a value (created with `ResultSuccess`)
- **Failure**: Contains an error (created with `ResultFailure`)

#### Creating Results

```go
// Success with value
result := lxtypes.ResultSuccess(42)

// Failure with error
result := lxtypes.ResultFailure[int](errors.New("operation failed"))
```

#### Checking Success with Value-Error Pattern

```go
result := divide(10, 2)

// Use Go's idiomatic (value, error) pattern
if value, err := result.Value(); err == nil {
    fmt.Println("Success:", value)
} else {
    fmt.Println("Error:", err)
}
```

#### Working with Structs

```go
type Config struct {
    Host string
    Port int
}

// With struct values
config := Config{Host: "localhost", Port: 8080}
result := lxtypes.ResultSuccess(config)

if cfg, err := result.Value(); err == nil {
    fmt.Printf("Config: %s:%d\n", cfg.Host, cfg.Port)
}

// With struct pointers
configPtr := &Config{Host: "example.com", Port: 443}
result2 := lxtypes.ResultSuccess(configPtr)

if cfg, err := result2.Value(); err == nil {
    fmt.Printf("Config: %s\n", cfg.Host)
}

// Failure with default
result3 := lxtypes.ResultFailure[Config](errors.New("config not found"))
defaultConfig := Config{Host: "default", Port: 80}
config = result3.ValueOr(defaultConfig)  // Returns defaultConfig
```

#### Safe Access with Defaults

```go
// Success returns original value
success := lxtypes.ResultSuccess(42)
value := success.ValueOr(0)  // 42

// Failure returns default value
failure := lxtypes.ResultFailure[int](errors.New("error"))
value2 := failure.ValueOr(99)  // 99
```

#### Practical Example: API Call

```go
// Simulate an API call
func FetchUser(id int) lxtypes.Result[User] {
    // Make API call...
    if err != nil {
        return lxtypes.ResultFailure[User](err)
    }
    return lxtypes.ResultSuccess(user)
}

// Handle the result
result := FetchUser(123)

// Use value-error pattern
if user, err := result.Value(); err == nil {
    fmt.Printf("User: %s\n", user.Name)
} else {
    fmt.Printf("Error: %v\n", err)
}

// Or use default
defaultUser := User{Name: "Guest"}
user := result.ValueOr(defaultUser)
```

#### Converting from Go's (value, error) Pattern

```go
// Standard library function
value, err := strconv.Atoi("42")

// Convert to Result
var result lxtypes.Result[int]
if err != nil {
    result = lxtypes.ResultFailure[int](err)
} else {
    result = lxtypes.ResultSuccess(value)
}

// Now use Result methods
finalValue := result.ValueOr(0)
```

**Methods:**
- `Value() (T, error)` - Returns (value, nil) if success, or (zero, err) if failure (Go's idiomatic pattern)
- `ValueOr(T) T` - Get value or default (no error checking needed)

**Use Cases:**
- Wrapping functions that return (value, error)
- Standardizing error handling across your codebase
- Avoiding nil pointer panics with default values
- Database operations that may fail
- File I/O operations
- Network requests
- Configuration loading
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
either := lxtypes.EitherLeft[string, int]("error")

// Right
either := lxtypes.EitherRight[string, int](42)
```

#### Accessing Values

```go
// Check and access Left
if left, ok := either.Left(); ok {
    fmt.Println("Left value:", left)
}

// Check and access Right
if right, ok := either.Right(); ok {
    fmt.Println("Right value:", right)
}

// Pattern matching style
if left, ok := either.Left(); ok {
    fmt.Println("Error:", left)
} else if right, ok := either.Right(); ok {
    fmt.Println("Success:", right)
}
```

#### Safe Access with Defaults

```go
// Get value or default
leftVal := either.LeftOr("default")
rightVal := either.RightOr(0)
```

**Methods:**
- `Left() (L, bool)` - Returns left value and true if Left, or zero value and false if Right
- `Right() (R, bool)` - Returns right value and true if Right, or zero value and false if Left
- `LeftOr(L) L` - Returns left value or default if Right
- `RightOr(R) R` - Returns right value or default if Left

**Use Cases:**
- Validation with custom error types
- Parsing that returns one of two types
- Union types before pattern matching
- Polymorphic return values
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
        return lxtypes.EitherLeft[ValidationError, User](
            ValidationError{Field: "age", Message: "must be non-negative"},
        )
    }
    if name == "" {
        return lxtypes.EitherLeft[ValidationError, User](
            ValidationError{Field: "name", Message: "cannot be empty"},
        )
    }
    return lxtypes.EitherRight[ValidationError, User](User{Name: name, Age: age})
}

// Usage
result := ValidateUser("Alice", 30)
if user, ok := result.Right(); ok {
    fmt.Printf("Valid user: %s, age %d\n", user.Name, user.Age)
} else if validationErr, ok := result.Left(); ok {
    fmt.Printf("Validation error in %s: %s\n", validationErr.Field, validationErr.Message)
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

## Lazy Evaluation

### Lazy[T]

Represents a value that may be computed immediately (eager) or on first access (deferred). Provides a unified interface for both strategies with automatic caching and thread-safe deferred evaluation.

**Key Features:**
- **Deferred computation**: Compute values only when needed
- **Automatic caching**: Results cached after first computation
- **Thread-safe**: Uses `sync.Once` for safe concurrent access
- **Error handling**: Supports computations that may fail
- **Unified interface**: Treat eager and deferred values uniformly

#### Creating Lazy Values

##### LazyEager - Immediate Value

Wraps an already-computed value with no deferred computation.

```go
// Wrap an immediate value
lazy := lxtypes.LazyEager(42)
value, _ := lazy.Get()  // Returns immediately: 42

// Check if evaluated (always true for eager)
fmt.Println(lazy.IsEvaluated())  // true
```

##### LazyEagerOrError - Wrap Existing Results

Convert an existing `(value, error)` result into a Lazy.

```go
result, err := someFunction()
lazy := lxtypes.LazyEagerOrError(result, err)

// Later access
value, err := lazy.Get()  // Returns result and err immediately
```

##### LazyDeferred - Deferred Computation

Compute value lazily on first access. The computation function is called at most once.

```go
// Create lazy computation
expensive := lxtypes.LazyDeferred(func() (int, error) {
    // Expensive operation (database query, file read, etc.)
    time.Sleep(time.Second)
    return 42, nil
})

fmt.Println(expensive.IsEvaluated())  // false (not computed yet)

// Compute on first access
value, err := expensive.Get()  // Takes 1 second
fmt.Println(value)                    // 42

// Subsequent calls return cached value
value2, _ := expensive.Get()  // Returns instantly
fmt.Println(expensive.IsEvaluated())  // true
```

#### Methods

##### Get() (T, error)

Returns the value, computing it if necessary. For deferred values, the computation happens on first call and the result is cached.

```go
lazy := lxtypes.LazyDeferred(func() (string, error) {
    return "computed", nil
})

value, err := lazy.Get()
if err != nil {
    // Handle error
}
fmt.Println(value)  // "computed"
```

##### MustGet() T

Returns the value, panicking if computation fails. Useful when you're certain the computation will succeed.

```go
lazy := lxtypes.LazyDeferred(func() (int, error) {
    return 42, nil
})

value := lazy.MustGet()  // 42 (panics if error)
```

##### IsEvaluated() bool

Returns true if the value has been computed (for deferred) or was provided immediately (for eager).

```go
lazy := lxtypes.LazyDeferred(func() (int, error) {
    return 42, nil
})

fmt.Println(lazy.IsEvaluated())  // false
lazy.Get()
fmt.Println(lazy.IsEvaluated())  // true
```

#### Use Cases

**1. Expensive Resource Initialization**

```go
// Database connection pool that's only created when needed
dbPool := lxtypes.LazyDeferred(func() (*sql.DB, error) {
    return sql.Open("postgres", connectionString)
})

// In fast path, database never initialized
if cachedResult, ok := cache.Get(key); ok {
    return cachedResult
}

// In slow path, database initialized on first access
db, err := dbPool.Get()
if err != nil {
    return err
}
result := db.Query(...)
```

**2. Configuration Loading**

```go
type Config struct {
    APIKey string
    Endpoint string
}

// Config only loaded when actually needed
config := lxtypes.LazyDeferred(func() (Config, error) {
    return loadConfigFromFile("config.json")
})

// Conditional usage
if needsAPI {
    cfg := config.MustGet()
    callAPI(cfg.APIKey, cfg.Endpoint)
}
```

**3. Conditional Computations**

```go
// Expensive fallback that's only computed if primary fails
fallbackData := lxtypes.LazyDeferred(func() ([]byte, error) {
    return fetchFromSlowBackup()
})

// Try primary source
data, err := fetchFromPrimary()
if err != nil {
    // Fallback computation only happens here
    data, err = fallbackData.Get()
}
```

**4. Thread-Safe Singletons**

```go
var logger = lxtypes.LazyDeferred(func() (*Logger, error) {
    return NewLogger("app.log")
})

// Safe to call from multiple goroutines
func Log(msg string) {
    log := logger.MustGet()  // Initialized exactly once
    log.Write(msg)
}
```

**5. Caching Expensive Computations**

```go
// Compute once, reuse many times
fibonacci := lxtypes.LazyDeferred(func() ([]int, error) {
    // Generate first 100 fibonacci numbers
    fib := make([]int, 100)
    fib[0], fib[1] = 0, 1
    for i := 2; i < 100; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }
    return fib, nil
})

// Use multiple times without recomputation
for i := 0; i < 10; i++ {
    fib, _ := fibonacci.Get()  // Only computed once
    fmt.Println(fib[i])
}
```

#### Error Handling

Lazy supports operations that may fail. Errors are cached along with values.

```go
// Computation that may fail
lazy := lxtypes.LazyDeferred(func() (Data, error) {
    data, err := loadData()
    if err != nil {
        return Data{}, fmt.Errorf("failed to load: %w", err)
    }
    return data, nil
})

// Handle error on access
data, err := lazy.Get()
if err != nil {
    log.Printf("Error: %v", err)
    return
}

// Or use MustGet if failure should panic
data := lazy.MustGet()
```

#### Thread Safety

`LazyDeferred` is thread-safe and ensures the computation function is called exactly once, even with concurrent access.

```go
lazy := lxtypes.LazyDeferred(func() (int, error) {
    time.Sleep(100 * time.Millisecond)
    return 42, nil
})

// Launch 100 concurrent goroutines
var wg sync.WaitGroup
for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        value, _ := lazy.Get()  // Function called exactly once
        fmt.Println(value)
    }()
}
wg.Wait()
```

#### Eager vs Deferred Comparison

| Feature | LazyEager | LazyDeferred |
|---------|-----------|--------------|
| Computation timing | Immediate | On first `Get()` |
| Performance | Instant access | First access has delay |
| Use case | Wrapping existing values | Expensive computations |
| `IsEvaluated()` | Always `true` | `false` until first `Get()` |
| Thread safety | N/A (immutable) | Thread-safe with `sync.Once` |
| Memory | Stores value immediately | Stores function until evaluation |

**When to use LazyEager:**
- Wrapping existing computation results
- Converting `(value, error)` pairs to Lazy interface
- No benefit from deferring (value already available)

**When to use LazyDeferred:**
- Expensive computations (database queries, file I/O, API calls)
- Conditional usage (may not be needed in all code paths)
- Singleton initialization
- Caching computed values across multiple accesses

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
- ✅ Simple (value, bool) pattern like Optional
- ✅ Type-safe access with boolean checks
- ✅ Polymorphic return values
- ✅ Union types representation
- ✅ No panics - always safe to access

## Related Packages

- **lxslices**: Functional slice operations (uses Predicate types)
- **lxptrs**: Pointer utilities
- **lxconstraints**: Generic type constraints

## Contributing

See the main [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.

## License

Apache 2.0 - See [LICENSE](../LICENSE) for details.

