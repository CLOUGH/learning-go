// Package workerpool: a job-processing worker pool that respects
// context cancellation.
package workerpool

import "context"

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID int
	Value int
}

// Run starts `numWorkers` goroutines that pull Jobs from `jobs`, call
// `process` on each, and send the Result to the returned channel.
//
// Requirements:
//   - At most `numWorkers` calls to `process` run concurrently.
//   - The returned channel is closed once all workers have exited (so
//     callers can safely `for r := range Run(...)`).
//   - Workers stop pulling new jobs once `jobs` is closed and drained -
//     this is the normal, successful shutdown path.
//   - Workers ALSO stop early if ctx is cancelled - mid-flight `process`
//     calls don't need to abort, but a worker must not block forever
//     trying to receive a job, and must not block forever trying to SEND
//     a result if nobody is reading anymore because the caller gave up.
//     Both the job-receive and the result-send need to race against
//     ctx.Done().
//   - Run itself must return immediately (it starts goroutines and
//     returns the channel - it does not block waiting for workers).
//
// TODO: implement Run.
func Run(ctx context.Context, jobs <-chan Job, numWorkers int, process func(Job) Result) <-chan Result {
	panic("TODO: implement Run")
}
