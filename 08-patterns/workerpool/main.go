package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID  int
	Output int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs { // exits automatically when jobs is closed and drained
		time.Sleep(10 * time.Millisecond) // simulate work
		results <- Result{JobID: j.ID, Output: j.ID * j.ID}
		fmt.Printf("worker %d handled job %d\n", id, j.ID)
	}
}

func main() {
	const numJobs = 9
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs) // no more jobs coming - lets workers' range loops finish

	// Close results once all workers are done, so the range below can
	// finish too, instead of us having to know the exact count up front.
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("result for job %d: %d\n", r.JobID, r.Output)
	}
}

/*
Expected output (one possible interleaving - which worker gets which job,
and the exact print order, both vary between runs; only the *set* of
9 jobs handled and their squared results is guaranteed):

worker 3 handled job 3
worker 2 handled job 2
result for job 3: 9
worker 1 handled job 1
result for job 1: 1
result for job 2: 4
worker 3 handled job 4
result for job 4: 16
worker 1 handled job 6
result for job 6: 36
worker 2 handled job 5
result for job 5: 25
worker 2 handled job 9
result for job 9: 81
worker 3 handled job 7
result for job 7: 49
worker 1 handled job 8
result for job 8: 64
*/
