// Package solutions holds reference solutions for the 05-select katas.
// Try the kata yourself first - see ../README.md.
package solutions

import "time"

// First - kata 1.
func First(a, b <-chan int) int {
	select {
	case v := <-a:
		return v
	case v := <-b:
		return v
	}
}

// WithTimeout - kata 2.
func WithTimeout(ch <-chan int, d time.Duration) (int, bool) {
	select {
	case v := <-ch:
		return v, true
	case <-time.After(d):
		return 0, false
	}
}
