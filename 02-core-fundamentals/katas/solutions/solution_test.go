package solutions

import (
	"reflect"
	"testing"
)

func TestDedupeInts(t *testing.T) {
	got := Dedupe([]int{1, 2, 2, 3, 1, 4})
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Dedupe = %v, want %v", got, want)
	}
}

func TestDedupeStrings(t *testing.T) {
	got := Dedupe([]string{"a", "b", "a", "c", "b"})
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Dedupe = %v, want %v", got, want)
	}
}

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
