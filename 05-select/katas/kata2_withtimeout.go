package katas

import "time"

// WithTimeout waits for a value from ch, up to duration d. If a value
// arrives in time, it returns (value, true). If d elapses first, it
// returns (0, false) without waiting any longer.
//
// TODO: implement WithTimeout using select + time.After.
func WithTimeout(ch <-chan int, d time.Duration) (int, bool) {
	panic("TODO: implement WithTimeout")
}
