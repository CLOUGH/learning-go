package katas

import (
	"context"
	"testing"
	"time"
)

func TestDelayCompletesNormally(t *testing.T) {
	err := Delay(context.Background(), 20*time.Millisecond)
	if err != nil {
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
	err := Delay(ctx, time.Hour) // would hang for an hour if cancellation isn't respected
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

	err := Delay(ctx, time.Hour)
	if err == nil {
		t.Fatal("expected a deadline-exceeded error")
	}
}
