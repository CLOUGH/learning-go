package main

import (
	"fmt"
	"time"
)

func basicSelect() {
	fmt.Println("--- select waits on whichever channel is ready first ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(30 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("received:", msg)
		case msg := <-ch2:
			fmt.Println("received:", msg)
		}
	}
}

func nonBlockingWithDefault() {
	fmt.Println("--- default makes select non-blocking ---")
	ch := make(chan int)

	select {
	case v := <-ch:
		fmt.Println("got", v)
	default:
		fmt.Println("nothing ready, moving on instead of blocking")
	}
}

func timeoutPattern() {
	fmt.Println("--- time.After for timeouts ---")
	resultCh := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond) // simulate slow work
		resultCh <- "slow result"
	}()

	select {
	case res := <-resultCh:
		fmt.Println("got result:", res)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("timed out waiting for the slow goroutine")
	}
}

func cancellationPattern() {
	fmt.Println("--- done-channel cancellation ---")
	done := make(chan struct{})
	work := make(chan int)

	go func() {
		for i := 0; ; i++ {
			select {
			case <-done:
				fmt.Println("worker: told to stop")
				return
			case work <- i:
				// sent i, loop again
			}
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("main received:", <-work)
	}
	close(done)                       // tell the worker goroutine to stop
	time.Sleep(10 * time.Millisecond) // let the print above happen before main exits
}

func main() {
	basicSelect()
	fmt.Println()

	nonBlockingWithDefault()
	fmt.Println()

	timeoutPattern()
	fmt.Println()

	cancellationPattern()
}

/*
Expected output:

--- select waits on whichever channel is ready first ---
received: from ch2
received: from ch1

--- default makes select non-blocking ---
nothing ready, moving on instead of blocking

--- time.After for timeouts ---
timed out waiting for the slow goroutine

--- done-channel cancellation ---
main received: 0
main received: 1
main received: 2
worker: told to stop

"from ch2" prints before "from ch1" because ch2's goroutine sleeps 10ms
and ch1's sleeps 30ms - deterministic given those staggered delays, but
still ultimately at the scheduler's discretion, not a language guarantee.
*/
