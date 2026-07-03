package katas

// AtomicCounter is a concurrency-safe counter implemented using
// sync/atomic instead of a mutex - appropriate for a single integer like
// this, where there's no multi-step invariant to protect (compare with
// the mutex-based Counter from lesson 06, which is the right tool once
// more than a single value needs to change together).
//
// TODO: add whatever field you need (hint: int64, so you can use the
// sync/atomic functions/types that operate on int64) and implement Inc
// and Value below.
type AtomicCounter struct {
	// TODO: fields
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

// TODO: implement Inc using sync/atomic (e.g. atomic.AddInt64).
func (c *AtomicCounter) Inc() {
	panic("TODO: implement Inc")
}

// TODO: implement Value using sync/atomic (e.g. atomic.LoadInt64).
func (c *AtomicCounter) Value() int64 {
	panic("TODO: implement Value")
}
