# Katas — 05 select

```sh
go test -race ./05-select/katas/...
```

1. **`kata1_first.go`** — `First(a, b <-chan int) int`. Return whichever
   of two channels produces a value first.
2. **`kata2_withtimeout.go`** — `WithTimeout(ch <-chan int, d
   time.Duration) (int, bool)`. Receive from `ch`, but give up after `d`
   — the timeout pattern from the lesson, as a reusable function.
