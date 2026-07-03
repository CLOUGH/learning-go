# Katas — 08 concurrency patterns

```sh
go test -race ./08-patterns/katas/...
```

1. **`kata1_batch.go`** — `Batch(in <-chan int, size int) <-chan []int`.
   Group values from `in` into slices of up to `size`, emitting a batch
   as soon as it's full, plus one final (possibly smaller) batch when
   `in` closes.
2. **`kata2_runlimited.go`** — `RunLimited(tasks []func(), limit int)`.
   Run every task in `tasks`, with at most `limit` running concurrently
   at any moment — the worker-pool/semaphore idea from this lesson,
   applied to plain functions instead of typed jobs.

## Solutions

Have a real attempt first — then check `solutions/solution.go` against
what you wrote:

```sh
go test -race ./08-patterns/katas/solutions/...
```
