// Package solutions holds reference solutions for the 02-core-fundamentals
// katas. Try the kata yourself first - see ../README.md.
package solutions

// Dedupe - kata 1.
func Dedupe[T comparable](items []T) []T {
	seen := make(map[T]struct{}, len(items))
	out := make([]T, 0, len(items))
	for _, item := range items {
		if _, ok := seen[item]; ok {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	return out
}

// SumPointers - kata 2.
func SumPointers(nums []*int) int {
	total := 0
	for _, p := range nums {
		if p != nil {
			total += *p
		}
	}
	return total
}
