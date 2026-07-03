package pipeline

import (
	"reflect"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	got := Collect(Generate(1, 2, 3))
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Generate: got %v, want %v", got, want)
	}
}

func TestTransform(t *testing.T) {
	got := Collect(Transform(Generate(1, 2, 3), func(n int) int { return n * 10 }))
	want := []int{10, 20, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Transform: got %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	got := Collect(Filter(Generate(1, 2, 3, 4, 5, 6), func(n int) bool { return n%2 == 0 }))
	want := []int{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Filter: got %v, want %v", got, want)
	}
}

func TestFullPipeline(t *testing.T) {
	// generate 1..10, double, keep only values > 10
	source := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	doubled := Transform(source, func(n int) int { return n * 2 })
	filtered := Filter(doubled, func(n int) bool { return n > 10 })

	got := Collect(filtered)
	want := []int{12, 14, 16, 18, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("pipeline: got %v, want %v", got, want)
	}
}

func TestChannelsCloseWithinDeadline(t *testing.T) {
	// Guards against a stage that forgets to close its output channel -
	// Collect would hang forever on a channel that's never closed.
	done := make(chan struct{})
	go func() {
		Collect(Filter(Transform(Generate(1, 2, 3), func(n int) int { return n }), func(int) bool { return true }))
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("pipeline did not finish - a stage is probably not closing its output channel")
	}
}
