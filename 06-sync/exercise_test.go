package main

import (
	"sync"
	"testing"
)

// Run with: go test -race ./06-sync/...
// If Inc/Value aren't properly synchronized, this test will either report
// the wrong final count or -race will flag a data race directly.
func TestCounterConcurrentIncrements(t *testing.T) {
	c := NewCounter()
	var wg sync.WaitGroup

	const goroutines = 100
	const incrementsEach = 100

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsEach; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()

	want := goroutines * incrementsEach
	if got := c.Value(); got != want {
		t.Errorf("Value() = %d, want %d", got, want)
	}
}
