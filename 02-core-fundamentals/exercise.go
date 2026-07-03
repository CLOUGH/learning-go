package main

import "errors"

// --- Task 1: generics ---
//
// Implement the three classic higher-order functions using type
// parameters. They should work for ANY element type - that's the point
// of generics: one implementation, checked at compile time, no `any` and
// no type assertions needed anywhere in these bodies.

// Map applies f to every element of s and returns the results, in order.
// Note T and U can be different types (e.g. mapping []int to []string).
//
// TODO: implement Map.
func Map[T, U any](s []T, f func(T) U) []U {
	panic("TODO: implement Map")
}

// Filter returns a new slice containing only the elements of s for which
// keep returns true, preserving order.
//
// TODO: implement Filter.
func Filter[T any](s []T, keep func(T) bool) []T {
	panic("TODO: implement Filter")
}

// Reduce folds s down to a single value, starting from `initial` and
// combining one element at a time with f (left to right).
//
// TODO: implement Reduce.
func Reduce[T, Acc any](s []T, initial Acc, f func(Acc, T) Acc) Acc {
	panic("TODO: implement Reduce")
}

// --- Task 2: custom errors ---
//
// A NotFoundError is returned when a lookup fails. It must satisfy the
// standard `error` interface (one method: Error() string), and its
// message must include the ID that wasn't found, e.g.:
//
//	`item 42 not found`
//
// TODO: give NotFoundError whatever field(s) it needs, and implement
// Error() string on it.
type NotFoundError struct {
	// TODO: fields
}

func (e *NotFoundError) Error() string {
	panic("TODO: implement Error")
}

// ErrPermissionDenied is a sentinel error - callers check for it by
// identity via errors.Is, not by inspecting its message.
var ErrPermissionDenied = errors.New("permission denied")

// FindItem simulates a lookup with two failure modes:
//   - id < 0            -> wrap ErrPermissionDenied with context via fmt.Errorf's %w
//   - id not in `items`  -> return a *NotFoundError for that id
//   - otherwise         -> return the item's name and a nil error
//
// TODO: implement FindItem. `items` maps id -> name.
func FindItem(items map[int]string, id int) (string, error) {
	panic("TODO: implement FindItem")
}
