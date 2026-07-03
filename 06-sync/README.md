# 06 — The `sync` package

Channels aren't the only tool for concurrency in Go. When multiple
goroutines need to share a piece of state (a map, a counter, a cache), a
lock from the `sync` package is often the simpler, more direct answer than
building a channel-based protocol around it. Rule of thumb: **use channels
to pass ownership of data or signal events between goroutines; use a mutex
to protect a shared piece of state that multiple goroutines need to touch
in place.**

## The race this lesson starts with

```go
counter := 0
for i := 0; i < 1000; i++ {
    go func() { counter++ }()
}
```

`counter++` looks atomic but isn't — it's read, increment, write. Two
goroutines can both read the same value before either writes it back,
and an increment gets lost. This is a **data race**: unsynchronized
concurrent access to the same memory where at least one access is a write.
`main.go` reproduces this, then fixes it two ways.

## `sync.Mutex`

```go
var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()
```

Only one goroutine can hold the lock at a time; others calling `Lock()`
block until it's released. Always pair `Lock()` with a `defer Unlock()`
right after it, so the unlock happens even if the function returns early
or panics.

## `sync.RWMutex`

Same idea, but distinguishes readers from writers: any number of readers
can hold `RLock()` simultaneously, but `Lock()` (for writing) excludes
everyone. Use this when reads vastly outnumber writes and you want readers
to not block each other.

## `sync.Once`

Guarantees a function runs exactly once, no matter how many goroutines call
it concurrently — the standard way to do lazy, thread-safe initialization.

```go
var once sync.Once
once.Do(func() { fmt.Println("only prints once, ever") })
```

## `sync/atomic`

For simple counters, `sync/atomic` (e.g. `atomic.Int64`) can be faster and
simpler than a mutex, since it uses CPU-level atomic instructions instead
of a lock. Reach for a mutex first for anything beyond a single
counter/flag — it's easier to reason about correctly.

## Detecting races: `-race`

Never eyeball concurrent code for races — run it with the race detector,
which instruments every memory access and reports the exact conflicting
goroutines:

```sh
go test -race ./06-sync/...
```

Run the demo:

```sh
go run ./06-sync
go test -race ./06-sync/...
```

## Katas

Two more practice drills beyond the exercise above — see
[katas/README.md](katas/README.md).

```sh
go test -race ./06-sync/katas/...
```
