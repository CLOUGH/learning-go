package main

// SquareAll computes the square of every number in nums, concurrently -
// one goroutine per element - and returns the results in the SAME ORDER
// as the input.
//
// This is deliberately awkward with only what lesson 03 has taught you
// (goroutines + WaitGroup): you need somewhere to safely store each
// goroutine's result. A slice indexed by position, written to from each
// goroutine, works fine here because each goroutine writes to a distinct
// index - there's no shared index being written by two goroutines at
// once. (If that sentence doesn't fully make sense yet, lesson 06 on
// races will make it click.)
//
// TODO: implement this using a WaitGroup and one goroutine per element.
func SquareAll(nums []int) []int {
	panic("TODO: implement SquareAll")
}
