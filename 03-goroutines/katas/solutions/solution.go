// Package solutions holds reference solutions for the 03-goroutines
// katas. Try the kata yourself first - see ../README.md.
package solutions

import "sync"

// ParallelForEach - kata 1. Using wg.Go (Go 1.25+) instead of the manual
// wg.Add(1) + `go func(){ defer wg.Done(); ... }()` shape - see
// 03-goroutines/README.md's "Go 1.25+: wg.Go(f)" section. Either form is
// correct; this is just the more modern spelling.
func ParallelForEach(items []int, f func(int)) {
	var wg sync.WaitGroup
	for _, item := range items {
		// No `item := item` needed: this module targets Go 1.22+, where
		// each loop iteration already gets its own `item` (see 09-pitfalls).
		wg.Go(func() {
			f(item)
		})
	}
	wg.Wait()
}

// RunDone - kata 2. No WaitGroup needed here: a single goroutine, and the
// channel itself IS the "finished" signal - closing it after f() returns
// is enough for any number of receivers to observe completion.
func RunDone(f func()) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		f()
	}()
	return done
}
