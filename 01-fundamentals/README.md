# 01 — Fundamentals

A fast tour of the core language, just enough to make the concurrency
lessons make sense. If you already know Go basics, skim `main.go` and move
on — the real point of this curriculum starts at lesson 03.

## Things to notice while reading `main.go`

- **Static typing, but no type ceremony.** `x := 5` infers `int`.
- **Structs + methods, not classes.** A method is just a function with a
  receiver (`func (p Point) String() string`). There's no inheritance —
  Go uses composition (embedding) instead.
- **Interfaces are implicit.** A type satisfies an interface just by having
  the right methods — no `implements` keyword. This matters for
  concurrency later: things like `error` and `io.Reader` are interfaces
  used everywhere.
- **Errors are values, not exceptions.** Functions that can fail return an
  `error` as their last return value; the caller checks it with `if err !=
  nil`. There's no `try/catch` in normal control flow.
- **Closures capture variables by reference.** This is the single most
  important fact to internalize before touching goroutines — a closure
  that captures a variable sees later mutations to it, which is exactly
  the mechanism behind a whole class of concurrency bugs (see lesson 09).

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
