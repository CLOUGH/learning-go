# Katas — 01 fundamentals

Small, focused drills. Each is a stub in its own file with a matching
`_test.go` — fill in the function, then:

```sh
go test ./01-fundamentals/katas/...
```

1. **`kata1_palindrome.go`** — `IsPalindrome(s string) bool`. Basic
   string/loop practice.
2. **`kata2_counter.go`** — `NewCounter() func() int`. A closure that
   remembers state between calls — the same mechanism goroutines will
   later capture (for better or worse).
