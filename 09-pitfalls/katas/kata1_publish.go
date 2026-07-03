package katas

// Publish sends each value in `values`, in order, on the returned
// channel, then closes it. If `done` is closed before every value has
// been sent, Publish must stop sending and return (i.e. the goroutine it
// starts must exit) instead of blocking forever trying to send a value
// nobody will ever receive - that would leak the goroutine, exactly the
// bug this lesson is about.
//
// TODO: implement Publish. (Hint: every send needs to race against
// done via select, not just be a plain `out <- v`.)
func Publish(done <-chan struct{}, values []int) <-chan int {
	panic("TODO: implement Publish")
}
