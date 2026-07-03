package katas

// RunDone starts f in its own goroutine and immediately returns a channel
// that will be closed once f has finished running. RunDone itself must
// not block waiting for f.
//
// TODO: implement RunDone.
func RunDone(f func()) <-chan struct{} {
	panic("TODO: implement RunDone")
}
