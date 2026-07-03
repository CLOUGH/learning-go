package katas

import "testing"

// Complete both benchmarks below using b.RunParallel, the way lesson 10's
// own main_test.go does for SafeAdder. Each should hammer Inc() from many
// goroutines at once - that's the whole point of comparing them: which
// one holds up better under concurrent contention.
//
// Run with:
//   go test -bench=. -benchmem ./10-testing-and-race/katas/...
// and compare the ns/op between the two.

// TODO: implement, using b.RunParallel to call (&AtomicCounter{}).Inc()
// repeatedly.
func BenchmarkAtomicCounterInc(b *testing.B) {
	panic("TODO: implement this benchmark")
}

// TODO: implement, using b.RunParallel to call (&MutexCounter{}).Inc()
// repeatedly.
func BenchmarkMutexCounterInc(b *testing.B) {
	panic("TODO: implement this benchmark")
}
