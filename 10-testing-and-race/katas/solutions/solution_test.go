package solutions

import (
	"sync"
	"testing"
)

// go test -race ./10-testing-and-race/katas/solutions/...
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

// kata 2: complete both benchmarks using b.RunParallel.
//
//	go test -bench=. -benchmem ./10-testing-and-race/katas/solutions/...

func BenchmarkAtomicCounterInc(b *testing.B) {
	c := NewAtomicCounter()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc()
		}
	})
}

func BenchmarkMutexCounterInc(b *testing.B) {
	c := NewMutexCounter()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc()
		}
	})
}
