# 10 — Testing concurrent code

Concurrent code is exactly the kind of code where "it worked when I ran it"
means very little — races and deadlocks are often timing-dependent and
won't show up every run. This lesson is about the tools that make
concurrent code trustworthy instead of "probably fine."

## The race detector: `-race`

```sh
go test -race ./...
go run -race ./10-testing-and-race
```

`-race` instruments every memory read/write and every synchronization
event, and reports if two goroutines touched the same memory
concurrently with at least one write and no `happens-before` relationship
between them (i.e. nothing — no mutex, no channel — established an order
between the two accesses). It reports the exact two goroutine stacks
involved. This is not a linter guessing at problems — it's dynamic
instrumentation of the exact execution that ran, so it only catches races
that the code path you exercised actually triggered. **Always run your
concurrency tests with `-race` in CI.** `main.go`'s `RacyAdder` is a
demo you can run through `go vet`-style manual inspection and then
actually catch with `-race`.

If you've used C/C++, this is the same underlying technology as Clang/
GCC's ThreadSanitizer (`-fsanitize=thread`) — Go's race detector was built
directly on that project's runtime library. Java has nothing built into
`javac`/the JVM that catches races this precisely; the closest tools are
external (e.g. IBM's old ConTest, or just careful code review), which is
part of why data races historically stayed hidden longer in Java codebases
than they do in a Go codebase that runs `-race` in CI.

## Don't synchronize tests with `time.Sleep`

```go
go doWork(resultCh)
time.Sleep(100 * time.Millisecond) // fragile: guessing how long doWork takes
result := <-resultCh
```

This is flaky by construction — too short and the test fails on a slow CI
box; too long and your test suite crawls. Use the actual synchronization
primitive instead: block on the channel or `WaitGroup` directly. If you
need a timeout for safety (so a genuinely broken test doesn't hang the
suite forever), use `select` with `time.After`, not a fixed sleep-then-check.

```go
select {
case result := <-resultCh:
    // use result
case <-time.After(time.Second):
    t.Fatal("timed out waiting for result")
}
```

## `testing/synctest`: deterministic concurrency tests (Go 1.25+)

The previous section says "don't synchronize tests with `time.Sleep`" —
but sometimes the thing you're testing is genuinely time-based (a
timeout, a retry backoff), and shrinking real durations down to
milliseconds to keep the test suite fast just trades one flakiness
problem for another (a slow CI box blowing through a 10ms budget).

`synctest.Test` runs a test function inside an isolated "bubble" with a
**fake clock**. Inside the bubble, `time.Sleep`, timers, and context
deadlines all use virtual time, which jumps forward instantly the moment
every goroutine in the bubble is blocked waiting on something (a channel,
a timer, `sync.WaitGroup.Wait`, ...). The result: a test that "waits" for
a 5-second timeout runs in a few milliseconds of real time, and does so
**deterministically** — no race against the CI box's actual clock speed.

```go
func TestWaitForSignalTimesOut(t *testing.T) {
    synctest.Test(t, func(t *testing.T) {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        sig := make(chan struct{}) // nobody ever closes this
        err := waitForSignal(ctx, sig)

        if err != context.DeadlineExceeded {
            t.Fatalf("err = %v, want DeadlineExceeded", err)
        }
    })
}
```

`synctest.Wait()` is the other half: it blocks until every other goroutine
in the bubble is durably blocked, which lets you assert "the goroutine I
started has definitely finished its work" without a sleep or an
extra channel round-trip purely for synchronization. See
`synctest_test.go` in this lesson for both patterns running against the
same `waitForSignal` helper (the same `select` on `ctx.Done()` shape from
lessons 05 and 07).

One constraint worth knowing up front: everything a bubbled test touches
should be self-contained — real network I/O, goroutines started outside
the bubble, and anything else that can only be unblocked by something
outside the bubble will make `Test` hang/panic rather than resolve
instantly, since the fake clock only advances when *every* goroutine in
the bubble is blocked on something the bubble itself controls.

## `t.Parallel()`

Marks a test as safe to run concurrently with other parallel tests in the
same package (Go still runs non-parallel tests sequentially first, then
runs all parallel ones together). Use it for independent tests to speed up
the suite — but never for tests that share mutable state, or you've just
introduced the exact race you're supposed to be testing for. Comparable to
JUnit 5's `@Execution(CONCURRENT)` or TestNG's parallel methods — same
tradeoff applies there too (shared fixtures become a hazard, not a
convenience).

Go's built-in `testing` package (`go test`, `*_test.go`, `func TestXxx(t
*testing.T)`) plays the role JUnit plays for Java or PHPUnit for PHP —
it's just part of the standard toolchain instead of a separate dependency
you add, and there's no separate assertion library bundled in (no
`assertEquals` — you write the `if got != want { t.Errorf(...) }` yourself,
or reach for a third-party assertion library if you want one).

## Benchmarking concurrent code

```go
func BenchmarkCounterInc(b *testing.B) {
    c := NewCounter()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            c.Inc()
        }
    })
}
```

`b.RunParallel` runs the benchmark body across multiple goroutines,
which is what you want when benchmarking something meant to be called
concurrently (like a mutex-protected counter) — a single-goroutine
benchmark wouldn't tell you anything about lock contention.

### `testing.B.Loop` (Go 1.24+)

For an ordinary, single-goroutine benchmark (no `b.RunParallel`), Go 1.24
replaced the classic pattern:

```go
// old:
for i := 0; i < b.N; i++ {
    doWork()
}

// Go 1.24+:
for b.Loop() {
    doWork()
}
```

`b.Loop()` resets the timer the moment it's first called (so setup code
above the loop doesn't count), stops the timer once it returns `false`
(so cleanup code below the loop doesn't count either), and keeps the
compiler from optimizing away the loop body — all things the old `b.N`
pattern made you get right by hand. It only replaces the *outer*, plain
loop; `b.RunParallel`/`pb.Next()` is still the right (and different) tool
the moment you want multiple goroutines calling the benchmarked code at
once, which is most of what this lesson cares about. `main_test.go` has
both: `BenchmarkSafeAdderInc` (parallel, contention-focused) and
`BenchmarkSafeAdderAddSequential` (single-goroutine, `b.Loop()`).

Run it:

```sh
go run ./10-testing-and-race
go test -race ./10-testing-and-race/...
go test -bench=. ./10-testing-and-race/...
```

## Real-world use cases

- **`-race` in CI is the single highest-leverage tool in this lesson** —
  most real Go projects (the standard library itself included) require
  it as a merge gate specifically because concurrency bugs are
  timing-dependent and can pass a normal test run hundreds of times
  before showing up in production under real load.
- **`testing/synctest` is what makes testing a retry-with-backoff client
  or a session-expiry timeout fast and reliable** — without it, teams
  either shrink real durations for tests (flaky on a loaded CI box) or
  skip testing the timeout path at all; a fake-clock test runs the exact
  same logic in milliseconds, deterministically, every time.
- **Benchmarking with `b.RunParallel`** is how you'd validate a
  performance fix before shipping it — e.g. proving that switching a
  hot-path counter from `sync.Mutex` to `atomic.Int64` actually reduced
  contention, with a number to put in the PR description instead of just
  asserting it should be faster.

## Katas

See [katas/README.md](katas/README.md).

```sh
go test -race ./10-testing-and-race/katas/...
go test -bench=. ./10-testing-and-race/katas/...
```
