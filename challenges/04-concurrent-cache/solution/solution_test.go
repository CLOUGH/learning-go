package solution

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentAccess(t *testing.T) {
	c := NewCache()
	var wg sync.WaitGroup

	const goroutines = 50
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id%10)
			for j := 0; j < 100; j++ {
				c.Set(key, j)
				c.Get(key)
				c.Len()
			}
		}(i)
	}
	wg.Wait()
}
