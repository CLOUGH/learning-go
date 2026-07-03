package main

import (
	"fmt"
	"time"
)

func unbufferedDemo() {
	fmt.Println("--- unbuffered channel: send blocks until received ---")
	ch := make(chan string)

	go func() {
		fmt.Println("goroutine: about to send")
		ch <- "hello"
		fmt.Println("goroutine: send completed") // only prints after main receives
	}()

	time.Sleep(20 * time.Millisecond) // just so the prints above are visible in order
	msg := <-ch
	fmt.Println("main: received", msg)
}

func bufferedDemo() {
	fmt.Println("--- buffered channel: send only blocks when full ---")
	ch := make(chan int, 2)

	ch <- 1 // doesn't block, buffer has room
	ch <- 2 // doesn't block, buffer now full
	fmt.Println("sent two values without a receiver ready")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func producerConsumer() {
	fmt.Println("--- producer/consumer with close() + range ---")
	nums := make(chan int)

	go func() {
		defer close(nums) // signal "no more values" once done producing
		for i := 1; i <= 5; i++ {
			nums <- i
		}
	}()

	// range over a channel: reads until the channel is closed and drained
	for n := range nums {
		fmt.Println("consumed:", n)
	}
	fmt.Println("channel closed, loop exited cleanly")
}

func closedChannelBehavior() {
	fmt.Println("--- reading a closed channel: zero value + ok=false ---")
	ch := make(chan int, 1)
	ch <- 42
	close(ch)

	v, ok := <-ch
	fmt.Println("first read:", v, ok) // 42 true (buffered value still there)

	v, ok = <-ch
	fmt.Println("second read:", v, ok) // 0 false (channel closed and empty)
}

func main() {
	unbufferedDemo()
	fmt.Println()

	bufferedDemo()
	fmt.Println()

	producerConsumer()
	fmt.Println()

	closedChannelBehavior()
}

/*
Expected output:

--- unbuffered channel: send blocks until received ---
goroutine: about to send
main: received hello

--- buffered channel: send only blocks when full ---
sent two values without a receiver ready
1
2

--- producer/consumer with close() + range ---
consumed: 1
consumed: 2
consumed: 3
consumed: 4
consumed: 5
channel closed, loop exited cleanly

--- reading a closed channel: zero value + ok=false ---
first read: 42 true
second read: 0 false

One line's position isn't guaranteed: unbufferedDemo's goroutine prints
"goroutine: send completed" AFTER the channel rendezvous with main, as an
unsynchronized continuation racing everything that runs after it - it
usually appears right after "main: received hello", but can just as
easily show up interleaved into the next section's output instead
(it did, in the run this comment was captured from).
*/
