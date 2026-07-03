package katas

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

func TestDedupeEmpty(t *testing.T) {
	got := Dedupe([]int{})
	if len(got) != 0 {
		t.Errorf("Dedupe(empty) = %v, want empty", got)
	}
}
