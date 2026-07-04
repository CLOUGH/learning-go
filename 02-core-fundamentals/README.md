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
something to see what breaks. Once you've been through the mechanics
below, [Idiomatic usage, best practices, and things to avoid](#idiomatic-usage-best-practices-and-things-to-avoid)
covers what to actually *do* with them in real code.

```sh
go run ./02-core-fundamentals
```

## What's in each file

Comparisons below assume Java, JavaScript, some C++, some C, and some PHP —
adjust if that's not your background.

- **[01_variables_types.go](01_variables_types.go)** — `var` vs `:=`,
  zero values, `const` and `iota`, the numeric type zoo (`int8`..`int64`,
  `uint`, `float32/64`), explicit conversion (Go never converts types for
  you), untyped constants. `iota` is Go's answer to Java `enum`/C
  `#define`/C++ `enum class` — there's no dedicated enum keyword, you get
  the same effect from `const` + `iota`. Every variable having a zero
  value (never "uninitialized") is like Java fields, unlike C/C++ locals
  which are garbage until set.
- **[02_arrays_slices_maps.go](02_arrays_slices_maps.go)** — arrays are
  fixed-size values (copying an array copies every element, like C/C++
  `std::array` or a C array — NOT like a Java array, which is always a
  reference); slices are a *view* (pointer, length, capacity) over a
  backing array, which is why two slices can alias the same memory and
  why `append` sometimes mutates in place and sometimes doesn't — closest
  to a C++ `std::span` or `std::vector` grow semantics, and functionally
  similar to a Java `ArrayList`/JS `Array`/PHP indexed array, but with the
  aliasing behavior made very explicit instead of hidden. Maps: nil map
  reads vs. writes, the comma-ok idiom, why map iteration order is
  randomized on purpose (Java's `HashMap` and PHP's arrays both have this
  "don't rely on order" property too, for related reasons; JS `Map`
  and PHP's ordered arrays actually preserve insertion order, unlike Go).
  Getting a *deterministic* view of a map's keys is what
  [11_iterators.go](11_iterators.go) covers, via the `maps` package.
- **[03_strings_runes_bytes.go](03_strings_runes_bytes.go)** — strings are
  immutable UTF-8 byte sequences (like Java/JS/PHP strings: immutable);
  `len(s)` counts bytes, not characters — different from Java/JS, where
  `.length` counts UTF-16 code units, and from PHP, whose default string
  functions are byte-oriented like Go's `len`; `range` over a string
  yields runes; converting to `[]rune`/`[]byte`; `strings.Builder` for
  efficient concatenation (the same reason Java has `StringBuilder` and
  JS engines internally rope-optimize `+=` in a loop).
- **[04_pointers.go](04_pointers.go)** — `&`/`*`, when Go passes by value
  vs. when you need a pointer, nil pointer dereferences. This is the one
  concept here that's a direct match for C/C++ (`&`, `*`, no pointer
  arithmetic though, and no `free` — Go's garbage collector reclaims
  memory for you like Java/JS/PHP). Java, JS, and PHP have no explicit
  pointers at all — objects are always accessed through implicit
  references, so there's no equivalent of choosing value vs. pointer
  semantics for your own types the way Go (and C/C++) require.
- **[05_structs_embedding.go](05_structs_embedding.go)** — struct
  literals, embedding (Go's answer to inheritance: composition with
  automatic field/method promotion — no `extends`, no base classes;
  closer to PHP traits or "has-a" composition than to Java/C++ `extends`),
  struct tags (the mechanism `encoding/json` uses, similar in spirit to
  Java annotations but just a plain string, read via reflection), comparing
  structs with `==` (works out of the box, unlike Java where `==` on
  objects compares references and you need `.equals()`).
- **[06_methods.go](06_methods.go)** — value vs. pointer receivers, and
  the method-set gotcha: a value of type `T` only has `T`'s
  value-receiver methods, not `*T`'s pointer-receiver methods — which
  becomes visible (and confusing) the moment interfaces are involved. C++
  has a related distinction (calling a non-const method needs a non-const
  object/pointer), but Java/JS/PHP have nothing like this — every object
  is always accessed by reference, so there's no "value vs. pointer
  receiver" choice to make at all.
- **[07_interfaces.go](07_interfaces.go)** — implicit satisfaction, `any`
  (Go's equivalent of Java's `Object`/TypeScript's `any`/C++'s `void*` —
  but type-safe, recovered via type assertions instead of casts), type
  assertions and type switches, `fmt.Stringer`/`error` as the idiomatic
  examples, and the classic "typed nil in an interface is not a nil
  interface" trap — a Go-specific footgun with no real analogue in
  Java/JS/PHP (where `null`/`None` is just `null`) or C++ (where a raw
  `nullptr` doesn't carry hidden type info the way a Go interface value does).
- **[08_generics.go](08_generics.go)** — type parameters, constraints
  (`comparable`, `cmp.Ordered`), generic `Map`/`Filter`/`Reduce`, and a
  look at the stdlib's own generic `slices`/`maps` packages. Go generics
  (added in 1.18) are monomorphized at compile time like C++ templates
  (a real `Stack[int]` and `Stack[string]` are separately compiled, fast,
  type-safe code) rather than type-erased like Java generics (where
  `List<Integer>` is just `List` with casts inserted by the compiler).
  JS and PHP have no generics at all — everything's dynamically typed, so
  the problem generics solve doesn't really arise the same way.
- **[11_iterators.go](11_iterators.go)** *(Go 1.23+)* — range-over-func:
  writing your own iterator as a plain function shaped like
  `func(yield func(V) bool)` (the `iter.Seq[V]` type), so it can be used
  directly in a `for ... range` loop, plus the standard library's own
  iterator-returning functions (`slices.Values`, `slices.All`,
  `slices.Collect`, `slices.Sorted`, `maps.Keys`, `maps.Values`). If
  you've used C#, this is conceptually close to `IEnumerable<T>` +
  `yield return`; Java's `Iterator<T>`/`Iterable<T>` is the nearest analogue
  there, though Go's version is just a function value, not an interface
  with `hasNext()`/`next()`. JS generator functions (`function*` +
  `yield`) are the closest conceptual match of your four languages — Go's
  `yield` is a plain callback you invoke yourself rather than a language
  keyword that suspends execution, but the "produce values one at a time,
  consumer can stop early" shape is the same idea.
- **[09_errors.go](09_errors.go)** — wrapping errors with `fmt.Errorf` and
  `%w`, sentinel errors with `errors.Is`, custom error types with
  `errors.As`. `errors.Is`/`errors.As` walking a wrapped chain is the
  closest Go gets to Java/C++/PHP exception causes (`Throwable.getCause()`,
  `std::exception_ptr`, `Exception::getPrevious()`) — but it's still all
  plain return values, never thrown/caught.
- **[10_defer_panic_recover.go](10_defer_panic_recover.go)** — `defer`
  runs LIFO, always runs even on panic; `recover` only does anything
  inside a deferred function; how these compose to give Go a clean way to
  handle unexpected failures without exceptions. `defer` is close in
  spirit to Java's `try`/`finally` or C++ RAII destructors (guaranteed
  cleanup code), but it's a standalone statement, not tied to a `try`
  block. `panic`/`recover` LOOK like throw/catch (Java/JS/PHP/C++ all
  have those), but idiomatic Go reserves them for truly exceptional,
  unrecoverable situations — a normal, expected failure uses an `error`
  return, not a panic.
- **[shapes/](shapes/shapes.go)** — a tiny separate package, imported by
  `main.go`, to show exported vs. unexported identifiers for real: capital
  letter = visible outside the package, lowercase = not — no `public`/
  `private`/`protected` keywords like Java/C++/PHP, just capitalization,
  and it's enforced per-package rather than per-class (there's no
  "private to this type but visible to this file" in Go — it's "visible
  outside this package, or not," full stop). Note `main.go` can call
  `shapes.NewCircle` and `c.Area()` but cannot reach `shapes.piApprox` —
  the compiler enforces this, not a linter.

## Idiomatic usage, best practices, and things to avoid

The mechanics above tell you what's *legal*. This is what experienced Go
code actually does with it — and the mistakes that show up over and over
in code review.

**Variables & types**
- Do: use `:=` inside functions almost always; reach for `var` when you
  want the zero value on purpose, at package level, or when spelling out
  the type makes the code clearer than the right-hand side would.
- Do: default to `int` and `float64` unless you have a concrete reason
  (matching a binary format/protocol, a struct field layout at scale) to
  pick a sized type like `int32`/`uint8`.
- Avoid: reaching for a sized integer type for an ordinary counter or loop
  variable "to save memory" — it just adds overflow-bug surface for no
  benefit; `int` is the right default almost everywhere.
- Avoid: building elaborate `iota` const blocks for something that's only
  ever going to have two states — a `bool` is clearer.

**Arrays, slices, maps**
- Do: use slices for nearly everything; arrays are rare in idiomatic Go
  (fixed-size things like a hash digest, or when you deliberately want
  value/copy semantics).
- Do: preallocate with `make([]T, 0, n)` when you know the eventual size —
  avoids repeated reallocation as `append` grows the backing array.
- Do: use `copy()` (or `slices.Clone`) whenever you need data that's
  independent of another slice — never assume a sub-slice is safe from
  aliasing (see `02_arrays_slices_maps.go`'s `b`/`a` example).
- Avoid: returning a slice that aliases a struct's internal backing array
  from an exported method without copying it first — callers mutating
  what you handed them corrupts your internal state (this is exactly why
  the `SafeList` kata in lesson 09 requires `Items()` to return a copy).
- Avoid: relying on map iteration order for anything, ever — sort the keys
  yourself if order matters for output.

**Strings, runes, bytes**
- Do: use `range` for character-by-character work on a string, never a
  manually-indexed byte loop, unless you specifically want raw bytes.
- Do: use `strings.Builder` for concatenation inside a loop.
- Avoid: treating `len(s)` as a character count — it's bytes, and will
  silently give you the wrong answer the moment non-ASCII input shows up.
- Avoid: converting to `[]rune` in a hot path "just in case" — it
  allocates a full copy; only do it when you actually need random access
  to characters or a character count.

**Pointers**
- Do: use a pointer when the callee needs to mutate the caller's value,
  or the value is large enough that copying it is wasteful.
- Avoid: sprinkling pointers everywhere "for performance" on small
  structs — copying a small struct is often cheaper than the indirection
  and it keeps the data easier to reason about (no aliasing to track).
- Avoid: taking the address of a loop variable and stashing it (e.g. into
  a slice of pointers) without checking your Go version's loop semantics
  (see lesson 09) — historically the single most common pointer bug in Go.

**Structs & embedding**
- Do: use embedding for genuine composition ("built out of"), and give a
  type a constructor function (`NewX`) once it has invariants to enforce
  or unexported fields that must be initialized correctly.
- Avoid: reaching for embedding to fake class inheritance/polymorphism —
  if you find yourself wanting to override a promoted method the way you'd
  override a virtual method in Java/C++, that's usually a sign an
  interface is the right tool instead.
- Avoid: exporting every field by default "just in case" — keep fields
  unexported unless there's an actual reason external code needs direct
  access to them.

**Methods**
- Do: once any method on a type needs a pointer receiver, make all of
  that type's methods pointer receivers too, for consistency — mixing the
  two is the direct cause of the method-set-vs-interfaces gotcha in
  `06_methods.go`.

**Interfaces**
- Do: keep interfaces small. The standard-library proverb is "the bigger
  the interface, the weaker the abstraction" — `io.Reader` is one method
  and it's used everywhere for exactly that reason.
- Do: define an interface in the package that *consumes* it, not next to
  the type that implements it — you often don't need to declare the
  interface at all until a second implementation (like a test fake) shows up.
- Avoid: designing a large interface upfront "for flexibility" before you
  have a second implementation that needs it — you're speculating, and Go
  interfaces are cheap to add later since satisfaction is implicit.
- Avoid: having a constructor return an interface type when a concrete
  type would do — it hides the concrete type's other methods from callers
  for no benefit if there's only one implementation.
- Avoid: returning a concrete `*T` as an `error`/interface value without
  checking it's genuinely `nil` first — that's the typed-nil trap from
  `07_interfaces.go`; return a literal `nil`, not a nil-valued pointer.

**Generics**
- Do: reach for generics when you'd otherwise be duplicating real logic
  across types (a container, `Map`/`Filter`/`Reduce`-style helpers).
- Avoid: adding a type parameter to a function that only ever needs to
  work on one type today — plain, non-generic code is easier to read, and
  generics are easy to add later without breaking callers.
- Avoid: reaching for generics when what you actually want is an
  interface (dynamic dispatch over different behavior) rather than the
  same logic applied to different data types.

**Errors**
- Do: wrap with `%w` when you add context, so callers can still
  `errors.Is`/`errors.As` through your wrapping to the original cause.
- Do: use a sentinel error (`errors.New`) for a condition callers check
  for by identity; use a custom error type when callers need structured
  data out of the failure.
- Avoid: comparing errors with `==` — it breaks the moment any wrapping is
  involved (see `09_errors.go`'s final demo). Use `errors.Is`/`errors.As`.
- Avoid: swallowing an error (`_ = f()`) without a deliberate, commented
  reason — an ignored error is a bug waiting to happen silently.
- Avoid: stuttering wrapped messages ("failed to do X: failed to do Y:
  failed to do Z") — add context, don't just repeat "failed".

**defer / panic / recover**
- Do: `defer` a cleanup call immediately after the line that acquires the
  resource (`mu.Lock(); defer mu.Unlock()`, `f, _ := os.Open(...); defer
  f.Close()`) so it can never be forgotten on an early return.
- Avoid: `panic`/`recover` as a substitute for ordinary error handling —
  idiomatic Go reserves `panic` for programmer errors and truly
  unrecoverable situations; an expected failure (bad input, a missing
  file) returns an `error`.
- Avoid: `defer`ing inside a loop that runs many iterations (e.g.
  processing thousands of files) — each deferred call only runs when the
  *enclosing function* returns, not at the end of that loop iteration, so
  they all pile up and hold their resources open until the function exits.

**Packages & visibility**
- Do: keep the exported surface of a package as small as the callers
  actually need — you can always export more later; taking something back
  is a breaking change.
- Avoid: mutable package-level (global) state where you can help it — it
  makes code harder to test in isolation and, the moment goroutines are
  involved, an unsynchronized global is a data race waiting to happen.

## Real-world use cases

Where these specific mechanics show up once you're past tutorial code:

**Variables & types** — `iota` const blocks model closed sets of states
that map to a database enum or HTTP status family:

```go
type OrderStatus int

const (
    Pending OrderStatus = iota
    Shipped
    Delivered
    Cancelled
)
```

**Arrays, slices, maps** — a slice of structs read from a database query
or a JSON API response is the single most common data shape in a Go
service; `map[string]int` (or a struct value) is the default choice for
an in-memory counter/lookup table (word frequency, request counts per
endpoint, a config keyed by feature name).

**Strings, runes, bytes** — `strings.Builder` is what you reach for while
generating a CSV export or assembling a SQL query from clauses in a loop,
anywhere string concatenation would otherwise happen many times in a row.

**Pointers** — a pointer receiver is how a `*sql.DB`-backed repository
type mutates its internal connection pool state, and passing `*Config` to
a constructor instead of `Config` is how you let `nil` mean "use
defaults" for an optional dependency.

**Structs & embedding** — embedding `sync.Mutex` or a base `Model` struct
(with `ID`, `CreatedAt`, `UpdatedAt` fields) into every database-backed
type is a common pattern in Go ORMs/ORM-adjacent code:

```go
type Model struct {
    ID        int
    CreatedAt time.Time
}

type User struct {
    Model // promotes ID, CreatedAt onto User directly
    Name  string
}
```

**Methods** — a `String() string` method (satisfying `fmt.Stringer`) is
what makes a custom type print usefully in logs and error messages
instead of a raw struct dump — worth adding to almost any type with an
identity (an ID, a status enum).

**Interfaces** — this is the single most-used real-world pattern in the
whole lesson: define a small interface for whatever your code depends on
(a `PaymentGateway`, a `Clock`, a `UserStore`), so production code wires
up the real implementation and tests wire up a fake — no mocking
framework or dependency-injection container required, unlike Java/C# code
that typically needs Mockito/Moq for the same job.

**Generics** — a `Repository[T]` type is the classic real use: one
`Create`/`Get`/`List` implementation shared by every entity type in a
service instead of hand-writing `UserRepository`, `OrderRepository`,
`ProductRepository` with identical bodies:

```go
type Repository[T any] struct {
    items map[int]T
}

func (r *Repository[T]) Get(id int) (T, bool) {
    v, ok := r.items[id]
    return v, ok
}
```

**Errors** — a custom error type carrying structured data (a `ValidationError`
with a `Field` and `Message`) is how an HTTP handler turns a failure deep
in a call stack into the right status code and JSON body via `errors.As`,
without every layer in between needing to know about HTTP at all.

**defer / panic / recover** — `defer conn.Close()`/`defer tx.Rollback()`
right after opening a database connection or transaction is the standard
way production Go code guarantees cleanup on every exit path; `recover`
inside HTTP middleware is how a server turns a panic in one request
handler into a 500 response instead of crashing the whole process.

**Packages & visibility** — a package's exported types/functions are its
public API contract; keeping helper types and validation logic
unexported is how a `payments` package can change its internal
implementation without breaking every caller, the same discipline Java
gets from `public`/`package-private` and PHP from `public`/`private`.

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
