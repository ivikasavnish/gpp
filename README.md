# gpp — Go with Decorator Syntax and Nil-Safe Selector

A custom Go compiler that adds Python-style `@decorator` syntax and a
Kotlin/JS-style nil-safe selector (`?.`) to Go.

## Requirements

- Linux AMD64
- `curl`, `tar`

## Quick Start

```bash
git clone https://github.com/ivikasavnish/gpp
cd gpp
./install.sh
./gpp run test_decorator_syntax.go
./gpp run test_nil_safety.go
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

Known limitation: silently no-ops on methods (functions with a receiver)
rather than erroring.

## Nil-Safe Selector (`?.`)

Kotlin/JS-style chaining, not Go's usual `(value, ok)` idiom:

```go
var user *User
city := user?.Address?.City   // "" — no panic, single value

street := user2?.Address?.Street?.Name  // chains compose left to right
```

`a?.b?.c` evaluates to the zero value of the final field's type if any
pointer/interface link in the chain is nil, otherwise the real value.
Restricted to pointer/interface base types — `?.` on a plain value type
is a compile error, not silently ignored. See `test_nil_safety.go`.
