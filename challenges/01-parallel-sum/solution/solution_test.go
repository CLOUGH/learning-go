package solution

import "testing"

func TestParallelSum(t *testing.T) {
	cases := []struct {
		nums       []int
		numWorkers int
		want       int
	}{
		{[]int{}, 4, 0},
		{[]int{7}, 4, 7},
		{[]int{1, 2, 3}, 8, 6},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 36},
		{[]int{1, 2, 3}, 0, 6},
	}
	for _, c := range cases {
		if got := ParallelSum(c.nums, c.numWorkers); got != c.want {
			t.Errorf("ParallelSum(%v, %d) = %d, want %d", c.nums, c.numWorkers, got, c.want)
		}
	}
}
