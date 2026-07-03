package main

import "fmt"

// Stage 1: generates numbers onto a channel.
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// Stage 2: squares every number it receives.
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// Stage 3: keeps only even numbers.
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}

func main() {
	// Each stage runs in its own goroutine; data flows through the chain
	// concurrently rather than one stage fully finishing before the next starts.
	pipeline := filterEven(square(generate(1, 2, 3, 4, 5, 6)))

	for v := range pipeline {
		fmt.Println("result:", v)
	}
}

/*
Expected output (deterministic - a single-lane pipeline preserves order):

result: 4
result: 16
result: 36
*/
