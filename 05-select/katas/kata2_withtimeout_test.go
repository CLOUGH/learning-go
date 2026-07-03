package katas

import (
	"testing"
	"time"
)

func TestWithTimeoutReceivesInTime(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 99

	v, ok := WithTimeout(ch, 500*time.Millisecond)
	if !ok || v != 99 {
		t.Fatalf("WithTimeout = (%d, %v), want (99, true)", v, ok)
	}
}

func TestWithTimeoutTimesOut(t *testing.T) {
	ch := make(chan int) // nobody ever sends

	start := time.Now()
	_, ok := WithTimeout(ch, 50*time.Millisecond)
	elapsed := time.Since(start)

	if ok {
		t.Fatal("expected ok=false on timeout")
	}
	if elapsed < 50*time.Millisecond {
		t.Fatalf("returned after only %v, expected to wait out the timeout", elapsed)
	}
	if elapsed > 2*time.Second {
		t.Fatalf("took %v - looks like it didn't actually time out", elapsed)
	}
}
