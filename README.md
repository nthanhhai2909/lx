# lx

> Small, focused extensions to Go's standard library.

`lx` provides lightweight, idiomatic helper packages that complement Go’s standard library. Each package is intentionally small, well-tested, and easy to reason about.

## Features

* Go-idiomatic APIs
* Small, focused packages (no giant `utils`)
* Generics-first (Go 1.25+)
* Minimal dependencies
* Stable, predictable behavior

## Packages

* **`lxstrings`** – string helpers
* **`lxptrs`** – pointer helpers

> More packages may be added over time, but only when they provide clear value beyond the standard library.

## Installation

```bash
go get github.com/nthanhhai2909/lx
```

## Usage

```go
import "github.com/nthanhhai2909/lx/lxstrings"

if lxstrings.IsBlank(input) {
    return errors.New("input is blank")
}
```


## Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines and the contribution process.

## 📄 License

Apache License 2.0