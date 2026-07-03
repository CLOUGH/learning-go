package main

import (
	"fmt"
	"sync"
)

func racyCounter() {
	fmt.Println("--- racy counter (run with `go test -race` on the exercise to see this class of bug caught) ---")
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // READ, then WRITE - not atomic, this races
		}()
	}
	wg.Wait()
	fmt.Println("expected 1000, got:", counter) // often less than 1000
}

func mutexCounter() {
	fmt.Println("--- fixed with sync.Mutex ---")
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println("expected 1000, got:", counter) // always 1000
}

type SafeCache struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeCache() *SafeCache {
	return &SafeCache{data: make(map[string]int)}
}

func (c *SafeCache) Set(key string, value int) {
	c.mu.Lock() // exclusive: writers block everyone
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *SafeCache) Get(key string) (int, bool) {
	c.mu.RLock() // shared: many readers can hold this at once
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func rwMutexDemo() {
	fmt.Println("--- sync.RWMutex: many concurrent readers, exclusive writer ---")
	cache := NewSafeCache()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		cache.Set("answer", 42)
	}()
	wg.Wait()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if v, ok := cache.Get("answer"); ok {
				fmt.Println("reader", id, "saw:", v)
			}
		}(i)
	}
	wg.Wait()
}

func onceDemo() {
	fmt.Println("--- sync.Once: init logic runs exactly once ---")
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("initializing (only one goroutine will ever print this)")
			})
		}(i)
	}
	wg.Wait()
}

func main() {
	racyCounter()
	fmt.Println()

	mutexCounter()
	fmt.Println()

	rwMutexDemo()
	fmt.Println()

	onceDemo()
}

/*
Expected output (see caveats below - the first block is racy on purpose):

--- racy counter (run with `go test -race` on the exercise to see this class of bug caught) ---
expected 1000, got: 960

--- fixed with sync.Mutex ---
expected 1000, got: 1000

--- sync.RWMutex: many concurrent readers, exclusive writer ---
reader 4 saw: 42
reader 2 saw: 42
reader 3 saw: 42
reader 0 saw: 42
reader 1 saw: 42

--- sync.Once: init logic runs exactly once ---
initializing (only one goroutine will ever print this)

Caveats:
  - racyCounter's "got: X" is genuinely different every run (960 here;
    could be anything up to 1000) - that's the data race in action, not
    a mistake in this comment.
  - The five "reader N saw: 42" lines print in whatever order the
    scheduler happens to run those goroutines - any permutation of
    0-4 is correct.
*/
