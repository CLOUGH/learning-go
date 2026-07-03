# Katas — 03 goroutines

```sh
go test -race ./03-goroutines/katas/...
```

1. **`kata1_parallelforeach.go`** — `ParallelForEach(items []int, f func(int))`.
   A generalization of the lesson's `SquareAll` exercise: run `f` for every
   item concurrently and don't return until all of them have finished.
2. **`kata2_rundone.go`** — `RunDone(f func()) <-chan struct{}`. Run `f`
   in its own goroutine and return a channel that's closed the moment `f`
   finishes — a small building block you'll see again, formalized, once
   channels and `select` are introduced.
