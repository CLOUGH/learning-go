# 01 — Fundamentals

A fast tour of the core language, just enough to make the concurrency
lessons make sense. If you already know Go basics, skim `main.go` and move
on — the real point of this curriculum starts at lesson 03.

## Things to notice while reading `main.go`

- **Static typing, but no type ceremony.** `x := 5` infers `int`. Closer to
  `auto` in C++ or Java's `var` than to dynamically-typed JS/PHP — the type
  is still fixed at compile time, Go just doesn't make you spell it out.
- **Structs + methods, not classes.** A method is just a function with a
  receiver (`func (p Point) String() string`). There's no inheritance —
  Go uses composition (embedding) instead. Structurally a `Point` is like a
  C struct, but unlike C, Go lets you attach functions to it directly
  without a separate free-floating function and an explicit first
  argument. There's no `class`, no constructors, no `this`/`self` beyond
  the receiver name you choose.
- **Interfaces are implicit.** A type satisfies an interface just by having
  the right methods — no `implements` keyword like Java/PHP, and no
  virtual base class + vtable setup like C++. It's structural typing,
  closer to how TypeScript interfaces work than to Java's nominal typing.
  This matters for concurrency later: things like `error` and `io.Reader`
  are interfaces used everywhere.
- **Errors are values, not exceptions.** Functions that can fail return an
  `error` as their last return value; the caller checks it with `if err !=
  nil`. There's no `try/catch` in normal control flow — closer to C's
  "return an error code" convention (or PHP's old `mysqli` error-return
  style) than to Java/JS/PHP/C++ exceptions.
- **Closures capture variables by reference.** This is the single most
  important fact to internalize before touching goroutines — a closure
  that captures a variable sees later mutations to it, exactly like a JS
  closure over a `let`/`var`, and exactly the mechanism behind a whole
  class of concurrency bugs (see lesson 09) — including one that will look
  very familiar if you've ever debugged `var i` inside a JS loop + callback.
  C/C++ has no equivalent without an explicit capture list or pointer, and
  Java has no closures over mutable local variables at all (captured
  locals must be effectively final).

Run it:

```sh
go run ./01-fundamentals
```

Want the fuller picture of the language instead of the quick tour above —
generics, embedding, error wrapping, `defer`/`panic`/`recover`, package
visibility, and more? That's [02-core-fundamentals](../02-core-fundamentals/README.md).

## Real-world use cases

These aren't warm-up syntax — every one of them shows up in ordinary
production Go within the first few files of most services:

- **Implicit interfaces** are why Go code is easy to test without a
  mocking framework: a function that takes an interface it needs (say,
  something with a `Save(User) error` method) can be handed a real
  database in production and a fake in-memory struct in tests, with zero
  `implements`/`@Mock` ceremony:

  ```go
  type UserStore interface {
      Save(u User) error
  }

  func Register(store UserStore, u User) error { return store.Save(u) }
  // production: Register(postgresStore, u)
  // test:       Register(&fakeStore{}, u)
  ```

- **`error` as a return value** is the shape of essentially every
  standard-library and third-party function that can fail — a database
  query, an HTTP call, a file read. Getting comfortable with `if err !=
  nil { return err }` here is what makes reading real Go code fluent later.
- **Closures capturing by reference** are exactly the mechanism behind
  `http.HandlerFunc` middleware and dependency injection without a
  framework — a handler "closes over" a logger or a config value from its
  enclosing scope instead of needing it injected through a container:

  ```go
  func withLogging(logger *log.Logger, next http.HandlerFunc) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          logger.Println(r.Method, r.URL.Path) // closes over logger
          next(w, r)
      }
  }
  ```

## Katas

Two short practice drills — see [katas/README.md](katas/README.md).

```sh
go test ./01-fundamentals/katas/...
```
