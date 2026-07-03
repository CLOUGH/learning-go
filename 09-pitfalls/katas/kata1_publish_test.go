package katas

import (
	"testing"
	"time"
)

func TestPublishSendsEverything(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	out := Publish(done, []int{1, 2, 3})

	var got []int
	deadline := time.After(2 * time.Second)
loop:
	for {
		select {
		case v, ok := <-out:
			if !ok {
				break loop
			}
			got = append(got, v)
		case <-deadline:
			t.Fatal("timed out waiting for values")
		}
	}

	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestPublishStopsWhenDoneCloses(t *testing.T) {
	done := make(chan struct{})
	// An unbuffered channel with more values than any receiver will ever
	// read: if Publish doesn't select on `done` for every send, its
	// goroutine blocks forever trying to send the 2nd value - a leak.
	out := Publish(done, []int{1, 2, 3, 4, 5})

	<-out       // receive exactly one value
	close(done) // then tell the publisher to stop

	// Give the (correctly implemented) goroutine a moment to notice
	// `done` and exit. There's no direct way to assert "no goroutine
	// leaked" from here, but if Publish is wrong, this test - and the
	// whole `go test` run - will hang until the test binary's timeout,
	// which is the point: a leak here is loud, not silent.
	time.Sleep(50 * time.Millisecond)
}
