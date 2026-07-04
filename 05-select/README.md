# 05 — Select

`select` lets a goroutine wait on multiple channel operations at once, and
proceeds with whichever one is ready first. It's the switch statement of
concurrent Go.

If you've written C, the name and idea are a deliberate nod to the Unix
`select()`/`poll()` syscalls (wait on multiple file descriptors, proceed
with whichever is ready) — same concept, one level up, for channels
instead of file descriptors. If you know JS, the timeout pattern below is
similar in *purpose* to `Promise.race([task, timeoutPromise])`, though
`select` also handles ongoing repeated waiting (inside a loop) far more
naturally than a one-shot `Promise.race`. Java and C++ have no
single built-in construct that unifies "wait on any of these" the way
`select` does — you'd typically poll, use a `CompletableFuture.anyOf`
(Java), or hand-roll it with condition variables.

```go
select {
case v := <-ch1:
    fmt.Println("got", v, "from ch1")
case v := <-ch2:
    fmt.Println("got", v, "from ch2")
case ch3 <- 42:
    fmt.Println("sent 42 on ch3")
default:
    fmt.Println("none of the above were ready right now")
}
```

Key rules:

- If multiple cases are ready simultaneously, Go picks one **at random** —
  this is deliberate, to stop you writing code that accidentally depends on
  case order.
- `default` makes the whole `select` non-blocking: if no other case is
  ready *right now*, `default` runs instead of waiting. Omit `default` and
  `select` blocks until one case becomes ready.
- A `select{}` with no cases at all blocks forever (occasionally used
  intentionally to park a goroutine, but rare).

## The timeout pattern

`time.After(d)` returns a channel that receives a value after duration `d`.
Combined with `select`, that gives you a timeout for free, without touching
the operation you're timing out:

```go
select {
case res := <-resultCh:
    fmt.Println("got result:", res)
case <-time.After(2 * time.Second):
    fmt.Println("timed out waiting for result")
}
```

## The cancellation ("done channel") pattern

A goroutine that should stop early when told to typically selects on its
normal work channel and a `done` channel:

```go
for {
    select {
    case <-done:
        return // told to stop
    case work := <-workCh:
        process(work)
    }
}
```

This is the manual version of what `context.Context` (lesson 07) formalizes
and makes composable across an entire call chain.

Run it:

```sh
go run ./05-select
```

## Real-world use cases

- **Bounding a slow downstream call** — the timeout pattern above is
  exactly how you'd guard an HTTP client call to a flaky third-party API:
  race the real response against `time.After` so one slow dependency
  can't hang your entire request indefinitely.
- **A single goroutine multiplexing several event sources** — e.g. a
  connection-handling loop that selects on incoming data, an outgoing
  write queue, a heartbeat ticker, and a shutdown signal all at once, all
  in one `for { select { ... } }` loop instead of four separate goroutines
  each needing their own synchronization.
- **Non-blocking checks with `default`** — polling "is there work
  available right now without waiting for it" (e.g. draining whatever's
  currently in a buffered channel before doing something else), which is
  the shape health-check and metrics-collection loops often take.

## Katas

See [katas/README.md](katas/README.md).

```sh
go test -race ./05-select/katas/...
```
