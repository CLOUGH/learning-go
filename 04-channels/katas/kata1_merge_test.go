package katas

import (
	"sort"
	"testing"
	"time"
)

func chanOf(vals ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range vals {
			ch <- v
		}
	}()
	return ch
}

func TestMerge(t *testing.T) {
	merged := Merge(chanOf(1, 2), chanOf(3, 4), chanOf(5))

	var got []int
	deadline := time.After(2 * time.Second)
loop:
	for {
		select {
		case v, ok := <-merged:
			if !ok {
				break loop
			}
			got = append(got, v)
		case <-deadline:
			t.Fatal("timed out - Merge never closed its output channel")
		}
	}

	sort.Ints(got)
	want := []int{1, 2, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestMergeNoInputs(t *testing.T) {
	merged := Merge()
	select {
	case _, ok := <-merged:
		if ok {
			t.Fatal("expected the merged channel to be closed with no inputs")
		}
	case <-time.After(time.Second):
		t.Fatal("timed out - Merge with zero inputs should close immediately")
	}
}
