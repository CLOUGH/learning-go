package main

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

// waitForSignal is the same "select on ctx.Done() vs. the real event"
// shape from lessons 05 and 07 - the function under test doesn't know or
// care that it's being tested inside a synctest bubble.
func waitForSignal(ctx context.Context, sig <-chan struct{}) error {
	select {
	case <-sig:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Before testing/synctest, testing this honestly meant either actually
// waiting out the real 5-second timeout (slow test suite) or shrinking the
// timeout to milliseconds (flaky under CI load - a slow enough scheduler
// tick could still fail the assertion). synctest.Test runs the function
// inside an isolated "bubble" with a FAKE clock: time.Sleep, timers, and
// context deadlines all use virtual time that jumps forward instantly the
// moment every goroutine in the bubble is blocked waiting on something.
func TestWaitForSignalTimesOut(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		sig := make(chan struct{}) // nobody ever closes this - it should time out

		start := time.Now()
		err := waitForSignal(ctx, sig)
		elapsed := time.Since(start)

		if err != context.DeadlineExceeded {
			t.Fatalf("err = %v, want context.DeadlineExceeded", err)
		}
		// Within the bubble, this really did "wait" exactly 5 seconds of
		// FAKE time - deterministically, every single run - even though
		// the test itself completes in a few milliseconds of wall-clock time.
		if elapsed != 5*time.Second {
			t.Fatalf("elapsed = %v, want exactly 5s of fake time", elapsed)
		}
	})
}

func TestWaitForSignalReceivesInTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		sig := make(chan struct{})
		go func() {
			time.Sleep(1 * time.Second) // fake time - resolves instantly, still deterministically "1s" later
			close(sig)
		}()

		if err := waitForSignal(ctx, sig); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}

// synctest.Wait blocks until every other goroutine in the bubble is
// durably blocked (waiting on a channel/timer/WaitGroup - see the package
// doc for the exact list). It's how you deterministically assert "the
// goroutine I started has done its work" without a sleep OR a channel
// round-trip purely for synchronization.
func TestSynctestWait(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		done := false
		go func() {
			done = true
		}()

		synctest.Wait()
		if !done {
			t.Fatal("expected the goroutine to have finished by the time Wait returns")
		}
	})
}
