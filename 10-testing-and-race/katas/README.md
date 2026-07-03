# Katas — 10 testing and race

```sh
go test -race ./10-testing-and-race/katas/...
go test -bench=. ./10-testing-and-race/katas/...
```

1. **`kata1_atomiccounter.go`** — `AtomicCounter` using `sync/atomic`
   instead of a mutex. Same `Inc`/`Value` shape you've built before with
   `sync.Mutex` (lesson 06's `Counter` exercise), but using the
   lighter-weight tool mentioned in that lesson's README for simple
   counters.
2. **`kata2_benchmark_test.go`** — complete two benchmarks (one for
   `AtomicCounter`, one for a mutex-based counter given to you) using
   `b.RunParallel`, then run both with `-bench=.` and compare the
   `ns/op` — a concrete, measured answer to "is atomic actually faster
   here?" instead of just taking it on faith.

## Solutions

Have a real attempt first — then check `solutions/solution.go` against
what you wrote:

```sh
go test -race ./10-testing-and-race/katas/solutions/...
go test -bench=. -benchmem ./10-testing-and-race/katas/solutions/...
```
