# 04 — Channels

Channels are how goroutines talk to each other. Go's philosophy, stated in
the famous slogan:

> Don't communicate by sharing memory; share memory by communicating.

Instead of multiple goroutines fighting over the same variable (needing
locks), you pass values through a channel, and only one goroutine touches
the data at a time.

## Declaring and using a channel

```go
ch := make(chan int)       // unbuffered channel of ints
ch <- 5                    // send 5 into ch (blocks until someone receives)
v := <-ch                  // receive a value from ch (blocks until someone sends)
```

## Unbuffered vs. buffered

- **Unbuffered** (`make(chan int)`): a send blocks until a receiver is
  ready, and vice versa. This is a *rendezvous* — it synchronizes the two
  goroutines at that point, not just transfers data.
- **Buffered** (`make(chan int, 3)`): a send only blocks if the buffer is
  full; a receive only blocks if the buffer is empty. Useful for decoupling
  producer/consumer speed, but don't reach for a buffer just to "fix" a
  deadlock without understanding why it was blocking.

## Direction

Function signatures can restrict a channel to send-only or receive-only,
which the compiler enforces:

```go
func produce(out chan<- int)  { out <- 1 }   // can only send to out
func consume(in <-chan int) int { return <-in } // can only receive from in
```

This is a correctness tool, not a performance one — it stops you from
accidentally sending on a channel that's meant to be read-only in that
function, etc.

## Closing a channel

`close(ch)` signals "no more values are coming." Rules that matter:

- Only the **sender** should close a channel, never the receiver.
- Sending on a closed channel panics.
- Closing an already-closed channel panics.
- Receiving from a closed channel never blocks: it returns the zero value
  immediately. Use the two-value form to tell "closed" apart from "a real
  zero was sent": `v, ok := <-ch` — `ok` is `false` once the channel is
  closed and drained.
- `for v := range ch` receives values until the channel is closed, then
  exits the loop automatically. This is the idiomatic way to consume a
  channel.

## Nil channels

A `nil` channel (a `chan` variable that was never `make`'d) blocks forever
on both send and receive. This looks like a bug but is actually used
deliberately in `select` statements (lesson 05) to disable a case.

Run it:

```sh
go run ./04-channels
go test -race ./04-channels/...
```

## Katas

Two more practice drills beyond the exercise above — see
[katas/README.md](katas/README.md).

```sh
go test -race ./04-channels/katas/...
```
