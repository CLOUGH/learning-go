package katas

// RunLimited runs every function in tasks to completion, running at most
// `limit` of them concurrently at any given moment, and does not return
// until all tasks have finished.
//
// TODO: implement RunLimited using a buffered channel as a counting
// semaphore (see 08-patterns/semaphore/main.go for the pattern this is
// based on).
func RunLimited(tasks []func(), limit int) {
	panic("TODO: implement RunLimited")
}
