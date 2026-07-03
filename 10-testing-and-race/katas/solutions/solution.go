// Package solutions holds reference solutions for the 10-testing-and-race
// katas. Try the kata yourself first - see ../README.md.
package solutions

import (
	"sync"
	"sync/atomic"
)

// AtomicCounter - kata 1.
type AtomicCounter struct {
	n int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.n, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.n)
}

// MutexCounter - given to you in the kata as the baseline to benchmark
// AtomicCounter against; reproduced here so this package is self-contained.
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
