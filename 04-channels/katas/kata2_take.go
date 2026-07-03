package katas

// Take returns a channel that yields the first n values received from
// in, then closes. If in closes before producing n values, Take closes
// the output early (after forwarding whatever it did receive). Take must
// not block forever, and must not leak a goroutine that keeps blocking
// on a receive from `in` after Take itself has stopped.
//
// You don't need to keep draining `in` beyond the n values you take -
// this kata's tests only send exactly what's needed.
//
// TODO: implement Take.
func Take(in <-chan int, n int) <-chan int {
	panic("TODO: implement Take")
}
