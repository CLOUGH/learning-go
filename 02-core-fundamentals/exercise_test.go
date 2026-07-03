package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(n int) string {
		return string(rune('a' + n - 1))
	})
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Map = %v, want %v", got, want)
	}
}

func TestFilter(t *testing.T) {
	got := Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	want := []int{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Filter = %v, want %v", got, want)
	}
}

func TestReduce(t *testing.T) {
	sum := Reduce([]int{1, 2, 3, 4}, 0, func(acc, n int) int { return acc + n })
	if sum != 10 {
		t.Errorf("Reduce sum = %d, want 10", sum)
	}

	concat := Reduce([]string{"a", "b", "c"}, "", func(acc, s string) string { return acc + s })
	if concat != "abc" {
		t.Errorf("Reduce concat = %q, want %q", concat, "abc")
	}
}

func TestFindItemNotFound(t *testing.T) {
	items := map[int]string{1: "sword", 2: "shield"}
	_, err := FindItem(items, 99)
	if err == nil {
		t.Fatal("expected an error for a missing id")
	}
	var nf *NotFoundError
	if !errors.As(err, &nf) {
		t.Fatalf("expected a *NotFoundError, got %T: %v", err, err)
	}
}

func TestFindItemPermissionDenied(t *testing.T) {
	items := map[int]string{1: "sword"}
	_, err := FindItem(items, -1)
	if !errors.Is(err, ErrPermissionDenied) {
		t.Fatalf("expected errors.Is(err, ErrPermissionDenied), got %v", err)
	}
}

func TestFindItemSuccess(t *testing.T) {
	items := map[int]string{1: "sword"}
	name, err := FindItem(items, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if name != "sword" {
		t.Errorf("name = %q, want %q", name, "sword")
	}
}
