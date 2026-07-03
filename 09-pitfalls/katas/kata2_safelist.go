package katas

// SafeList is a concurrency-safe, append-only list of ints.
//
// TODO: add whatever fields you need and implement Add and Items below.
type SafeList struct {
	// TODO: fields
}

func NewSafeList() *SafeList {
	return &SafeList{}
}

// Add appends v to the list. Safe to call from many goroutines at once.
//
// TODO: implement Add.
func (l *SafeList) Add(v int) {
	panic("TODO: implement Add")
}

// Items returns a COPY of the current contents, safe to read/range over
// even if other goroutines call Add concurrently afterward - the caller
// must never be able to observe a torn or racing view of the internal
// slice, and must not be able to corrupt SafeList's internal state by
// mutating what Items() returned.
//
// TODO: implement Items.
func (l *SafeList) Items() []int {
	panic("TODO: implement Items")
}
