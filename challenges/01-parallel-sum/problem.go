// Package parallelsum: compute the sum of a large slice of ints, splitting
// the work across multiple goroutines instead of a single sequential loop.
package parallelsum

// ParallelSum returns the sum of all elements in nums, computed using
// `numWorkers` goroutines, each summing an contiguous chunk of the slice.
// The partial sums must be combined safely (no data race) into the final
// total.
//
// Requirements:
//   - Must produce the same result as a plain sequential sum.
//   - Must actually divide the work across `numWorkers` goroutines (not
//     just run everything in one goroutine).
//   - Must be safe under `go test -race`.
//   - Handle numWorkers <= 0 or len(nums) == 0 sensibly (don't panic;
//     0 elements sums to 0).
//
// Hint: give each goroutine its own chunk and its own local accumulator,
// then combine the per-goroutine totals (e.g. via a channel, or via a
// mutex-protected total, or by writing each goroutine's partial sum into
// its own slot in a results slice indexed by worker number).
//
// TODO: implement ParallelSum.
func ParallelSum(nums []int, numWorkers int) int {
	panic("TODO: implement ParallelSum")
}
