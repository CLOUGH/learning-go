package main

// Generator returns a channel that emits each value in nums, in order,
// then closes the channel. It must send the values from a goroutine so the
// caller can start consuming immediately (i.e. Generator itself must
// return right away, not block until all values are consumed).
//
// TODO: implement Generator.
func Generator(nums ...int) <-chan int {
	panic("TODO: implement Generator")
}

// Sum drains `in` completely (using range, so it naturally stops when the
// channel closes) and returns the sum of every value received.
//
// TODO: implement Sum.
func Sum(in <-chan int) int {
	panic("TODO: implement Sum")
}

// Pipeline wires Generator and Sum together: Sum(Generator(nums...)).
// It's given to you as a sanity check that the two functions above compose
// the way channels are meant to.
func Pipeline(nums ...int) int {
	return Sum(Generator(nums...))
}
