# Learning Go — with a focus on goroutines

A self-contained, hands-on curriculum for learning Go, built directly into this
repo so every lesson is real, runnable code. The path moves through the core
language quickly, then goes deep on goroutines and concurrency, since that's
the part of Go that's genuinely different from most other languages.

## How to use this

Each numbered folder is one lesson:

- `README.md` — the concept, explained, with things to notice.
- `main.go` — a runnable demo. Run it with `go run ./NN-topic`.
- Some lessons also have `exercise.go` (a stub with `TODO`s for you to fill
  in) and `exercise_test.go` (tests to check your work). Run tests with:

  ```sh
  go test ./NN-topic/...
  go test -race ./NN-topic/...   # catch data races
  ```

- Every lesson also has a `katas/` folder: 1-2 smaller, more focused
  practice drills (stub + test, same idea as `exercise.go` but quicker to
  attempt and meant for repetition rather than a one-time exercise). See
  each lesson's `katas/README.md`, or run all of them at once:

  ```sh
  go test -race ./.../katas/...   # per lesson
  ```

Work through the lessons in order — later ones assume earlier ones. Don't
just read the code; run it, then break it on purpose (comment things out,
remove a `wg.Done()`, remove a `close()`) and see what happens. That's where
the real understanding of goroutines comes from.

## Roadmap

| # | Topic | What you'll learn |
|---|-------|--------------------|
| [00](00-setup/README.md) | Setup | Toolchain, modules, `go run`/`go test`/`go vet` |
| [01](01-fundamentals/README.md) | Fundamentals (quick tour) | Types, structs, interfaces, errors, closures |
| [02](02-core-fundamentals/README.md) | Core fundamentals (deep dive) | Types, arrays/slices/maps internals, structs & embedding, interfaces, generics, errors, defer/panic/recover, packages |
| [03](03-goroutines/README.md) | Goroutines | `go` keyword, scheduler basics, `WaitGroup` |
| [04](04-channels/README.md) | Channels | Buffered/unbuffered, direction, `close`, `range` |
| [05](05-select/README.md) | Select | Multiplexing, timeouts, `default`, cancellation |
| [06](06-sync/README.md) | sync package | `Mutex`, `RWMutex`, `Once`, races vs. channels |
| [07](07-context/README.md) | Context | Cancellation, deadlines, request-scoped values |
| [08](08-patterns/README.md) | Concurrency patterns | Worker pools, pipelines, fan-in/out, rate limiting, semaphores |
| [09](09-pitfalls/README.md) | Pitfalls | Leaks, deadlocks, loop-variable capture, closed-channel panics |
| [10](10-testing-and-race/README.md) | Testing concurrent code | `-race`, `t.Parallel`, avoiding `time.Sleep` as sync |
| [11](11-standard-library/README.md) | Standard library tour | `strings`/`strconv`, `os`/`io`/`bufio`, `time`, `sort`, `encoding/json`, `net/http`, `log`/`slog`, `flag` |

## Challenges

Once you've been through the lessons (and their katas), [challenges/](challenges/README.md)
has bigger, independent coding problems (parallel sum, worker pool, pipeline,
concurrent cache) with tests to check yourself and a reference solution to
compare against afterward.

## Checklist

- [ ] 00 — Setup
- [ ] 01 — Fundamentals (quick tour)
- [ ] 02 — Core fundamentals (deep dive)
- [ ] 03 — Goroutines
- [ ] 04 — Channels
- [ ] 05 — Select
- [ ] 06 — sync package
- [ ] 07 — Context
- [ ] 08 — Concurrency patterns
- [ ] 09 — Pitfalls
- [ ] 10 — Testing concurrent code
- [ ] 11 — Standard library tour
- [ ] Challenge 01 — Parallel Sum
- [ ] Challenge 02 — Worker Pool
- [ ] Challenge 03 — Pipeline
- [ ] Challenge 04 — Concurrent Cache
