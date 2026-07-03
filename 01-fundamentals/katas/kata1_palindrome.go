package katas

import "strings"

// IsPalindrome reports whether s reads the same forwards and backwards,
// ignoring case. Treat s as a sequence of bytes (no need to worry about
// multi-byte runes for this kata - that's covered in lesson 02).
//
// TODO: implement IsPalindrome.
func IsPalindrome(s string) bool {
	// panic("TODO: implement IsPalindrome")

	if len(s) == 0 {
		return true
	}

	// read in for ward order
	for i := 0; i < len(s); i++ {
		reverseIndex := len(s) - i - 1
		if !strings.EqualFold(s[i:i+1], s[reverseIndex:reverseIndex+1]) {
			return false
		}
	}
	return true
}
