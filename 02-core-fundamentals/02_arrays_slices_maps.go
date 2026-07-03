package main

import "fmt"

func demoArraysSlicesMaps() {
	fmt.Println("--- arrays (value semantics) ---")

	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // COPIES all 3 elements - arrays are values, like structs
	arr2[0] = 99
	fmt.Println("arr1:", arr1, "arr2:", arr2) // arr1 unaffected

	fmt.Println()
	fmt.Println("--- slices: a view over a backing array, not a value ---")

	// make(type, len, cap): len is the visible length, cap is how much
	// room exists in the backing array before a new one must be allocated.
	s1 := make([]int, 3, 5)
	fmt.Printf("s1=%v len=%d cap=%d\n", s1, len(s1), cap(s1))

	// Slicing shares the SAME backing array - mutating through one slice
	// is visible through the other, until an append forces a reallocation.
	original := []int{1, 2, 3, 4, 5}
	view := original[1:3] // {2, 3}, but backed by the same array as `original`
	view[0] = 999
	fmt.Println("mutating `view` also changed `original`:", original) // [1 999 3 4 5]

	// append: if there's spare capacity, it writes into the same backing
	// array (and is visible to other slices sharing that array!). If
	// there ISN'T spare capacity, Go allocates a new, bigger backing array
	// instead - and from that point on the two slices are no longer related.
	a := make([]int, 2, 4) // len 2, cap 4 - 2 spare slots
	a[0], a[1] = 1, 2
	b := a[:cap(a)]    // reslice to the full capacity so b can "see" the spare slots too
	a = append(a, 100) // fits in spare capacity -> writes into the shared array at index 2
	fmt.Println("b reveals a's append because they share a backing array:", b)
	// b's own length never changed (append only grows the header it
	// RETURNS), but index 2 of the shared array is now 100, and b's
	// length already reached that far - so b[2] shows it.

	// copy() is the safe way to duplicate data when you don't want aliasing.
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = -1
	fmt.Println("src untouched after copy+mutate dst:", src, dst)

	fmt.Println()
	fmt.Println("--- maps ---")

	// The zero value of a map is nil. Reading from a nil map is fine (you
	// just get zero values); WRITING to a nil map panics. Always
	// make(map[K]V) or use a map literal before writing.
	var nilMap map[string]int
	fmt.Println("read from nil map is safe:", nilMap["missing"]) // 0, no panic

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	// The comma-ok idiom distinguishes "present with zero value" from "absent".
	if v, ok := m["a"]; ok {
		fmt.Println("a is present:", v)
	}
	if _, ok := m["z"]; !ok {
		fmt.Println("z is absent")
	}

	delete(m, "a") // deleting an absent key is also a no-op, never panics
	fmt.Println("after delete:", m)

	// Map iteration order is deliberately randomized by the runtime, on
	// every run, specifically so nobody writes code that accidentally
	// depends on it. Sort the keys yourself if you need stable output.
	fmt.Println("map length:", len(m))
}
