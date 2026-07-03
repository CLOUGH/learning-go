package solution

import (
	"reflect"
	"testing"
)

func TestFullPipeline(t *testing.T) {
	source := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	doubled := Transform(source, func(n int) int { return n * 2 })
	filtered := Filter(doubled, func(n int) bool { return n > 10 })

	got := Collect(filtered)
	want := []int{12, 14, 16, 18, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("pipeline: got %v, want %v", got, want)
	}
}
