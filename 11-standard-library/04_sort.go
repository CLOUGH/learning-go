package main

import (
	"fmt"
	"slices"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func demoSort() {
	fmt.Println("--- sort: convenience sorts for the built-in types ---")

	ints := []int{5, 2, 8, 1, 9}
	sort.Ints(ints)
	fmt.Println("sort.Ints:", ints)

	words := []string{"banana", "apple", "cherry"}
	sort.Strings(words)
	fmt.Println("sort.Strings:", words)

	fmt.Println()
	fmt.Println("--- sort.Slice: sorting your own types with a less-func ---")

	people := []person{
		{"Charlie", 35},
		{"Alice", 30},
		{"Bob", 25},
	}

	// sort.Slice takes a "less" function: given two indexes, report
	// whether the element at i should sort before the element at j.
	// It's not generic (predates generics), so it works via reflection +
	// an index-based callback rather than passing values directly.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("sorted by age:", people)

	// sort.SliceStable is the same idea, but guarantees equal elements
	// keep their original relative order - matters when you sort by one
	// field after already having sorted by another.
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("sorted by name:", people)

	fmt.Println()
	fmt.Println("--- the newer, generic alternative: slices.SortFunc (Go 1.21+) ---")

	// Covered in 02-core-fundamentals/08_generics.go too - slices.SortFunc
	// takes values directly (not indexes) and a comparator returning
	// negative/zero/positive, matching the cmp package's convention. For
	// new code, prefer slices.Sort/slices.SortFunc over the sort package;
	// sort.Slice is what you'll see in most pre-generics (pre-2023) Go code.
	slices.SortFunc(people, func(a, b person) int {
		return a.Age - b.Age
	})
	fmt.Println("slices.SortFunc by age:", people)
}

/*
Expected output (from demoSort, called via main.go):

--- sort: convenience sorts for the built-in types ---
sort.Ints: [1 2 5 8 9]
sort.Strings: [apple banana cherry]

--- sort.Slice: sorting your own types with a less-func ---
sorted by age: [{Bob 25} {Alice 30} {Charlie 35}]
sorted by name: [{Alice 30} {Bob 25} {Charlie 35}]

--- the newer, generic alternative: slices.SortFunc (Go 1.21+) ---
slices.SortFunc by age: [{Bob 25} {Alice 30} {Charlie 35}]
*/
