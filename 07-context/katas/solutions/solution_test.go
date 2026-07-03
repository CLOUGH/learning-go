package solutions

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestDelayCompletesNormally(t *testing.T) {
	if err := Delay(context.Background(), 20*time.Millisecond); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDelayCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	start := time.Now()
	err := Delay(ctx, time.Hour)
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected an error from a cancelled context")
	}
	if elapsed > time.Second {
		t.Fatalf("Delay took %v - it didn't respect cancellation", elapsed)
	}
}

func TestDelayTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	if err := Delay(ctx, time.Hour); err == nil {
		t.Fatal("expected a deadline-exceeded error")
	}
}

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
		time.Sleep(time.Hour)
	}, 30*time.Millisecond)
	elapsed := time.Since(start)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v, want context.DeadlineExceeded", err)
	}
	if elapsed > 2*time.Second {
		t.Fatalf("RunWithTimeout took %v - it should have returned right after the timeout", elapsed)
	}
}
