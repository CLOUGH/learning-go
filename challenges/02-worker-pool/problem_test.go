package workerpool

import (
	"context"
	"sync"
	"testing"
	"time"
)

func collectWithTimeout(t *testing.T, results <-chan Result, timeout time.Duration) []Result {
	t.Helper()
	var got []Result
	deadline := time.After(timeout)
	for {
		select {
		case r, ok := <-results:
			if !ok {
				return got
			}
			got = append(got, r)
		case <-deadline:
			t.Fatal("timed out waiting for results channel to close - possible deadlock or goroutine leak")
			return nil
		}
	}
}

func TestRunProcessesAllJobs(t *testing.T) {
	const numJobs = 20
	jobs := make(chan Job, numJobs)
	for i := 0; i < numJobs; i++ {
		jobs <- Job{ID: i, Value: i}
	}
	close(jobs)

	results := Run(context.Background(), jobs, 4, func(j Job) Result {
		return Result{JobID: j.ID, Value: j.Value * 2}
	})

	got := collectWithTimeout(t, results, 2*time.Second)
	if len(got) != numJobs {
		t.Fatalf("got %d results, want %d", len(got), numJobs)
	}

	seen := make(map[int]int)
	for _, r := range got {
		seen[r.JobID] = r.Value
	}
	for i := 0; i < numJobs; i++ {
		if seen[i] != i*2 {
			t.Errorf("job %d: got value %d, want %d", i, seen[i], i*2)
		}
	}
}

func TestRunRespectsWorkerLimit(t *testing.T) {
	const numJobs = 12
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	for i := 0; i < numJobs; i++ {
		jobs <- Job{ID: i, Value: i}
	}
	close(jobs)

	var mu sync.Mutex
	current, max := 0, 0

	results := Run(context.Background(), jobs, numWorkers, func(j Job) Result {
		mu.Lock()
		current++
		if current > max {
			max = current
		}
		mu.Unlock()

		time.Sleep(20 * time.Millisecond)

		mu.Lock()
		current--
		mu.Unlock()

		return Result{JobID: j.ID, Value: j.Value}
	})

	got := collectWithTimeout(t, results, 3*time.Second)
	if len(got) != numJobs {
		t.Fatalf("got %d results, want %d", len(got), numJobs)
	}

	mu.Lock()
	defer mu.Unlock()
	if max > numWorkers {
		t.Errorf("observed %d concurrent process() calls, want at most %d", max, numWorkers)
	}
	if max < 2 {
		t.Errorf("observed max concurrency of %d - workers don't seem to be running in parallel at all", max)
	}
}

func TestRunStopsOnCancellation(t *testing.T) {
	jobs := make(chan Job) // unbuffered, never closed - workers would block forever without cancellation
	ctx, cancel := context.WithCancel(context.Background())

	results := Run(ctx, jobs, 2, func(j Job) Result {
		return Result{JobID: j.ID, Value: j.Value}
	})

	cancel() // cancel immediately - no jobs will ever be sent

	// Run must close `results` promptly once cancelled, even though `jobs`
	// is empty and never closed. If Run doesn't respect ctx, this hangs
	// until the test's own timeout fires.
	got := collectWithTimeout(t, results, 2*time.Second)
	if len(got) != 0 {
		t.Fatalf("got %d results after immediate cancellation, want 0", len(got))
	}
}
