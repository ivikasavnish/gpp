package main

import (
	"fmt"
	"time"
)

// timed wraps fn to measure its execution time.
// Signature: func(func()) func() — takes the function, returns a wrapped version.
func timed(fn func()) func() {
	return func() {
		start := time.Now()
		fn()
		fmt.Printf("  [timed: %v]\n", time.Since(start))
	}
}

// logged wraps fn to log calls by name.
// Curried: logged("name") returns a func(func()) func().
func logged(name string) func(func()) func() {
	return func(fn func()) func() {
		return func() {
			fmt.Printf("  [logged: → %s]\n", name)
			fn()
			fmt.Printf("  [logged: ← %s]\n", name)
		}
	}
}

// cached wraps fn with a cache stub.
// Curried: cached(size) returns a func(func()) func().
func cached(size int) func(func()) func() {
	return func(fn func()) func() {
		return func() {
			fmt.Printf("  [cached: size=%d]\n", size)
			fn()
		}
	}
}

// repeat wraps fn to call it n times.
// Curried: repeat(n) returns a func(func()) func().
func repeat(n int) func(func()) func() {
	return func(fn func()) func() {
		return func() {
			for i := 1; i <= n; i++ {
				fmt.Printf("  [repeat: run %d/%d]\n", i, n)
				fn()
			}
		}
	}
}

// retried wraps fn to retry on panic up to n times.
// Curried: retried(n) returns a func(func()) func().
func retried(n int) func(func()) func() {
	return func(fn func()) func() {
		return func() {
			for attempt := 1; attempt <= n; attempt++ {
				ok := func() (success bool) {
					defer func() {
						if r := recover(); r != nil {
							fmt.Printf("  [retried: attempt %d/%d failed → %v]\n", attempt, n, r)
						} else {
							success = true
						}
					}()
					fn()
					return true
				}()
				if ok {
					fmt.Printf("  [retried: succeeded on attempt %d/%d]\n", attempt, n)
					return
				}
			}
			fmt.Printf("  [retried: all %d attempts exhausted]\n", n)
		}
	}
}

// throttled wraps fn to enforce a minimum delay between calls.
// Curried: throttled(delay) returns a func(func()) func().
func throttled(delay time.Duration) func(func()) func() {
	return func(fn func()) func() {
		var lastCall time.Time
		return func() {
			if wait := delay - time.Since(lastCall); wait > 0 {
				fmt.Printf("  [throttled: waiting %v]\n", wait.Round(time.Millisecond))
				time.Sleep(wait)
			}
			lastCall = time.Now()
			fn()
		}
	}
}

// normalFunc has no decorators.
func normalFunc() {
	fmt.Println("  Normal function called")
}

// timedFunc is wrapped by timed — it will log elapsed time after the body.
@timed
func timedFunc() {
	fmt.Println("  Timed function called")
}

// loggedFunc is wrapped by logged — logs entry and exit.
@logged("myFunction")
func loggedFunc() {
	fmt.Println("  Logged function called")
}

// multiDecoratedFunc is wrapped by all three, outermost first:
//   timed( logged("multiDecorated")( cached(100)( func(){body} ) ) )()
@timed
@logged("multiDecorated")
@cached(100)
func multiDecoratedFunc() {
	fmt.Println("  Multi-decorated function called")
}

// repeatedFunc runs 100 times via the repeat decorator.
@repeat(100)
func repeatedFunc() {
	fmt.Println("  Repeated function called")
}

// flakyFunc simulates a function that fails on first 2 calls.
var callCount int

@retried(5)
func flakyFunc() {
	callCount++
	if callCount < 3 {
		panic(fmt.Sprintf("not ready yet (call %d)", callCount))
	}
	fmt.Println("  Flaky function succeeded")
}

// throttledFunc is called rapidly but throttled to 200ms between calls.
@logged("throttledFunc")
@throttled(200 * time.Millisecond)
@repeat(5)
func throttledFunc() {
	fmt.Println("  Throttled function called")
}

// retriedAndTimed combines retried + timed.
@timed
@retried(3)
func retriedAndTimed() {
	fmt.Println("  Retried+timed function called")
}

func main() {
	fmt.Println("=== Testing Go Decorator Syntax ===\n")

	fmt.Println("1. Normal function:")
	normalFunc()

	fmt.Println("\n2. @timed — wraps before AND after:")
	timedFunc()

	fmt.Println("\n3. @logged(\"myFunction\") — wraps with entry/exit log:")
	loggedFunc()

	fmt.Println("\n4. @timed @logged @cached — three layers, outermost=timed:")
	multiDecoratedFunc()

	fmt.Println("\n5. @repeat(100) — runs 100 times:")
	repeatedFunc()

	fmt.Println("\n6. @retried(5) — retries on panic:")
	callCount = 0
	flakyFunc()

	fmt.Println("\n7. @logged @throttled(200ms) @repeat(5) — stacked:")
	throttledFunc()

	fmt.Println("\n8. @timed @retried(3) — timed + retried:")
	retriedAndTimed()
}
