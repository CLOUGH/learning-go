// Package solutions holds reference solutions for the 06-sync katas. Try
// the kata yourself first - see ../README.md.
package solutions

import "sync"

// OnceValue - kata 1.
func OnceValue[T any](f func() T) func() T {
	var (
		once  sync.Once
		value T
	)
	return func() T {
		once.Do(func() {
			value = f()
		})
		return value
	}
}

// VisitorSet - kata 2.
type VisitorSet struct {
	mu   sync.Mutex
	seen map[string]struct{}
}

func NewVisitorSet() *VisitorSet {
	return &VisitorSet{seen: make(map[string]struct{})}
}

func (v *VisitorSet) Add(id string) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.seen[id] = struct{}{}
}

func (v *VisitorSet) Count() int {
	v.mu.Lock()
	defer v.mu.Unlock()
	return len(v.seen)
}
