// Package solution is the reference solution for challenge 01 (parallel sum).
package solution

import "sync"

// ParallelSum splits nums into numWorkers contiguous chunks, sums each
// chunk in its own goroutine, and combines the per-goroutine partial sums.
func ParallelSum(nums []int, numWorkers int) int {
	if len(nums) == 0 {
		return 0
	}
	if numWorkers <= 0 {
		numWorkers = 1
	}
	if numWorkers > len(nums) {
		numWorkers = len(nums) // no point in more workers than elements
	}

	chunkSize := (len(nums) + numWorkers - 1) / numWorkers // ceil division
	partials := make([]int, numWorkers)

	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		start := w * chunkSize
		if start >= len(nums) {
			break
		}
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}

		wg.Add(1)
		go func(workerIdx, start, end int) {
			defer wg.Done()
			sum := 0
			for _, n := range nums[start:end] {
				sum += n
			}
			partials[workerIdx] = sum // each goroutine writes a distinct index - no race
		}(w, start, end)
	}
	wg.Wait()

	total := 0
	for _, p := range partials {
		total += p
	}
	return total
}
