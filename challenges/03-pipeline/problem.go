// Package pipeline: build a generic, composable channel pipeline out of
// three stages - generate, transform, filter.
package pipeline

// Generate returns a channel that emits each value in nums, in order,
// then closes it. Must send from a goroutine so Generate itself returns
// immediately.
//
// TODO: implement Generate.
func Generate(nums ...int) <-chan int {
	panic("TODO: implement Generate")
}

// Transform reads every value from in, applies f, and sends the result to
// the returned channel. The returned channel must close once `in` is
// closed and drained.
//
// TODO: implement Transform.
func Transform(in <-chan int, f func(int) int) <-chan int {
	panic("TODO: implement Transform")
}

// Filter reads every value from in and forwards only the ones where
// keep(value) is true. The returned channel must close once `in` is
// closed and drained.
//
// TODO: implement Filter.
func Filter(in <-chan int, keep func(int) bool) <-chan int {
	panic("TODO: implement Filter")
}

// Collect drains a channel into a slice. Given to you - useful for tests
// and for trying this out in a `go run`-able main if you write one.
func Collect(in <-chan int) []int {
	var out []int
	for v := range in {
		out = append(out, v)
	}
	return out
}
