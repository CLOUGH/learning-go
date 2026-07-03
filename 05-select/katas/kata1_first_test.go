package katas

import (
	"testing"
	"time"
)

func TestFirstFromA(t *testing.T) {
	a := make(chan int, 1)
	b := make(chan int)
	a <- 42

	got := First(a, b)
	if got != 42 {
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

	got := First(a, b)
	if got != 7 {
		t.Errorf("First = %d, want 7", got)
	}
}
