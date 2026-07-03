# 02 — Core fundamentals (deep dive)

Lesson 01 was a fast tour, just enough to read goroutine code. This lesson
is the fuller picture: the parts of the core language that come up
constantly once you're actually writing Go, including several things that
trip people up the first time (slice aliasing, method sets, the "nil
interface" gotcha). Nothing here is concurrency-specific — you could read
this lesson before or instead of the concurrency track and it would stand
on its own.

Each topic is its own file so you can jump around; `main.go` just calls
each demo in sequence. Read the file, run it, then open it and change
something to see what breaks.

```sh
go run ./02-core-fundamentals
```

## What's in each file

- **[01_variables_types.go](01_variables_types.go)** — `var` vs `:=`,
  zero values, `const` and `iota`, the numeric type zoo (`int8`..`int64`,
  `uint`, `float32/64`), explicit conversion (Go never converts types for
  you), untyped constants.
- **[02_arrays_slices_maps.go](02_arrays_slices_maps.go)** — arrays are
  fixed-size values (copying an array copies every element); slices are a
  *view* (pointer, length, capacity) over a backing array, which is why
  two slices can alias the same memory and why `append` sometimes mutates
  in place and sometimes doesn't. Maps: nil map reads vs. writes, the
  comma-ok idiom, why map iteration order is randomized on purpose.
- **[03_strings_runes_bytes.go](03_strings_runes_bytes.go)** — strings are
  immutable UTF-8 byte sequences; `len(s)` counts bytes, not characters;
  `range` over a string yields runes; converting to `[]rune`/`[]byte`;
  `strings.Builder` for efficient concatenation.
- **[04_pointers.go](04_pointers.go)** — `&`/`*`, when Go passes by value
  vs. when you need a pointer, nil pointer dereferences.
- **[05_structs_embedding.go](05_structs_embedding.go)** — struct
  literals, embedding (Go's answer to inheritance: composition with
  automatic field/method promotion), struct tags, comparing structs.
- **[06_methods.go](06_methods.go)** — value vs. pointer receivers, and
  the method-set gotcha: a value of type `T` only has `T`'s
  value-receiver methods, not `*T`'s pointer-receiver methods — which
  becomes visible (and confusing) the moment interfaces are involved.
- **[07_interfaces.go](07_interfaces.go)** — implicit satisfaction, `any`,
  type assertions and type switches, `fmt.Stringer`/`error` as the
  idiomatic examples, and the classic "typed nil in an interface is not a
  nil interface" trap.
- **[08_generics.go](08_generics.go)** — type parameters, constraints
  (`comparable`, `cmp.Ordered`), generic `Map`/`Filter`/`Reduce`, and a
  look at the stdlib's own generic `slices`/`maps` packages.
- **[09_errors.go](09_errors.go)** — wrapping errors with `fmt.Errorf` and
  `%w`, sentinel errors with `errors.Is`, custom error types with
  `errors.As`.
- **[10_defer_panic_recover.go](10_defer_panic_recover.go)** — `defer`
  runs LIFO, always runs even on panic; `recover` only does anything
  inside a deferred function; how these compose to give Go a clean way to
  handle unexpected failures without exceptions.
- **[shapes/](shapes/shapes.go)** — a tiny separate package, imported by
  `main.go`, to show exported vs. unexported identifiers for real: capital
  letter = visible outside the package, lowercase = not. Note `main.go`
  can call `shapes.NewCircle` and `c.Area()` but cannot reach
  `shapes.piApprox` — the compiler enforces this, not a linter.

## Exercise

`exercise.go` / `exercise_test.go` has two independent tasks:

1. Generic `Map`, `Filter`, `Reduce` — the classic exercise for getting
   comfortable with type parameters.
2. A custom error type plus a wrapping function, checked with
   `errors.Is` and `errors.As` — the realistic version of Go error
   handling beyond a plain `if err != nil`.

```sh
go test ./02-core-fundamentals/...
```

## Katas

Two more, smaller drills covering ground the exercise above doesn't
(generics + maps, and pointers) — see [katas/README.md](katas/README.md).

```sh
go test ./02-core-fundamentals/katas/...
```
