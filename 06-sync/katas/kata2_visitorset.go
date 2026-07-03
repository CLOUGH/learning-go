package katas

// VisitorSet tracks the set of distinct visitor IDs seen so far, safe for
// concurrent use by many goroutines at once.
//
// TODO: add whatever fields you need and implement Add and Count below.
type VisitorSet struct {
	// TODO: fields
}

func NewVisitorSet() *VisitorSet {
	return &VisitorSet{}
}

// Add records id as having visited. Adding the same id more than once
// (including concurrently from different goroutines) must not increase
// Count beyond 1 for that id.
//
// TODO: implement Add.
func (v *VisitorSet) Add(id string) {
	panic("TODO: implement Add")
}

// Count returns the number of distinct ids added so far.
//
// TODO: implement Count.
func (v *VisitorSet) Count() int {
	panic("TODO: implement Count")
}
