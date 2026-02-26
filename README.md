# gpp — Go with Decorator Syntax

A custom Go compiler that adds Python-style `@decorator` syntax to Go.

## Requirements

- Linux AMD64
- `curl`, `tar`

## Quick Start

```bash
git clone https://github.com/ivikasavnish/gpp
cd gpp
./install.sh
./gpp run test_decorator_syntax.go
```

## Decorator Syntax

```go
// simple decorator
@timed
func myFunc() { ... }

// parameterised decorator
@logged("name")
func myFunc() { ... }

// stacked decorators (outermost first)
@timed
@logged("name")
@repeat(100)
func myFunc() { ... }
```

## Available Decorators (in test file)

| Decorator | Description |
|---|---|
| `@timed` | prints elapsed time after the call |
| `@logged("name")` | prints entry/exit with name |
| `@cached(size)` | cache stub |
| `@repeat(n)` | calls the function n times |
| `@retried(n)` | retries on panic up to n times |
| `@throttled(delay)` | enforces minimum delay between calls |

## Decorator Signature

```go
// no params → func(func()) func()
func timed(fn func()) func() { ... }

// with params → curried: func(...) func(func()) func()
func repeat(n int) func(func()) func() { ... }
```
