package main

import (
	"fmt"
	"sync"
)

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

// square is the "slow" stage we want to parallelize (fan-out): several
// goroutines all call square(in), each reading from the SAME input channel,
// so the work is naturally spread across them - whichever goroutine is
// free next grabs the next value.
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

// merge fans multiple input channels into one output channel (fan-in).
func merge(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(in <-chan int) {
			defer wg.Done()
			for v := range in {
				out <- v
			}
		}(in)
	}

	// Close out only once every input goroutine above has finished,
	// otherwise we'd close it while some goroutine might still send.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	source := generate(1, 2, 3, 4, 5, 6, 7, 8)

	// Fan-out: 3 goroutines all pulling from the same `source` channel.
	const workers = 3
	squared := make([]<-chan int, workers)
	for i := 0; i < workers; i++ {
		squared[i] = square(source)
	}

	// Fan-in: merge the 3 result channels back into one.
	for v := range merge(squared...) {
		fmt.Println("squared:", v)
	}
	// Note: output order is not guaranteed - values race across 3 workers.
}

/*
Expected output (one possible order - only the *set* {1,4,9,16,25,36,49,64}
is guaranteed, not this ordering):

squared: 1
squared: 9
squared: 4
squared: 25
squared: 16
squared: 36
squared: 64
squared: 49
*/
