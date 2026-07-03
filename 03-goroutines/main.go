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

	interleaving()
}
