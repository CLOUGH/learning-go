package parallelsum

import (
	"testing"
)

func sequentialSum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func TestParallelSum(t *testing.T) {
	cases := []struct {
		name       string
		nums       []int
		numWorkers int
	}{
		{"empty", []int{}, 4},
		{"single element", []int{7}, 4},
		{"fewer elements than workers", []int{1, 2, 3}, 8},
		{"evenly divides", []int{1, 2, 3, 4, 5, 6, 7, 8}, 4},
		{"does not evenly divide", makeRange(1, 100), 7},
		{"zero workers falls back sensibly", []int{1, 2, 3}, 0},
		{"one worker", makeRange(1, 50), 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want := sequentialSum(c.nums)
			got := ParallelSum(c.nums, c.numWorkers)
			if got != want {
				t.Errorf("ParallelSum(%v, %d) = %d, want %d", c.nums, c.numWorkers, got, want)
			}
		})
	}
}

func TestParallelSumLarge(t *testing.T) {
	nums := makeRange(1, 1_000_000)
	want := sequentialSum(nums)
	got := ParallelSum(nums, 8)
	if got != want {
		t.Errorf("ParallelSum on 1,000,000 elements = %d, want %d", got, want)
	}
}

func makeRange(start, end int) []int {
	nums := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		nums = append(nums, i)
	}
	return nums
}
