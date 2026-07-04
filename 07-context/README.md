# 07 — Context

`context.Context` is the standard way to carry cancellation signals,
deadlines, and (sparingly) request-scoped values across API boundaries and
goroutines. If you've done lesson 05, this formalizes the "done channel"
pattern into something composable across an entire call chain — an HTTP
handler, the database call it makes, the goroutines that call spawns — all
sharing one cancellation signal.

The closest thing you've likely used is JS's `AbortController`/
`AbortSignal` (`fetch(url, { signal })`) — same idea, a cancellation token
threaded through to whatever might need to stop early. Java's nearest
equivalents are `Future.cancel()` or a `CompletableFuture` with a timeout,
though neither is as pervasively threaded through every API the way `ctx`
is in Go. C/C++ and PHP have no standard-library equivalent at all — you'd
build this yourself (an atomic cancellation flag, checked periodically).

## Creating a context

```go
ctx := context.Background() // root context, usually created once in main() or at a request's entry point

ctx, cancel := context.WithCancel(ctx)   // derived context you can cancel manually
defer cancel()                           // always call cancel, even if you don't use it explicitly - see below

ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // auto-cancels after 2s
defer cancel()

ctx, cancel := context.WithDeadline(ctx, someTime) // auto-cancels at a specific time
defer cancel()
```

**Always call the returned `cancel` function**, even on the success path,
even if the timeout would have fired anyway. It releases resources
associated with the context; forgetting it is a (small) leak.
`defer cancel()` right after creating the context is the idiom.

## Respecting a context inside a goroutine

```go
func worker(ctx context.Context, work <-chan int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("cancelled:", ctx.Err())
            return
        case w := <-work:
            process(w)
        }
    }
}
```

`ctx.Done()` returns a channel that's closed when the context is cancelled
or times out. `ctx.Err()` then tells you why: `context.Canceled` or
`context.DeadlineExceeded`.

## Passing a context

Convention: `ctx` is always the **first parameter**, named `ctx`, never
stored in a struct field. Every function in a call chain that might block
or do I/O should accept and propagate a `ctx`.

## `context.WithValue`

Carries request-scoped data (a request ID, an auth token) across API
boundaries. Use it sparingly — it's not a general-purpose way to pass
optional parameters, and overusing it makes data flow hard to trace. If a
function needs a value to do its job, prefer passing it as a normal
argument.

Run it:

```sh
go run ./07-context
```

## Real-world use cases

- **Every `net/http` handler already receives one** — `r.Context()` on an
  incoming request is cancelled automatically the moment the client
  disconnects or the request times out, so passing that context down into
  a database query means the query itself gets cancelled instead of
  running to completion for a client that's no longer listening:

  ```go
  func handler(w http.ResponseWriter, r *http.Request) {
      ctx := r.Context()
      user, err := db.GetUserContext(ctx, userID) // cancelled if the client hangs up
      ...
  }
  ```

- **Bounding an outgoing API call's total time** — wrapping a call to a
  third-party service in `context.WithTimeout` so a single slow
  dependency can't make your own service hang indefinitely; this is the
  standard way Go services enforce their own SLAs on calls they don't
  control.
- **Carrying a request ID or trace ID** through `context.WithValue` so
  every log line and downstream call in a request's lifetime can be
  correlated back to the same incoming request — the main legitimate use
  of `WithValue` in most real codebases, alongside auth claims extracted
  once at the edge.

## Katas

See [katas/README.md](katas/README.md).

```sh
go test -race ./07-context/katas/...
```
