package katas

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunWithTimeoutFinishesInTime(t *testing.T) {
	err := RunWithTimeout(func() {
		time.Sleep(10 * time.Millisecond)
	}, 500*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRunWithTimeoutExceeded(t *testing.T) {
	start := time.Now()
	err := RunWithTimeout(func() {
		time.Sleep(time.Hour) // must not make the test actually wait an hour
	}, 30*time.Millisecond)
	elapsed := time.Since(start)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v, want context.DeadlineExceeded", err)
	}
	if elapsed > 2*time.Second {
		t.Fatalf("RunWithTimeout took %v - it should have returned right after the timeout", elapsed)
	}
}
