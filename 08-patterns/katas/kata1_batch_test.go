package katas

import (
	"reflect"
	"testing"
	"time"
)

func intChan(vals ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range vals {
			ch <- v
		}
	}()
	return ch
}

func collectBatches(t *testing.T, ch <-chan []int) [][]int {
	t.Helper()
	var got [][]int
	deadline := time.After(2 * time.Second)
	for {
		select {
		case b, ok := <-ch:
			if !ok {
				return got
			}
			got = append(got, b)
		case <-deadline:
			t.Fatal("timed out - Batch never closed its output channel")
			return nil
		}
	}
}

func TestBatchEvenlyDivides(t *testing.T) {
	got := collectBatches(t, Batch(intChan(1, 2, 3, 4, 5, 6), 2))
	want := [][]int{{1, 2}, {3, 4}, {5, 6}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBatchWithLeftover(t *testing.T) {
	got := collectBatches(t, Batch(intChan(1, 2, 3, 4, 5), 2))
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBatchEmpty(t *testing.T) {
	got := collectBatches(t, Batch(intChan(), 3))
	if len(got) != 0 {
		t.Errorf("got %v, want no batches", got)
	}
}
