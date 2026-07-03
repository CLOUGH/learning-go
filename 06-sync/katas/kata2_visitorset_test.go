package katas

import (
	"fmt"
	"sync"
	"testing"
)

func TestVisitorSetBasic(t *testing.T) {
	v := NewVisitorSet()
	v.Add("alice")
	v.Add("bob")
	v.Add("alice") // duplicate

	if got := v.Count(); got != 2 {
		t.Errorf("Count() = %d, want 2", got)
	}
}

// go test -race ./06-sync/katas/...
func TestVisitorSetConcurrent(t *testing.T) {
	v := NewVisitorSet()
	var wg sync.WaitGroup

	const goroutines = 50
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Only 10 distinct ids, so lots of duplicate Adds across goroutines.
			v.Add(fmt.Sprintf("visitor-%d", id%10))
		}(i)
	}
	wg.Wait()

	if got := v.Count(); got != 10 {
		t.Errorf("Count() = %d, want 10", got)
	}
}
