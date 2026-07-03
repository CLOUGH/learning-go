package katas

import "time"

// RunWithTimeout runs f (which takes no arguments and returns nothing) in
// the background. If f finishes within `timeout`, RunWithTimeout returns
// nil. If `timeout` elapses first, it returns context.DeadlineExceeded
// and does NOT wait any further for f to finish (f keeps running in the
// background - this kata isn't about cancelling f, just about not
// blocking the caller past the timeout).
//
// TODO: implement RunWithTimeout.
func RunWithTimeout(f func(), timeout time.Duration) error {
	panic("TODO: implement RunWithTimeout")
}
