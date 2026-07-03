# 05 — Select

`select` lets a goroutine wait on multiple channel operations at once, and
proceeds with whichever one is ready first. It's the switch statement of
concurrent Go.

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

## Katas

See [katas/README.md](katas/README.md).

```sh
go test -race ./05-select/katas/...
```
