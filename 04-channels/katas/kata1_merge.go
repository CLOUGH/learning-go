package katas

// Merge fans-in any number of input channels into a single output
// channel: every value sent on any of chans should appear on the
// returned channel. The returned channel must close once ALL of chans
// have been closed and drained - not as soon as the first one closes.
//
// TODO: implement Merge. (Hint: one goroutine per input channel, plus a
// WaitGroup to know when they've all finished, so you know when it's
// safe to close the output.)
func Merge(chans ...<-chan int) <-chan int {
	panic("TODO: implement Merge")
}
