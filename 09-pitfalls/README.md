# 09 — Pitfalls

Concurrency bugs are the main thing that separates "wrote some Go" from
"understands goroutines." Each of these is demonstrated in `main.go`,
one function per pitfall, with the broken version and the fix side by side.

## 1. Goroutine leaks

A goroutine that blocks forever (waiting on a channel nobody will ever
write to, or that nobody reads from) never gets garbage collected — it
just sits there, leaked, for the life of the program. This is the most
common concurrency bug in real Go services: a goroutine started per
request that gets stuck waiting on a channel because the caller gave up
early (timed out, client disconnected). Fix: always give a blocking
goroutine a way out, usually a `context.Context` or `done` channel it also
selects on.

## 2. Loop-variable capture (pre-1.22 vs. 1.22+)

Before Go 1.22, this printed `3 3 3` (or similar), not `0 1 2`:

```go
for i := 0; i < 3; i++ {
    go func() { fmt.Println(i) }()
}
```

`i` was one variable, reused every iteration; by the time the goroutines
actually ran, the loop had usually already finished and `i` was `3`. Go
1.22 (what's installed here) changed the language spec so each iteration
gets its own `i` — this now correctly prints `0 1 2` (in some order). You
still need to know the *old* behavior because: (a) any codebase/tutorial
targeting Go <1.22 relies on the old semantics and uses the classic
workaround of passing the variable in as a parameter (`go func(i int)
{...}(i)`), and (b) it's a very frequently asked interview/code-review
question.

## 3. Deadlock: all goroutines asleep

```go
ch := make(chan int)
ch <- 1 // unbuffered send with no receiver anywhere -> blocks forever
```

If every goroutine (including main) ends up blocked waiting on something
that will never happen, the Go runtime detects it and crashes with `fatal
error: all goroutines are asleep - deadlock!`. This is actually a gift —
much better than a silent hang — learn to read that message; it dumps
every goroutine's stack, which tells you exactly where each one is stuck.

## 4. Sending on / closing a closed channel

- Sending on a closed channel panics: `panic: send on closed channel`.
- Closing an already-closed channel panics: `panic: close of closed
  channel`.
- Rule: only the sender closes a channel, and only once. If multiple
  goroutines might all think they're "the one" that should close it, use
  a `sync.Once` around the close, or restructure so only one goroutine
  owns closing it.

## 5. Forgetting `wg.Done()`

If a goroutine you `wg.Add(1)`'d for panics or returns early *before*
calling `wg.Done()`, `wg.Wait()` blocks forever. Always `defer wg.Done()`
as the very first line after starting the goroutine, so it runs no matter
how the goroutine exits.

Run it:

```sh
go run ./09-pitfalls
```

(The deadlock example is commented out in `main` by default, since a real
deadlock crashes the whole program — uncomment `demonstrateDeadlock()` in
`main.go` and run it on its own to see the crash and stack dump.)

## Katas

Two "don't leak, don't race" drills — see [katas/README.md](katas/README.md).

```sh
go test -race ./09-pitfalls/katas/...
```
