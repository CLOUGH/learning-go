// Package solutions holds reference solutions for the 04-channels katas.
// Try the kata yourself first - see ../README.md.
package solutions

import "sync"

// Merge - kata 1. One goroutine per input channel forwards its values to
// the shared output; a WaitGroup tracks when all of them have finished
// draining their input, which is the signal it's safe to close out.
func Merge(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	for _, in := range chans {
		go func(in <-chan int) {
			defer wg.Done()
			for v := range in {
				out <- v
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Take - kata 2. Stops receiving (and returns) after n values, or sooner
// if in closes first - either way it closes its own output exactly once.
func Take(in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			v, ok := <-in
			if !ok {
				return
			}
			out <- v
		}
	}()
	return out
}
