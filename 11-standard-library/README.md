# 11 — The standard library, a tour of the fundamentals

Go's standard library is unusually complete and unusually stable — a huge
amount of real-world Go code is written using nothing but stdlib packages,
no third-party dependencies at all. This lesson is a guided tour of the
packages you'll reach for constantly, so you know what's available before
you go looking for a library to do it for you. It's not concurrency-specific
(nothing here needs anything from lessons 03–10), so — like lesson
02 — you could read it any time.

Same format as the rest of the curriculum: one file per topic, `main.go`
calls each demo in sequence.

```sh
go run ./11-standard-library
```

## What's in each file

- **[01_strings_strconv.go](01_strings_strconv.go)** — `strings`:
  `Contains`/`Split`/`Join`/`Replace`/`TrimSpace`/`Fields`/`Builder`.
  `strconv`: converting between strings and numbers/bools
  (`Atoi`/`Itoa`/`ParseFloat`/`ParseBool`/`FormatInt`/`Quote`). If you've
  used Java, this is `String`'s instance methods plus
  `Integer.parseInt`/`String.valueOf` split into two packages of free
  functions instead of methods on the type; if you've used JS, it's
  `String.prototype` methods plus `parseInt`/`Number()`/`.toString()`;
  PHP's global `str_*`/`intval`/`floatval` functions are a close match too
  — Go just groups them into packages instead of putting them all in the
  global namespace.
- **[02_os_io_bufio.go](02_os_io_bufio.go)** — `os`: environment
  variables, args, real file I/O (`WriteFile`/`ReadFile`/`CreateTemp`).
  `io`: the `Reader`/`Writer` interfaces that an enormous amount of Go
  code is built around — a file, a network connection, an in-memory
  buffer, and `os.Stdin`/`Stdout` are all just `io.Reader`/`io.Writer`,
  so the same `io.Copy` works on all of them. `bufio`: buffered,
  line-oriented reading (`Scanner`) on top of any `io.Reader`. C's
  `<stdio.h>` (`FILE*`, `fread`/`fwrite`) and Java's
  `InputStream`/`OutputStream` are the closest analogues to
  `io.Reader`/`io.Writer` — Go's version is two one-method interfaces
  instead of a class hierarchy, so anything implementing `Read`/`Write`
  slots into every function that accepts one.
- **[03_time.go](03_time.go)** — `time.Time` (instants) and
  `time.Duration` (elapsed time), constructing/comparing/formatting dates.
  The formatting layout (`"2006-01-02 15:04:05"`) trips up everyone
  coming from another language the first time: Go doesn't use `strftime`
  verbs like `%Y-%m-%d` (C, PHP, Python) or token letters like `yyyy-MM-dd`
  (Java); the layout string IS an example, applied to one specific
  reference date, and wherever a token sits in your layout is where that
  component goes in the output.
- **[04_sort.go](04_sort.go)** — `sort.Ints`/`sort.Strings` for the
  built-ins, `sort.Slice` for your own types via an index-based less-func
  (the pattern you'll see in most Go code written before generics went
  mainstream), and a reminder of `slices.SortFunc` (lesson 02's generics
  file) as the newer, type-safe alternative for new code.
- **[05_encoding_json.go](05_encoding_json.go)** — `json.Marshal`/
  `Unmarshal`, struct tags controlling the JSON key names, `omitempty`,
  and what happens to fields the struct doesn't know about (silently
  dropped, unlike some languages' stricter decoders). Comparable to
  Java's Jackson/Gson, JS's built-in `JSON.parse`/`JSON.stringify` (Go's
  version needs the struct tags because Go doesn't have JS's dynamic
  objects — every field has to be declared), or PHP's `json_encode`/
  `json_decode`.
- **[06_net_http.go](06_net_http.go)** — a real HTTP client (`http.Get`)
  talking to a real local server (`httptest.NewServer` + an
  `http.HandlerFunc`) — self-contained, no real network dependency, which
  is also exactly how you'd test HTTP code for real. A production server
  normally uses `http.ListenAndServe` instead, which blocks forever
  serving requests. Go's net/http needs no framework (no Express/Flask/
  Spring Boot equivalent) for something this simple — a `Handler` is just
  a function, and the standard library's server is fast enough that many
  real Go services never reach for a web framework at all.
- **[07_log_slog.go](07_log_slog.go)** — `log` (the classic
  timestamp-prefixed logger, writing to stderr) and `log/slog` (Go 1.21+,
  structured key-value logging, with a JSON handler for machine-readable
  output). Comparable to Java's `java.util.logging`/SLF4J, or Node's
  `console.log` vs. a structured logger like `pino` — `slog` is Go's
  built-in answer to "I want key-value fields, not just a formatted
  string," without needing a third-party dependency.
- **[08_flag.go](08_flag.go)** — command-line flag parsing. Shown via
  `flag.NewFlagSet` with an explicit argument slice (rather than the
  package-level `flag.Parse()` + real `os.Args`) so the demo doesn't
  interfere with the process's actual arguments — which is also exactly
  how you'd unit-test flag-parsing code. Comparable to Python's `argparse`
  or Node's `process.argv` (usually paired with a library like `yargs`) —
  Go's version is smaller and built in, at the cost of being less
  full-featured (no automatic subcommands, for instance — those exist as
  third-party packages like `cobra`).

## Idiomatic usage, best practices, and things to avoid

- Do: reach for `strings.Builder` (not `+=` in a loop) whenever you're
  building up a string across more than a couple of concatenations — see
  02-core-fundamentals's note on why repeated `+=` reallocates every time.
- Do: treat `io.Reader`/`io.Writer` as the "give me the narrowest
  interface that works" default when writing a function that consumes or
  produces a stream of bytes — it'll work with a file, a network
  connection, an `httptest` response body, or an in-memory buffer without
  you writing it four times.
- Do: use `time.Since(start)` instead of manually subtracting two
  `time.Now()` calls — it's the idiomatic spelling and reads better at
  the call site.
- Avoid: comparing two `time.Time` values with `==` — use `.Equal()`
  instead. Two `time.Time` values can represent the exact same instant
  but compare unequal with `==` if they carry different internal
  monotonic-clock or location data; `.Equal()` accounts for this.
- Avoid: parsing dates with a hand-rolled layout you reconstruct from
  memory each time — the standard library exposes common ones as
  constants (`time.RFC3339`, `time.Kitchen`, ...) worth knowing exist
  before you write your own.
- Do: check the error return from `json.Unmarshal`/`json.Marshal` (and
  everything in this lesson, really) — errors are values in Go, and a
  discarded error is a silently-swallowed bug waiting to happen (see
  02-core-fundamentals's errors section).
- Avoid: logging with `fmt.Println` in real services — reach for `log`
  or `log/slog` instead, so output carries a timestamp (and, with
  `slog`, structured fields) automatically instead of you re-inventing
  that by hand in every call site.
- Avoid: calling `flag.Parse()` more than once, or after `os.Args` has
  already been consumed elsewhere — it's meant to be called exactly once,
  early in `main()`.

## Run it

```sh
go run ./11-standard-library
```
