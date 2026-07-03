package katas

import "testing"

func intPtr(n int) *int { return &n }

func TestSumPointers(t *testing.T) {
	got := SumPointers([]*int{intPtr(1), intPtr(2), nil, intPtr(3)})
	if got != 6 {
		t.Errorf("SumPointers = %d, want 6", got)
	}
}

func TestSumPointersAllNil(t *testing.T) {
	got := SumPointers([]*int{nil, nil})
	if got != 0 {
		t.Errorf("SumPointers = %d, want 0", got)
	}
}

func TestSumPointersEmpty(t *testing.T) {
	got := SumPointers(nil)
	if got != 0 {
		t.Errorf("SumPointers = %d, want 0", got)
	}
}
