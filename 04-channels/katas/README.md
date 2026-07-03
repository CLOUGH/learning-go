# Katas — 04 channels

```sh
go test -race ./04-channels/katas/...
```

1. **`kata1_merge.go`** — `Merge(chans ...<-chan int) <-chan int`. Fan-in
   an arbitrary number of channels into one, closing the output once every
   input has closed.
2. **`kata2_take.go`** — `Take(in <-chan int, n int) <-chan int`. Read
   only the first `n` values from `in`, then close the output — practice
   with directional channels and closing at the right moment.
