// Package solutions holds reference solutions for the 08-patterns katas.
// Try the kata yourself first - see ../README.md.
package solutions

import "sync"

// Batch - kata 1.
func Batch(in <-chan int, size int) <-chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)
		batch := make([]int, 0, size)
		for v := range in {
			batch = append(batch, v)
			if len(batch) == size {
				out <- batch
				batch = make([]int, 0, size)
			}
		}
		if len(batch) > 0 {
			out <- batch
		}
	}()
	return out
}

// RunLimited - kata 2. A buffered channel used purely for its capacity is
// a counting semaphore: acquiring a slot is sending into it, releasing is
// receiving from it (see 08-patterns/semaphore/main.go).
func RunLimited(tasks []func(), limit int) {
	sem := make(chan struct{}, limit)
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		sem <- struct{}{} // acquire, blocks once `limit` are outstanding
		go func(task func()) {
			defer wg.Done()
			defer func() { <-sem }() // release
			task()
		}(task)
	}

	wg.Wait()
}
