// Package solutions holds reference solutions for the 01-fundamentals
// katas. Try the kata yourself first - see ../README.md.
package solutions

import "strings"

// IsPalindrome - kata 1.
func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if !strings.EqualFold(s[i:i+1], s[j:j+1]) {
			return false
		}
	}
	return true
}

// NewCounter - kata 2. The closure captures `n` by reference, so each
// call sees (and mutates) the same variable - and each call to NewCounter
// creates a brand new `n`, independent of any other counter.
func NewCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
