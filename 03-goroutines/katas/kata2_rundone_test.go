package katas

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestRunDoneClosesWhenFinished(t *testing.T) {
	var ran int32

	done := RunDone(func() {
		time.Sleep(20 * time.Millisecond)
		atomic.StoreInt32(&ran, 1)
	})

	select {
	case <-done:
		if atomic.LoadInt32(&ran) != 1 {
			t.Fatal("done was closed before f finished running")
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for done to close")
	}
}

func TestRunDoneDoesNotBlock(t *testing.T) {
	start := time.Now()
	RunDone(func() { time.Sleep(200 * time.Millisecond) })
	if elapsed := time.Since(start); elapsed > 50*time.Millisecond {
		t.Fatalf("RunDone blocked for %v - it should return immediately", elapsed)
	}
}
