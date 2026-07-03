package main

import (
	"reflect"
	"testing"
)

func TestSquareAll(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{in: []int{}, want: []int{}},
		{in: []int{1, 2, 3}, want: []int{1, 4, 9}},
		{in: []int{-2, 0, 5}, want: []int{4, 0, 25}},
	}

	for _, c := range cases {
		got := SquareAll(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SquareAll(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
