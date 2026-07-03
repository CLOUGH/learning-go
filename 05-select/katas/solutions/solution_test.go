package solutions

import (
	"testing"
	"time"
)

func TestFirstFromA(t *testing.T) {
	a := make(chan int, 1)
	b := make(chan int)
	a <- 42

	if got := First(a, b); got != 42 {
		t.Errorf("First = %d, want 42", got)
	}
}

func TestFirstFromB(t *testing.T) {
	a := make(chan int)
	b := make(chan int)

	go func() {
		time.Sleep(10 * time.Millisecond)
		b <- 7
	}()

	if got := First(a, b); got != 7 {
		t.Errorf("First = %d, want 7", got)
	}
}

func TestWithTimeoutReceivesInTime(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 99

	v, ok := WithTimeout(ch, 500*time.Millisecond)
	if !ok || v != 99 {
		t.Fatalf("WithTimeout = (%d, %v), want (99, true)", v, ok)
	}
}

func TestWithTimeoutTimesOut(t *testing.T) {
	ch := make(chan int)

	start := time.Now()
	_, ok := WithTimeout(ch, 50*time.Millisecond)
	elapsed := time.Since(start)

	if ok {
		t.Fatal("expected ok=false on timeout")
	}
	if elapsed < 50*time.Millisecond {
		t.Fatalf("returned after only %v, expected to wait out the timeout", elapsed)
	}
}
