# 00 — Setup

You already have Go installed (`go version` → 1.26+; `go.mod` at the repo
root declares `go 1.26`). A few things worth knowing before you start:

## Modules

This whole repo is one Go module (`go.mod` at the root, module `learn-go`).
Every lesson is its own `package main` in its own directory — that's normal
in Go: one directory = one package, and only one `package main` + `func
main()` per runnable directory.

`go.mod` plays the role `package.json` plays for npm, `pom.xml`/`build.gradle`
for Maven/Gradle, or `composer.json` for PHP: it names the module and pins
dependency versions (there's also a `go.sum`, like `package-lock.json`).
There's no build step to think about the way C/C++ has a Makefile/CMake to
wrangle — `go build`/`go run` compile the whole dependency graph directly,
and it's fast enough that "compile" barely feels like a separate step from
"run".

## Commands you'll use constantly

```sh
go run ./01-fundamentals        # compile + run a package
go build ./...                  # compile everything, catch errors, no run
go test ./...                   # run all tests in the module
go test -race ./...             # run tests with the race detector on
go vet ./...                    # static checks for suspicious code
gofmt -l .                      # list files that aren't formatted correctly
```

`-race` matters a lot for this curriculum: it instruments the binary to
detect concurrent unsynchronized access to memory. It's slower, but it turns
"my program had a subtle bug that only showed up sometimes" into "here's the
exact line and goroutine stack that raced." Get in the habit of running
tests with `-race` any time goroutines are involved — you'll use it from
lesson 03 onward.

If you've used C/C++, this is the same idea as Clang/GCC's ThreadSanitizer
(`-fsanitize=thread`) — Go's race detector is actually built on the same
underlying technology. Java and JS mostly don't need an equivalent: the JVM
and JS engines either synchronize memory model details for you or (for JS)
never run your code on more than one thread in the first place.

## A note on Go versions, 1.22 through 1.26

This curriculum was originally written against Go 1.22 and has since been
updated for Go 1.26 — both the installed toolchain and `go.mod`'s `go`
directive moved forward, and several lessons now cover language/stdlib
features that didn't exist in 1.22. Two version-specific things worth
knowing up front:

- **Go 1.22** changed a long-standing gotcha: in older Go, a `for` loop's
  variable was reused across iterations, which meant capturing the loop
  variable in a goroutine closure would often give you the wrong value. Go
  1.22 gives each iteration its own variable. Lesson
  [09-pitfalls](../09-pitfalls/README.md) covers this in detail, including
  why you'll still see the old workaround in code and tutorials written
  before 2024.
- **`go vet`'s `stdversion` analyzer** (added in 1.23) flags any use of a
  standard-library symbol that's newer than the `go` line in `go.mod` —
  e.g. it would catch `sync.WaitGroup.Go` (added in 1.25) being used in a
  module that still declared `go 1.22`. This is exactly why `go.mod` here
  was bumped to `go 1.26` alongside the toolchain upgrade, rather than left
  at 1.22 — some of the newer material genuinely needs it.

Where a lesson uses something newer than plain Go 1.22, it says so
explicitly and names the version that introduced it — see 03-goroutines
(`sync.WaitGroup.Go`, container-aware `GOMAXPROCS`), 02-core-fundamentals
(range-over-func iterators, `iter.Seq`, the iterator-based `slices`/`maps`
functions), and 10-testing-and-race (`testing/synctest`, `testing.B.Loop`).

## Katas

Two tiny "spot and fix it" drills for `gofmt` and `go vet` — see
[katas/README.md](katas/README.md).
