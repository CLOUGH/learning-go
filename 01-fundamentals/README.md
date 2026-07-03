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

## Katas

Two short practice drills — see [katas/README.md](katas/README.md).

```sh
go test ./01-fundamentals/katas/...
```
