package main

import (
	"fmt"
	"sync"
	"time"
)

func broken() {
	fmt.Println("--- broken: main may exit before the goroutine runs ---")
	go fmt.Println("hello from a goroutine")
	fmt.Println("hello from main")
	// No synchronization: the goroutine above is racing main's return.
	// Try commenting out the time.Sleep in main() and running this
	// repeatedly - sometimes you'll see the goroutine's line, sometimes not.
}

func withWaitGroup() {
	fmt.Println("--- fixed: sync.WaitGroup makes main wait ---")
	var wg sync.WaitGroup

	messages := []string{"one", "two", "three"}
	for _, msg := range messages {
		wg.Add(1)
		go func(m string) {
			defer wg.Done() // always runs, even if this goroutine panics
			fmt.Println("worker got:", m)
		}(msg) // passed as an argument, not captured - see 09-pitfalls for why this matters
	}

	wg.Wait() // blocks here until all 3 have called wg.Done()
	fmt.Println("all workers done")
}

func withWaitGroupGo() {
	fmt.Println("--- Go 1.25+: wg.Go(f) replaces Add(1)+go func(){defer Done(); f()}() ---")
	var wg sync.WaitGroup

	messages := []string{"one", "two", "three"}
	for _, msg := range messages {
		// wg.Go only takes a func(), so msg must be captured by the closure
		// rather than passed in as a parameter the way withWaitGroup does
		// above. That's safe here only because this module targets Go
		// 1.22+: each loop iteration gets its own `msg` (see 09-pitfalls for
		// the pre-1.22 version of this gotcha, which this exact pattern used
		// to trigger).
		wg.Go(func() {
			fmt.Println("worker got:", msg)
		})
	}

	wg.Wait()
	fmt.Println("all workers done")
	// wg.Go(f) does exactly what withWaitGroup above does by hand:
	// wg.Add(1), then `go func() { defer wg.Done(); f() }()`. It's pure
	// convenience - same semantics, less boilerplate, one fewer place to
	// forget a wg.Done(). It doesn't change anything else about WaitGroup.
}

func interleaving() {
	fmt.Println("--- goroutines interleave; order is not guaranteed ---")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(5-id) * time.Millisecond)
			fmt.Println("goroutine", id, "finished")
		}(i)
	}
	wg.Wait()
	// Run this a few times: the print order changes because it depends on
	// scheduling, not on the order goroutines were started in.
}

func main() {
	broken()
	time.Sleep(50 * time.Millisecond) // crude hack so you can SEE the race; never do this for real sync
	fmt.Println()

	withWaitGroup()
	fmt.Println()

	withWaitGroupGo()
	fmt.Println()

	interleaving()
}
