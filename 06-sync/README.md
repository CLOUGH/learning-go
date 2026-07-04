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

`sync.Mutex` is the same concept as Java's `synchronized` keyword/
`ReentrantLock`, C's `pthread_mutex_t`, or C++'s `std::mutex` — the main
difference is Go has no `synchronized` block sugar or automatic
lock-on-scope-exit (no RAII), which is exactly why the `defer Unlock()`
idiom exists: it's Go's manual stand-in for what C++'s
`std::lock_guard`/`std::unique_lock` destructor gives you for free. JS and
PHP don't need this in ordinary code — JS because it's single-threaded,
PHP because a request's execution generally isn't sharing mutable memory
with another concurrently-running request the way threads do.

## `sync.RWMutex`

Same idea, but distinguishes readers from writers: any number of readers
can hold `RLock()` simultaneously, but `Lock()` (for writing) excludes
everyone. Use this when reads vastly outnumber writes and you want readers
to not block each other. Direct equivalent of Java's `ReadWriteLock`/
`ReentrantReadWriteLock` and C++'s `std::shared_mutex`.

## `sync.Once`

Guarantees a function runs exactly once, no matter how many goroutines call
it concurrently — the standard way to do lazy, thread-safe initialization.
Solves the same problem Java developers reach for double-checked locking
or a `static` initializer for, and that C++ solves with a function-local
`static` variable (whose initialization the language itself guarantees is
thread-safe, exactly once) — `sync.Once` is Go's explicit version of that
same guarantee.

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

## Real-world use cases

- **An in-memory cache shared across request-handling goroutines** — a
  `map[string]Value` guarded by a `sync.RWMutex` is the standard shape
  for a process-local cache (rate-limit counters, feature-flag values, a
  session store) that many concurrent HTTP requests read constantly and
  occasionally write:

  ```go
  type Cache struct {
      mu   sync.RWMutex
      data map[string]string
  }

  func (c *Cache) Get(key string) (string, bool) {
      c.mu.RLock()
      defer c.mu.RUnlock()
      v, ok := c.data[key]
      return v, ok
  }
  ```

- **Lazy, thread-safe singleton initialization** — `sync.Once` around
  setting up a database connection pool or parsing a config file exactly
  once, no matter how many goroutines call the accessor concurrently
  during startup.
- **`sync/atomic` counters for metrics** — request counters, in-flight
  request gauges, and similar high-frequency single-value updates in a
  server's hot path use `atomic.Int64` instead of a mutex, since the
  lock's overhead would be measurable at that call frequency.

## Katas

Two more practice drills beyond the exercise above — see
[katas/README.md](katas/README.md).

```sh
go test -race ./06-sync/katas/...
```
