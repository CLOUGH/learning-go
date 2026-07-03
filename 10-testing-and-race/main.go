package main

import (
	"fmt"
	"sync"
)

// RacyAdder is deliberately broken: Add is not synchronized. Running the
// demo below normally might look fine; running it with `go run -race` (or
// exercising it from a test with `go test -race`) reliably reports the
// race, because -race instruments every access rather than relying on
// timing.
type RacyAdder struct {
	total int
}

func (a *RacyAdder) Add(n int) {
	a.total += n // unsynchronized read-modify-write
}

// SafeAdder is the fix: a mutex makes Add safe to call concurrently.
type SafeAdder struct {
	mu    sync.Mutex
	total int
}

func (a *SafeAdder) Add(n int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.total += n
}

func (a *SafeAdder) Total() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.total
}

func main() {
	fmt.Println("--- run this file with -race to see RacyAdder get caught ---")
	fmt.Println("    go run -race ./10-testing-and-race")

	racy := &RacyAdder{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			racy.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println("RacyAdder total (may not be 100):", racy.total)

	safe := &SafeAdder{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			safe.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println("SafeAdder total (always 100):", safe.Total())
}
