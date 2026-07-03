# 08 — Concurrency patterns

These are the recurring shapes goroutines + channels get assembled into.
Each is its own runnable subfolder.

## [workerpool/](workerpool/main.go)

A fixed number of goroutines ("workers") pull jobs from a shared channel
and process them concurrently, capping how much work runs in parallel.
This is the pattern you reach for whenever you have N independent tasks
and want to bound concurrency instead of spawning one goroutine per task.
Directly analogous to Java's `ExecutorService`/`ThreadPoolExecutor`
(a fixed-size thread pool pulling from a task queue) or C++'s (less
commonly hand-rolled) thread pool pattern — Go just doesn't have a
built-in pool type, so this lesson shows you the channel-based way to
build one yourself.

```sh
go run ./08-patterns/workerpool
```

## [pipeline/](pipeline/main.go)

A chain of stages, each its own goroutine, connected by channels: stage 1's
output channel is stage 2's input channel, and so on. Data flows through
the pipeline stage by stage, and stages run concurrently with each other
(stage 2 can be processing item 1 while stage 1 produces item 2).

```sh
go run ./08-patterns/pipeline
```

## [fanoutfanin/](fanoutfanin/main.go)

**Fan-out**: multiple goroutines read from the same input channel to
parallelize work across them. **Fan-in**: multiple channels are merged into
a single output channel. Combined, this is how you parallelize a pipeline
stage that's slow, then converge results back into one stream.

```sh
go run ./08-patterns/fanoutfanin
```

## [ratelimiter/](ratelimiter/main.go)

Using `time.Ticker` to only allow one operation per tick, regardless of how
fast requests come in — the basic building block behind rate-limited API
clients.

```sh
go run ./08-patterns/ratelimiter
```

## [semaphore/](semaphore/main.go)

A buffered channel used purely for its capacity, as a counting semaphore:
acquiring a "slot" is sending into it, releasing is receiving from it. This
bounds how many goroutines run some section of code concurrently, without
a full worker-pool structure. Same concept as Java's `java.util.concurrent.
Semaphore` or C++'s `std::counting_semaphore` — Go just repurposes the
channel you already know instead of adding a dedicated semaphore type.

```sh
go run ./08-patterns/semaphore
```

## Katas

See [katas/README.md](katas/README.md).

```sh
go test -race ./08-patterns/katas/...
```
