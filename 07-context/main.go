package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int, work <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: stopping (%v)\n", id, ctx.Err())
			return
		case w, ok := <-work:
			if !ok {
				return
			}
			fmt.Printf("worker %d: processing %d\n", id, w)
		}
	}
}

func timeoutDemo() {
	fmt.Println("--- context.WithTimeout cancels workers automatically ---")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel() // always call cancel, even though the timeout will fire

	work := make(chan int)
	done := make(chan struct{})
	go func() {
		worker(ctx, 1, work)
		close(done)
	}()

	// Feed a couple of items, then go quiet - the worker should time out
	// and stop on its own after 50ms, without us telling it to explicitly.
	work <- 1
	work <- 2
	<-done
}

func manualCancelDemo() {
	fmt.Println("--- context.WithCancel: cancel triggered manually ---")
	ctx, cancel := context.WithCancel(context.Background())

	work := make(chan int)
	done := make(chan struct{})
	go func() {
		worker(ctx, 2, work)
		close(done)
	}()

	work <- 1
	cancel() // explicitly tell the worker to stop
	<-done
}

type ctxKey string

const requestIDKey ctxKey = "requestID"

func handle(ctx context.Context) {
	// Retrieving a request-scoped value. Note the unexported key type
	// (ctxKey) - this avoids collisions with other packages that also
	// stash values in the same context under a plain string key.
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Println("handling request:", reqID)
	}
}

func valueDemo() {
	fmt.Println("--- context.WithValue: request-scoped data (use sparingly) ---")
	ctx := context.WithValue(context.Background(), requestIDKey, "req-42")
	handle(ctx)
}

func main() {
	timeoutDemo()
	fmt.Println()

	manualCancelDemo()
	fmt.Println()

	valueDemo()
}

/*
Expected output:

--- context.WithTimeout cancels workers automatically ---
worker 1: processing 1
worker 1: processing 2
worker 1: stopping (context deadline exceeded)

--- context.WithCancel: cancel triggered manually ---
worker 2: processing 1
worker 2: stopping (context canceled)

--- context.WithValue: request-scoped data (use sparingly) ---
handling request: req-42
*/
