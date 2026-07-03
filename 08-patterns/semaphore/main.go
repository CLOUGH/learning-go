package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// A buffered channel used purely for its capacity: sending into it
	// "acquires" one of maxConcurrent slots, receiving "releases" one.
	// Once the buffer is full, further sends block until a slot frees up.
	const maxConcurrent = 2
	sem := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem <- struct{}{}        // acquire a slot (blocks if all are taken)
			defer func() { <-sem }() // release the slot when done

			fmt.Printf("task %d: running (at most %d run at once)\n", id, maxConcurrent)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("task %d: done\n", id)
		}(i)
	}
	wg.Wait()
}
