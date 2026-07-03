package katas

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestOnceValueRunsOnce(t *testing.T) {
	var calls int32
	get := OnceValue(func() int {
		atomic.AddInt32(&calls, 1)
		return 42
	})

	var wg sync.WaitGroup
	results := make([]int, 100)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = get()
		}(i)
	}
	wg.Wait()

	if atomic.LoadInt32(&calls) != 1 {
		t.Fatalf("f was called %d times, want exactly 1", calls)
	}
	for i, r := range results {
		if r != 42 {
			t.Fatalf("results[%d] = %d, want 42", i, r)
		}
	}
}
