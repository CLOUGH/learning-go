# 00 — Setup

You already have Go installed (`go version` → 1.22+). A few things worth
knowing before you start:

## Modules

This whole repo is one Go module (`go.mod` at the root, module `learn-go`).
Every lesson is its own `package main` in its own directory — that's normal
in Go: one directory = one package, and only one `package main` + `func
main()` per runnable directory.

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

## A note on Go version 1.22

Go 1.22 (what's installed here) changed a long-standing gotcha: in older Go,
a `for` loop's variable was reused across iterations, which meant capturing
the loop variable in a goroutine closure would often give you the wrong
value. Go 1.22 gives each iteration its own variable. Lesson
[09-pitfalls](../09-pitfalls/README.md) covers this in detail, including why
you'll still see the old workaround in code and tutorials written before
2024.

## Katas

Two tiny "spot and fix it" drills for `gofmt` and `go vet` — see
[katas/README.md](katas/README.md).
