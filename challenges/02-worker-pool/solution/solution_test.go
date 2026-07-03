package solution

import (
	"context"
	"testing"
	"time"
)

func TestRunProcessesAllJobs(t *testing.T) {
	const numJobs = 20
	jobs := make(chan Job, numJobs)
	for i := 0; i < numJobs; i++ {
		jobs <- Job{ID: i, Value: i}
	}
	close(jobs)

	results := Run(context.Background(), jobs, 4, func(j Job) Result {
		return Result{JobID: j.ID, Value: j.Value * 2}
	})

	var got []Result
	deadline := time.After(2 * time.Second)
loop:
	for {
		select {
		case r, ok := <-results:
			if !ok {
				break loop
			}
			got = append(got, r)
		case <-deadline:
			t.Fatal("timed out")
		}
	}

	if len(got) != numJobs {
		t.Fatalf("got %d results, want %d", len(got), numJobs)
	}
}
