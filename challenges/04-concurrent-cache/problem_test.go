package concurrentcache

import (
	"fmt"
	"sync"
	"testing"
)

func TestBasicGetSet(t *testing.T) {
	c := NewCache()
	if _, ok := c.Get("missing"); ok {
		t.Fatal("expected Get on missing key to return ok=false")
	}

	c.Set("a", 1)
	v, ok := c.Get("a")
	if !ok || v != 1 {
		t.Fatalf("Get(a) = (%d, %v), want (1, true)", v, ok)
	}

	c.Set("a", 2) // overwrite
	v, ok = c.Get("a")
	if !ok || v != 2 {
		t.Fatalf("Get(a) after overwrite = (%d, %v), want (2, true)", v, ok)
	}

	if c.Len() != 1 {
		t.Fatalf("Len() = %d, want 1", c.Len())
	}
}

func TestDelete(t *testing.T) {
	c := NewCache()
	c.Set("a", 1)
	c.Delete("a")
	if _, ok := c.Get("a"); ok {
		t.Fatal("expected key to be gone after Delete")
	}
	c.Delete("a") // deleting an already-absent key must not panic
	if c.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", c.Len())
	}
}

// Run with: go test -race ./challenges/04-concurrent-cache/...
func TestConcurrentAccess(t *testing.T) {
	c := NewCache()
	var wg sync.WaitGroup

	const goroutines = 50
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id%10) // deliberate overlap across goroutines
			for j := 0; j < 100; j++ {
				c.Set(key, j)
				c.Get(key)
				c.Len()
				if j%10 == 0 {
					c.Delete(key)
				}
			}
		}(i)
	}
	wg.Wait()
	// No assertion on final contents - the goroutines race by design.
	// This test's only real job is to let `-race` catch unsynchronized access.
}
