package katas

// Batch reads values from in and groups them into slices of up to size
// elements, sending each full batch on the returned channel as soon as
// it's ready. When in closes, Batch sends whatever partial batch it has
// accumulated (if any - it's fine to send nothing if there were zero
// leftover values) and then closes the returned channel.
//
// TODO: implement Batch. size will always be >= 1 in this kata's tests.
func Batch(in <-chan int, size int) <-chan []int {
	panic("TODO: implement Batch")
}
