package katas

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunLimitedRunsEverything(t *testing.T) {
	var count int32
	tasks := make([]func(), 20)
	for i := range tasks {
		tasks[i] = func() { atomic.AddInt32(&count, 1) }
	}

	RunLimited(tasks, 4)

	if count != 20 {
		t.Errorf("count = %d, want 20", count)
	}
}

func TestRunLimitedRespectsLimit(t *testing.T) {
	var mu sync.Mutex
	current, max := 0, 0

	const limit = 3
	tasks := make([]func(), 15)
	for i := range tasks {
		tasks[i] = func() {
			mu.Lock()
			current++
			if current > max {
				max = current
			}
			mu.Unlock()

			time.Sleep(15 * time.Millisecond)

			mu.Lock()
			current--
			mu.Unlock()
		}
	}

	RunLimited(tasks, limit)

	mu.Lock()
	defer mu.Unlock()
	if max > limit {
		t.Errorf("observed %d concurrent tasks, want at most %d", max, limit)
	}
	if max < 2 {
		t.Errorf("observed max concurrency of %d - tasks don't seem to run in parallel at all", max)
	}
}
