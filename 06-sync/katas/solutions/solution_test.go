package solutions

import (
	"fmt"
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

func TestVisitorSetBasic(t *testing.T) {
	v := NewVisitorSet()
	v.Add("alice")
	v.Add("bob")
	v.Add("alice")

	if got := v.Count(); got != 2 {
		t.Errorf("Count() = %d, want 2", got)
	}
}

func TestVisitorSetConcurrent(t *testing.T) {
	v := NewVisitorSet()
	var wg sync.WaitGroup

	const goroutines = 50
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			v.Add(fmt.Sprintf("visitor-%d", id%10))
		}(i)
	}
	wg.Wait()

	if got := v.Count(); got != 10 {
		t.Errorf("Count() = %d, want 10", got)
	}
}
