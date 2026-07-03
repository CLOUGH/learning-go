# Katas — 09 pitfalls

```sh
go test -race ./09-pitfalls/katas/...
```

1. **`kata1_publish.go`** — `Publish(done <-chan struct{}, values []int)
   <-chan int`. Send every value on the returned channel, in order - but
   stop immediately (without leaking a blocked goroutine) if `done` is
   closed before all values have been sent.
2. **`kata2_safelist.go`** — `SafeList`, a mutex-protected slice (`Add`,
   `Items`). The same "always guard shared state" lesson as the
   `Counter`/`VisitorSet` katas, on a slice this time - and a reminder
   that `Items()` must return a copy, not the live internal slice, or a
   caller could race with future `Add` calls just by holding the result.
