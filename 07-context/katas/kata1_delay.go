package katas

import (
	"context"
	"time"
)

// Delay waits for duration d, but returns early with ctx.Err() if ctx is
// cancelled or times out before d elapses. Returns nil if the full delay
// completes normally.
//
// TODO: implement Delay using select on time.After(d) and ctx.Done().
func Delay(ctx context.Context, d time.Duration) error {
	panic("TODO: implement Delay")
}
