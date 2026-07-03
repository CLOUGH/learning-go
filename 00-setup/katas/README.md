# Katas — 00 setup

Two tiny "spot and fix it" drills to get comfortable with the tools from
this lesson. Neither has a Go test to run — you verify these by running
the tool itself and confirming it's quiet.

## 1. `gofmt` (`gofmt/messy.go`)

The file is valid Go but badly formatted (inconsistent spacing, alignment,
brace placement). Find every issue and fix it.

```sh
gofmt -l 00-setup/katas/gofmt          # lists files that need formatting
gofmt -d 00-setup/katas/gofmt/messy.go # shows the diff gofmt would apply
gofmt -w 00-setup/katas/gofmt/messy.go # applies it
```

Try fixing it by hand first, then compare against what `gofmt -w` does.

## 2. `go vet` (`vet/vetty.go`)

The file compiles and runs, but has a bug `go vet` is specifically built to
catch: a `Printf`-style format verb that doesn't match its argument.

```sh
go vet ./00-setup/katas/vet/...
```

Read the warning, find the mismatched verb, and fix it so `go vet` is
silent.
