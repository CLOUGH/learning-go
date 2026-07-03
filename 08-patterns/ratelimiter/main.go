package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// A Ticker sends a value on its channel every `d`. Gating each request
	// on a tick caps throughput to 1 request per tick, regardless of how
	// fast `requests` could otherwise be drained.
	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop() // always stop a Ticker, or its goroutine leaks forever

	start := time.Now()
	for req := range requests {
		<-limiter.C // block until the next tick
		fmt.Printf("request %d handled at %v\n", req, time.Since(start).Round(time.Millisecond))
	}
}

/*
Expected output (the exact millisecond values will drift a little run to
run - real wall-clock timing, not simulated - but each should land
roughly 200ms after the previous one):

request 1 handled at 200ms
request 2 handled at 401ms
request 3 handled at 601ms
request 4 handled at 800ms
request 5 handled at 1s
*/
