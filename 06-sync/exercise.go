package main

// Counter is a concurrency-safe counter. Multiple goroutines will call
// Inc concurrently; Value must always reflect exactly the number of Inc
// calls that have completed - no lost increments.
//
// TODO: add whatever field(s) you need (hint: a sync.Mutex and an int) and
// implement Inc and Value below.
type Counter struct {
	// TODO: fields
}

func NewCounter() *Counter {
	return &Counter{}
}

// TODO: implement Inc so it's safe to call from many goroutines at once.
func (c *Counter) Inc() {
	panic("TODO: implement Inc")
}

// TODO: implement Value so it's safe to call concurrently with Inc.
func (c *Counter) Value() int {
	panic("TODO: implement Value")
}
