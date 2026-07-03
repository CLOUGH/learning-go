// Package solutions holds reference solutions for the 07-context katas.
// Try the kata yourself first - see ../README.md.
package solutions

import (
	"context"
	"time"
)

// Delay - kata 1.
func Delay(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RunWithTimeout - kata 2. f keeps running in the background if the
// timeout wins the race - this kata is about not blocking the CALLER past
// the timeout, not about cancelling f itself.
func RunWithTimeout(f func(), timeout time.Duration) error {
	done := make(chan struct{})
	go func() {
		defer close(done)
		f()
	}()

	select {
	case <-done:
		return nil
	case <-time.After(timeout):
		return context.DeadlineExceeded
	}
}
