package katas

import (
	"sort"
	"sync"
	"testing"
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

	sort.Ints(seen) // order isn't guaranteed - sort before comparing
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
