package main

import (
	"cmp"
	"fmt"
	"slices"
)

// A type parameter, `[T cmp.Ordered]`, lets Max work for any type that
// satisfies the cmp.Ordered constraint (stdlib "cmp" package: all the
// built-in numeric types plus string) - one function body, many types,
// with full compile-time type checking (no runtime type assertions
// involved, unlike `any`).
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Number is a custom constraint: an interface listing the exact set of
// underlying types allowed. `~int` (with the tilde) means "int, or any
// named type whose underlying type is int" - without the tilde it would
// only match the literal type `int`, not e.g. `type Meters int`.
type Number interface {
	~int | ~float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// A generic type: Stack[T] works for a stack of any element type, with
// no code duplication and no loss of type safety (a Stack[int]'s Pop()
// really does return an int, not `any`).
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func demoGenerics() {
	fmt.Println("--- generic functions ---")
	fmt.Println("Max(3, 7):", Max(3, 7))
	fmt.Println("Max(3.5, 1.2):", Max(3.5, 1.2))
	fmt.Println("Max(\"go\", \"rust\"):", Max("go", "rust")) // works for string too - cmp.Ordered covers it

	fmt.Println("Sum of ints:", Sum([]int{1, 2, 3, 4}))
	fmt.Println("Sum of float64s:", Sum([]float64{1.5, 2.5}))

	fmt.Println()
	fmt.Println("--- generic types ---")
	var s Stack[string]
	s.Push("a")
	s.Push("b")
	s.Push("c")
	for {
		v, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Println("popped:", v)
	}

	fmt.Println()
	fmt.Println("--- the standard library's own generic helpers (\"slices\", \"maps\", \"cmp\") ---")
	nums := []int{5, 2, 8, 1, 9}
	slices.Sort(nums) // in-place generic sort, no less-func boilerplate needed for ordered types
	fmt.Println("sorted:", nums)
	fmt.Println("contains 8:", slices.Contains(nums, 8))
	fmt.Println("max:", slices.Max(nums))

	// The exercise for this lesson asks you to write Map/Filter/Reduce
	// yourself - the classic generics exercise - so they're deliberately
	// not spoiled here.
}

/*
Expected output (from demoGenerics, called via main.go):

--- generic functions ---
Max(3, 7): 7
Max(3.5, 1.2): 3.5
Max("go", "rust"): rust
Sum of ints: 10
Sum of float64s: 4

--- generic types ---
popped: c
popped: b
popped: a

--- the standard library's own generic helpers ("slices", "maps", "cmp") ---
sorted: [1 2 5 8 9]
contains 8: true
max: 9
*/
