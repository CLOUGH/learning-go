package katas

// OnceValue wraps f so that, no matter how many times (or from how many
// goroutines) the returned function is called, f itself is only ever
// invoked once - every call to the returned function gets f's single
// return value.
//
// TODO: implement OnceValue using sync.Once.
func OnceValue[T any](f func() T) func() T {
	panic("TODO: implement OnceValue")
}
