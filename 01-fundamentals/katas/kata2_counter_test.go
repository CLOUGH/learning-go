package katas

import "testing"

func TestNewCounter(t *testing.T) {
	next := NewCounter()
	for want := 1; want <= 3; want++ {
		if got := next(); got != want {
			t.Fatalf("next() = %d, want %d", got, want)
		}
	}
}

func TestNewCounterIndependence(t *testing.T) {
	a := NewCounter()
	b := NewCounter()

	a()
	a()
	if got := b(); got != 1 {
		t.Fatalf("a's calls leaked into b: b() = %d, want 1", got)
	}
}
