# Katas — 07 context

```sh
go test -race ./07-context/katas/...
```

1. **`kata1_delay.go`** — `Delay(ctx context.Context, d time.Duration)
   error`. Sleep for `d`, unless `ctx` is cancelled or times out first —
   returns `ctx.Err()` in that case, `nil` otherwise.
2. **`kata2_runwithtimeout.go`** — `RunWithTimeout(f func(), timeout
   time.Duration) error`. Run `f` in the background; return `nil` if it
   finishes within `timeout`, otherwise return `context.DeadlineExceeded`
   without waiting any further for it.

## Solutions

Have a real attempt first — then check `solutions/solution.go` against
what you wrote:

```sh
go test -race ./07-context/katas/solutions/...
```
