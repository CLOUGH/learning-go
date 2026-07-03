# 03 — Goroutines

This is the lesson the whole curriculum is built around.

## What is a goroutine?

A goroutine is a function running concurrently with other goroutines, in
the same address space. You start one by putting `go` in front of a
function call:

```go
go doSomething()
```

That line returns *immediately* — it does not wait for `doSomething` to
finish. `doSomething` now runs concurrently with whatever comes after that
line.

Goroutines are **not** OS threads. Go's runtime multiplexes potentially
hundreds of thousands of goroutines onto a small number of OS threads (an
"M:N scheduler"). That's why they're cheap: a goroutine starts with a tiny
(≈2KB) growable stack, versus megabytes for an OS thread. This is the
whole reason Go programs can casually spin up thousands of goroutines where
an equivalent Java or C++ program would be careful about spinning up
thousands of threads.

## The catch: `main` doesn't wait

`func main()` is itself running in a goroutine (the "main goroutine"). When
`main` returns, the program exits — immediately, without waiting for any
other goroutines to finish. This is the first bug everyone writes:

```go
go fmt.Println("hello from goroutine")
fmt.Println("hello from main")
// program may exit before the goroutine ever runs
```

You need a way to wait. The tool for that is `sync.WaitGroup`:

- `wg.Add(n)` — say "there are `n` more goroutines to wait for."
- `wg.Done()` — call this when a goroutine finishes (usually via `defer`
  right after starting it).
- `wg.Wait()` — blocks until the count returns to zero.

`main.go` walks through the broken version, then the fixed version, then a
version that shows goroutines actually interleaving.

## What this lesson does *not* cover yet

Goroutines on their own have no way to communicate results back — the demo
below just prints from inside the goroutine. Getting data *out* of a
goroutine safely is what channels (lesson 04) and mutexes (lesson 06) are
for. `exercise.go` in this lesson nudges you toward discovering why that's
needed.

Run it:

```sh
go run ./03-goroutines
go test -race ./03-goroutines/...
```

## Katas

Two more practice drills beyond the exercise above — see
[katas/README.md](katas/README.md).

```sh
go test -race ./03-goroutines/katas/...
```
