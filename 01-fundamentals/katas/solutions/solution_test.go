package solutions

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"a", true},
		{"racecar", true},
		{"RaceCar", true},
		{"hello", false},
		{"ab", false},
		{"aa", true},
	}
	for _, c := range cases {
		if got := IsPalindrome(c.in); got != c.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestNewCounter(t *testing.T) {
	next := NewCounter()
	for want := 1; want <= 3; want++ {
		if got := next(); got != want {
			t.Fatalf("next() = %d, want %d", got, want)
		}
	}
}

func TestNewCounterIndependence(t *testing.T) {
	a := NewCounter()
	b := NewCounter()
	a()
	a()
	if got := b(); got != 1 {
		t.Fatalf("a's calls leaked into b: b() = %d, want 1", got)
	}
}
