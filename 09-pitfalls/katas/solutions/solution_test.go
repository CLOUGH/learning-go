package solutions

import (
	"reflect"
	"sync"
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
	out := Publish(done, []int{1, 2, 3, 4, 5})

	<-out
	close(done)

	time.Sleep(50 * time.Millisecond)
}

func TestSafeListBasic(t *testing.T) {
	l := NewSafeList()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	got := l.Items()
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Items() = %v, want %v", got, want)
	}
}

func TestSafeListItemsIsACopy(t *testing.T) {
	l := NewSafeList()
	l.Add(1)

	items := l.Items()
	items[0] = 999

	got := l.Items()
	if got[0] != 1 {
		t.Fatalf("mutating the result of Items() leaked into SafeList's internal state: %v", got)
	}
}

func TestSafeListConcurrentAccess(t *testing.T) {
	l := NewSafeList()
	var wg sync.WaitGroup

	const goroutines = 50
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			l.Add(n)
			_ = l.Items()
		}(i)
	}
	wg.Wait()

	if got := len(l.Items()); got != goroutines {
		t.Errorf("len(Items()) = %d, want %d", got, goroutines)
	}
}
