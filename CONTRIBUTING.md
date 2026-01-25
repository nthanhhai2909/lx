# Contributing

Thanks for your interest in **langx**

This project aims to provide small, focused, and idiomatic extensions to Go’s standard library. Please read the guidelines below to help keep the codebase clean, consistent, and easy to maintain.

---

## Contribute

1. **Fork** the repository
2. Create a feature branch from `main`:

   ```bash
   git checkout -b feat/short-description
   ```
3. Make your changes
4. Add or update tests
5. Run tests locally:

   ```bash
   go test ./...
   ```
6. Commit using clear, conventional messages
7. Open a Pull Request

---

## Branch naming

Use short-lived branches with the following prefixes:

* `feat/` – new features
* `fix/` – bug fixes
* `docs/` – documentation only
* `chore/` – tooling, CI, refactors

Examples:

```
feat/strx-isblank
fix/slicex-empty-input
docs/readme
```

---

## Commit message format

This project follows **Conventional Commits**:

```
feat(strx): add IsBlank helper
fix(slicex): handle nil slices
docs: update README
chore: update ci config
```

---

## Tests

Testing is a core requirement for this project.

* **All new functionality must include tests**
* Bug fixes should include **regression tests** that fail before the fix and pass after
* Prefer **table-driven tests**, following common Go conventions
* Tests should be:
  * Deterministic (no flaky or time-dependent behavior)
  * Easy to read and focused on one behavior per test case
* Place tests in the same package unless black-box testing is explicitly needed

Before opening a Pull Request, make sure:

```bash
go test ./...
```

passes locally without errors or warnings.
