// Package solutions holds reference solutions for the 09-pitfalls katas.
// Try the kata yourself first - see ../README.md.
package solutions

import "sync"

// Publish - kata 1. Every send races against done via select, so a
// blocked send can never leak the goroutine once the caller stops
// listening - exactly the fix for the "goroutine leak" pitfall this
// lesson is about.
func Publish(done <-chan struct{}, values []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range values {
			select {
			case out <- v:
			case <-done:
				return
			}
		}
	}()
	return out
}

// SafeList - kata 2.
type SafeList struct {
	mu    sync.Mutex
	items []int
}

func NewSafeList() *SafeList {
	return &SafeList{}
}

func (l *SafeList) Add(v int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.items = append(l.items, v)
}

// Items returns a COPY, so a caller can never observe (or cause) a race by
// holding onto the result while another goroutine calls Add.
func (l *SafeList) Items() []int {
	l.mu.Lock()
	defer l.mu.Unlock()
	out := make([]int, len(l.items))
	copy(out, l.items)
	return out
}
