package katas

import (
	"reflect"
	"testing"
	"time"
)

func collectWithDeadline(t *testing.T, ch <-chan int, timeout time.Duration) []int {
	t.Helper()
	var got []int
	deadline := time.After(timeout)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return got
			}
			got = append(got, v)
		case <-deadline:
			t.Fatal("timed out waiting for the channel to close")
			return nil
		}
	}
}

func TestTakeFewerThanAvailable(t *testing.T) {
	got := collectWithDeadline(t, Take(chanOfN(1, 2, 3, 4, 5), 3), 2*time.Second)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Take(...,3) = %v, want %v", got, want)
	}
}

func TestTakeMoreThanAvailable(t *testing.T) {
	got := collectWithDeadline(t, Take(chanOfN(1, 2), 5), 2*time.Second)
	want := []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Take(...,5) = %v, want %v", got, want)
	}
}

func TestTakeZero(t *testing.T) {
	got := collectWithDeadline(t, Take(chanOfN(1, 2, 3), 0), 2*time.Second)
	if len(got) != 0 {
		t.Errorf("Take(...,0) = %v, want empty", got)
	}
}

func chanOfN(vals ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range vals {
			ch <- v
		}
	}()
	return ch
}
