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
an equivalent Java program (`new Thread(...)`, or even a
`ExecutorService` thread pool) or C/C++ program (`pthread_create`,
`std::thread`) would be careful about spinning up thousands of OS threads.

If your mental model of "concurrency" comes from JavaScript, recalibrate:
JS is single-threaded with an event loop — `async`/`await` and Promises
give you *concurrency* (interleaving), but never true *parallelism* (two
pieces of your JS never run at the exact same instant, only Node's
underlying C++ thread pool for I/O does). Goroutines are genuinely
scheduled across multiple OS threads/CPU cores — real parallelism, not
just interleaving — which is exactly why unsynchronized access to shared
memory (lesson 06) is a real hazard in Go in a way it mostly isn't in
plain JS. PHP's classic model (one request = one process/thread, nothing
concurrent inside a single request) is even further from this — goroutines
have no direct PHP equivalent until you reach for newer, less common tools
like Fibers or a async extension.

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

`WaitGroup` plays the role of Java's `CountDownLatch` or calling `.join()`
on every `Thread`, or C++'s `std::thread::join()` — "block here until N
things finish." The closest JS analogue is `Promise.all([...])`, though
that's waiting on already-scheduled async work rather than explicitly
counting down.

### Go 1.25+: `wg.Go(f)`

Go 1.25 added a shortcut for the exact pattern above:

```go
// before (still perfectly valid, and what you'll see in any pre-1.25 codebase):
wg.Add(1)
go func() {
    defer wg.Done()
    doWork()
}()

// Go 1.25+:
wg.Go(func() {
    doWork()
})
```

`wg.Go(f)` does precisely `Add(1)` + start a goroutine + `defer Done()`
around `f` — it's the same three concepts you just learned, just spelled
with less boilerplate and one fewer place to forget the `Done()`. It
doesn't change anything else about how `WaitGroup` behaves. `main.go` has
both versions side by side (`withWaitGroup` and `withWaitGroupGo`) so you
can see they do the same thing.

### A scheduler note: container-aware `GOMAXPROCS` (Go 1.25+)

`GOMAXPROCS` controls how many OS threads can run goroutines simultaneously
— it's the "N" in the M:N scheduler mentioned above. Historically it
defaulted to the number of logical CPUs Go could see, which caused a
real problem in containers (Docker/Kubernetes): a container limited to,
say, 2 CPU-equivalents on a 64-core host would still see `GOMAXPROCS=64`,
massively over-scheduling and hurting performance. As of Go 1.25, the
runtime reads Linux cgroup CPU limits and sets `GOMAXPROCS` accordingly
by default, updating it automatically if the limit changes at runtime.
You generally don't need to set `GOMAXPROCS` by hand anymore in a
containerized deployment because of this.

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

## Real-world use cases

- **An HTTP server is already doing this for you.** `net/http`'s server
  spawns one goroutine per incoming request — that's *why* a Go web
  service handles thousands of concurrent requests on ordinary hardware
  without you writing a single line of concurrency code yourself; it's
  the framework's default, not something you opt into.
- **Fire-and-forget background work** — sending a confirmation email,
  writing an audit log entry, pushing an analytics event — after
  responding to a request, so the caller isn't kept waiting on work they
  don't need to see the result of:

  ```go
  func handleSignup(w http.ResponseWriter, r *http.Request) {
      user := createUser(r)
      go sendWelcomeEmail(user) // don't make the HTTP response wait on this
      w.WriteHeader(http.StatusCreated)
  }
  ```

  (Lesson 09 covers why this specific pattern needs care in a real
  service — a goroutine like this outlives the request and needs its own
  way to be cancelled if the process is shutting down.)
- **`wg.Go`/`WaitGroup`** is exactly the shape of "fan out N independent
  API calls or file reads, then proceed once they've all finished" — e.g.
  fetching a user's profile, orders, and preferences from three services
  in parallel before rendering a dashboard, instead of three sequential
  round trips.

## Katas

Two more practice drills beyond the exercise above — see
[katas/README.md](katas/README.md).

```sh
go test -race ./03-goroutines/katas/...
```
