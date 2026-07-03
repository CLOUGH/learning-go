# Challenges

Independent problems to test what you learned in lessons 02-09. Each
challenge is a separate package:

- `problem.go` — a stub with `TODO`s. This is what you edit.
- `problem_test.go` — tests that check correctness (and, where relevant,
  that your solution is actually concurrent / actually safe under `-race`).
- `solution/` — a complete reference solution in its own package. Don't
  look until you've had a real attempt — that's the entire point.

## Workflow

```sh
cd challenges/01-parallel-sum
go test ./...              # test your problem.go
go test -race ./...        # check for data races
go test ./solution/...      # the reference solution, once you want to compare
```

## The challenges

1. **[Parallel Sum](01-parallel-sum/)** — split a large slice across N
   goroutines and sum it concurrently. Tests goroutines + WaitGroup +
   safely combining partial results.
2. **[Worker Pool](02-worker-pool/)** — process a stream of jobs with a
   fixed number of workers, respecting `context` cancellation. Tests
   channels, WaitGroup, context together.
3. **[Pipeline](03-pipeline/)** — a 3-stage channel pipeline (generate →
   transform → filter). Tests channel composition and directionality.
4. **[Concurrent Cache](04-concurrent-cache/)** — a `Get`/`Set` cache safe
   for concurrent use, verified with `-race` under heavy concurrent
   load. Tests `sync.RWMutex`.

Work through them in order — each leans on the previous lesson's concepts.
