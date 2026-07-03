package main

import "testing"

func TestGeneratorAndSum(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{in: []int{}, want: 0},
		{in: []int{1, 2, 3}, want: 6},
		{in: []int{-5, 5, 10}, want: 10},
	}

	for _, c := range cases {
		got := Pipeline(c.in...)
		if got != c.want {
			t.Errorf("Pipeline(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestGeneratorClosesChannel(t *testing.T) {
	ch := Generator(1, 2)
	<-ch
	<-ch
	_, ok := <-ch
	if ok {
		t.Fatal("expected channel to be closed after all values consumed")
	}
}
