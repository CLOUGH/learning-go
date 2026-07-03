package katas

import (
	"reflect"
	"sync"
	"testing"
)

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
	items[0] = 999 // mutating the returned slice must not affect the list

	got := l.Items()
	if got[0] != 1 {
		t.Fatalf("mutating the result of Items() leaked into SafeList's internal state: %v", got)
	}
}

// go test -race ./09-pitfalls/katas/...
func TestSafeListConcurrentAccess(t *testing.T) {
	l := NewSafeList()
	var wg sync.WaitGroup

	const goroutines = 50
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			l.Add(n)
			_ = l.Items() // concurrent reads while other goroutines are still adding
		}(i)
	}
	wg.Wait()

	if got := len(l.Items()); got != goroutines {
		t.Errorf("len(Items()) = %d, want %d", got, goroutines)
	}
}
