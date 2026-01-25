# langx

> Small, focused extensions to Go's standard library.

`langx` provides lightweight, idiomatic helper packages that complement Go’s standard library. Each package is intentionally small, well-tested, and easy to reason about.

## Features

* Go-idiomatic APIs
* Small, focused packages (no giant `utils`)
* Generics-first (Go 1.25+)
* Minimal dependencies
* Stable, predictable behavior

## Packages

* **`stringx`** – string helpers (blank checks, safe joins, truncation, casing)
* **`slicex`** – generic slice helpers (map, filter, unique)
* **`mapx`** – map helpers (safe get, defaults)
* **`timex`** – time helpers (parsing, formatting)
* **`errorx`** – error helpers (wrapping, joining)

> More packages may be added over time, but only when they provide clear value beyond the standard library.

## Installation

```bash
go get github.com/nthanhhai2909/langx
```

## Usage

```go
import "github.com/nthanhhai2909/langx/stringx"

if stringx.IsBlank(input) {
    return errors.New("input is blank")
}
```

```go
import "github.com/nthanhhai2909/langx/slicex"

ids := slicex.Unique([]int{1, 2, 2, 3})
```

## Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines and the contribution process.

## 📄 License

Apache License 2.0