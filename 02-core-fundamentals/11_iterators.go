package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

// Countdown is a custom iterator: a function shaped like iter.Seq[int]
// (func(yield func(int) bool)). Calling yield hands the next value to
// whoever is ranging over this function; yield's return value tells the
// iterator whether to keep going (false means the consumer `break`s early).
//
// Since Go 1.23, a function with this exact shape can be used directly as
// the operand of a `for ... range` loop - this is "range-over-func". It's
// the same mechanism `for range aSlice` and `for range aMap` already used
// internally; Go 1.23 just opened it up for your own types.
func Countdown(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return // consumer broke out of the range loop early
			}
		}
	}
}

// Enumerate is a custom iter.Seq2: it yields TWO values per step (here,
// an index and a value), the same shape `maps.All` and `slices.All` use.
func Enumerate[T any](s []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

func demoIterators() {
	fmt.Println("--- range-over-func: custom iterators (Go 1.23+) ---")

	for n := range Countdown(3) {
		fmt.Println("countdown:", n)
	}

	fmt.Println("--- breaking early stops the iterator ---")
	for n := range Countdown(10) {
		if n == 7 {
			break // yield(7) returns false internally; Countdown's loop returns
		}
		fmt.Println("still counting:", n)
	}

	fmt.Println()
	fmt.Println("--- iter.Seq2: two values per step ---")
	for i, v := range Enumerate([]string{"a", "b", "c"}) {
		fmt.Println(i, v)
	}

	fmt.Println()
	fmt.Println("--- the standard library's own iterators: slices and maps ---")

	nums := []int{30, 10, 20}

	// slices.Values returns an iter.Seq[int] over the elements (no indexes).
	for v := range slices.Values(nums) {
		fmt.Println("value:", v)
	}

	// slices.All is the iter.Seq2 form: index AND value, same shape as a
	// plain `for i, v := range nums` - mostly useful when passing "the
	// sequence of this slice" around as a value instead of ranging directly.
	for i, v := range slices.All(nums) {
		fmt.Println("index", i, "->", v)
	}

	// slices.Collect turns any iter.Seq[T] back into a []T - the inverse
	// of slices.Values/slices.All. slices.Sorted does the same but sorts
	// as it collects.
	sorted := slices.Sorted(slices.Values(nums))
	fmt.Println("sorted copy:", sorted, "(original untouched:", nums, ")")

	m := map[string]int{"b": 2, "a": 1, "c": 3}

	// maps.Keys/maps.Values return iterators (NOT slices, as of Go 1.23 -
	// if you're thinking of an older tutorial that showed maps.Keys
	// returning []K, that was golang.org/x/exp/maps, a different,
	// pre-iterator package). Combine with slices.Sorted/slices.Collect to
	// get a concrete, ordered slice back out - the idiomatic way to get
	// deterministic output from a map (see 02_arrays_slices_maps.go for
	// why map iteration order is randomized in the first place).
	for _, k := range slices.Sorted(maps.Keys(m)) {
		fmt.Println(k, "=", m[k])
	}
}
