package solutions

import (
	"sort"
	"sync"
	"testing"
	"time"
)

func TestParallelForEach(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	var mu sync.Mutex
	var seen []int

	ParallelForEach(items, func(n int) {
		mu.Lock()
		seen = append(seen, n*n)
		mu.Unlock()
	})

	sort.Ints(seen)
	want := []int{1, 4, 9, 16, 25}
	if len(seen) != len(want) {
		t.Fatalf("got %v, want %v", seen, want)
	}
	for i := range want {
		if seen[i] != want[i] {
			t.Errorf("got %v, want %v", seen, want)
			break
		}
	}
}

func TestParallelForEachEmpty(t *testing.T) {
	called := false
	ParallelForEach(nil, func(int) { called = true })
	if called {
		t.Fatal("f should not be called for an empty slice")
	}
}

func TestRunDoneClosesWhenFinished(t *testing.T) {
	var ran bool
	var mu sync.Mutex

	done := RunDone(func() {
		time.Sleep(20 * time.Millisecond)
		mu.Lock()
		ran = true
		mu.Unlock()
	})

	select {
	case <-done:
		mu.Lock()
		defer mu.Unlock()
		if !ran {
			t.Fatal("done was closed before f finished running")
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for done to close")
	}
}

func TestRunDoneDoesNotBlock(t *testing.T) {
	start := time.Now()
	RunDone(func() { time.Sleep(200 * time.Millisecond) })
	if elapsed := time.Since(start); elapsed > 50*time.Millisecond {
		t.Fatalf("RunDone blocked for %v - it should return immediately", elapsed)
	}
}
