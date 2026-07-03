package katas

import "sync"

// MutexCounter is given to you fully implemented, as the mutex-based
// baseline to benchmark AtomicCounter (kata 1) against.
type MutexCounter struct {
	mu sync.Mutex
	n  int64
}

func NewMutexCounter() *MutexCounter {
	return &MutexCounter{}
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}
