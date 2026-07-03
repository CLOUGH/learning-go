# Katas — 02 core fundamentals

```sh
go test ./02-core-fundamentals/katas/...
```

1. **`kata1_dedupe.go`** — `Dedupe[T comparable](items []T) []T`. Generics
   + maps practice: remove duplicates while preserving first-seen order.
2. **`kata2_sumpointers.go`** — `SumPointers(nums []*int) int`. Pointer
   practice: sum the values behind a slice of pointers, treating any nil
   pointer as 0 instead of panicking.

## Solutions

Have a real attempt first — then check `solutions/solution.go` against
what you wrote:

```sh
go test ./02-core-fundamentals/katas/solutions/...
```
