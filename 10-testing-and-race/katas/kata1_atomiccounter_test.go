package katas

import (
	"sync"
	"testing"
)

// go test -race ./10-testing-and-race/katas/...
func TestAtomicCounterConcurrentIncrements(t *testing.T) {
	c := NewAtomicCounter()
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

	want := int64(goroutines * incrementsEach)
	if got := c.Value(); got != want {
		t.Errorf("Value() = %d, want %d", got, want)
	}
}
