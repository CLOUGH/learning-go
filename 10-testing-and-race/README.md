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

## `t.Parallel()`

Marks a test as safe to run concurrently with other parallel tests in the
same package (Go still runs non-parallel tests sequentially first, then
runs all parallel ones together). Use it for independent tests to speed up
the suite — but never for tests that share mutable state, or you've just
introduced the exact race you're supposed to be testing for.

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

Run it:

```sh
go run ./10-testing-and-race
go test -race ./10-testing-and-race/...
go test -bench=. ./10-testing-and-race/...
```

## Katas

See [katas/README.md](katas/README.md).

```sh
go test -race ./10-testing-and-race/katas/...
go test -bench=. ./10-testing-and-race/katas/...
```
