package main

import (
	"fmt"
	"sync"
	"testing"
)

// go test -race ./10-testing-and-race/...  should flag DATA RACE here.
func TestRacyAdder(t *testing.T) {
	a := &RacyAdder{}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.Add(1)
		}()
	}
	wg.Wait()
	t.Logf("RacyAdder total: %d (expected 50, may differ - and -race should flag the underlying race regardless)", a.total)
}

// The same test shape, but race-free - this is what "fixed" looks like.
func TestSafeAdder(t *testing.T) {
	a := &SafeAdder{}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.Add(1)
		}()
	}
	wg.Wait()
	if got := a.Total(); got != 50 {
		t.Errorf("Total() = %d, want 50", got)
	}
}

func TestParallelSubtests(t *testing.T) {
	cases := []int{1, 2, 3, 4}
	for _, c := range cases {
		c := c // needed for correctness on Go <1.22; harmless on 1.22+
		t.Run(fmt.Sprintf("case-%d", c), func(t *testing.T) {
			t.Parallel() // safe here because each subtest only touches its own local `c`
			if c <= 0 {
				t.Fatalf("expected positive, got %d", c)
			}
		})
	}
}

func BenchmarkSafeAdderInc(b *testing.B) {
	a := &SafeAdder{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			a.Add(1)
		}
	})
}

// BenchmarkSafeAdderAddSequential benchmarks from a single goroutine (no
// concurrency, no lock contention) - this is the case for testing.B.Loop
// (Go 1.24+), which replaced the classic `for i := 0; i < b.N; i++` idiom.
// b.RunParallel/pb.Next() above is still the right tool once you actually
// want multiple goroutines hammering the same benchmarked code at once.
func BenchmarkSafeAdderAddSequential(b *testing.B) {
	a := &SafeAdder{}
	for b.Loop() {
		a.Add(1)
	}
}
