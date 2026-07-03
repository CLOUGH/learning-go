# Katas — 06 sync

```sh
go test -race ./06-sync/katas/...
```

1. **`kata1_oncevalue.go`** — `OnceValue[T any](f func() T) func() T`.
   Wrap any zero-argument function so it only ever actually runs once,
   no matter how many goroutines call the wrapped version concurrently —
   `sync.Once` combined with generics.
2. **`kata2_visitorset.go`** — `VisitorSet`, a concurrency-safe set of
   strings (`Add`, `Count`). The same mutex-protected-state pattern as the
   lesson's `Counter` exercise, applied to a different data shape — the
   repetition is the point.
