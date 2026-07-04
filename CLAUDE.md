# CLAUDE.md — context for continuing this project

This repo is a self-contained, hands-on Go curriculum built entirely by
Claude Code for one learner, working session by session. It is **not** a
production codebase — every file here exists to teach. If you're picking
this up in a new session (possibly on a different machine), read this
whole file before touching anything; it captures conventions that aren't
obvious from the code alone.

## Who this is for

The user is learning Go from a background in **Java, JavaScript, some
C++, some C, and some PHP**. Explanations throughout the repo (READMEs,
code comments) default to comparing new Go concepts against those
languages — do the same when discussing this project in chat, and when
adding new lesson content. Skip the comparison when a feature has no real
analogue (multiple return values, explicit conversion rules) — don't
force one.

## What this repo is

A numbered-lesson curriculum, `00-setup` through `11-standard-library`,
with a heavy concurrency focus (that's most of lessons 03–10) bookended by
core-language material (01, 02) and a standard-library tour (11). See the
root [README.md](README.md) for the full roadmap and how-to-use-this
instructions — don't duplicate that here, just know it exists and is kept
in sync with whatever lessons actually exist.

## Repo conventions (follow these for anything new)

- **One module.** `go.mod` at the root, module `learn-go`, currently
  `go 1.26`. Every lesson directory is its own `package main` (or, for
  02-core-fundamentals and 11-standard-library, split across several
  numbered files in one `package main`, orchestrated by `main.go`).
- **Per-lesson layout**: `README.md` (concept + idiomatic comparisons +
  "Run it"), `main.go` (+ numbered demo files for the two multi-file
  lessons), optionally `exercise.go`/`exercise_test.go` (one bigger
  practice problem), and always a `katas/` folder (1-2 smaller drills).
- **Every lesson's README also has a "Real-world use cases" section**,
  placed after "Idiomatic usage" where that section exists, otherwise
  right before "Katas"/"Exercise". It's short prose plus a small
  illustrative code snippet (not a full runnable demo, no expected-output
  comment needed) tying the lesson's concept to where it actually shows up
  in production Go code — an HTTP service, a CLI tool, a job queue, a
  cache. This is deliberately distinct from "Idiomatic usage": that
  section is about *how* to use a feature well once you've reached for
  it; this one is about *where/why* you'd reach for it at all.
- **Katas and exercises are stubs, not partial implementations.** Every
  unimplemented function body is exactly `panic("TODO: implement X")` —
  never a partial/guessed implementation. This is deliberate: `go test`
  on an unsolved kata/exercise MUST fail with a clean panic, never a
  compile error and never a silent wrong answer.
- **Solutions live in a sibling `solutions/` (katas) or `solution/`
  (challenges) subdirectory**, as a separate package, so they don't spoil
  the stub at a glance and so both can be tested independently:
  `go test ./NN-lesson/katas/solutions/...`. `00-setup`'s two tooling
  katas (gofmt/vet) are the exception — their "solution" is just the
  corrected file, checked with `diff`.
- **`challenges/`** holds bigger, standalone problems (not tied to one
  lesson) with the same stub/test/solution shape, one level up.
- **Every demo `.go` file ends with a trailing `/* Expected output: ... */`
  comment** showing real captured `go run` output, so the file can be read
  and understood without running it. Where output is inherently
  non-deterministic (goroutine scheduling order, timing, `os.Args`, log
  timestamps, map iteration order), the comment says so explicitly and
  states what invariant actually holds instead of a literal fixed order.
  **If you add or change a demo file, regenerate this comment from a real
  `go run` — don't hand-write guessed output.**
- **Comment style**: no comments explaining *what* code does; comments
  explain *why* (a gotcha, a non-obvious invariant, a version-specific
  behavior). This matches the general Claude Code style guidance, applied
  extra strictly here since the "why" comments are half the pedagogical
  content.
- **Language-comparison callouts** (Java/JS/C/C++/PHP) belong in
  `README.md` prose, not code comments — keep code comments focused on
  Go-specific "why."

## Go version state (important - read before assuming anything about tooling)

- The system's Go was upgraded mid-project from an apt-installed 1.22.2 to
  **Go 1.26.4**, installed the official way to `/usr/local/go` (the old
  `golang-go`/`golang-1.22-go` apt packages were removed). `PATH` was
  updated in both `~/.zshrc` (interactive shells) and `~/.zshenv`
  (non-interactive shells, which is what a coding-assistant's shell tool
  typically runs) to put `/usr/local/go/bin` first.
- **On a new machine**: if `go version` doesn't show 1.26+, `go.mod`'s
  `go 1.26` directive will trigger Go's automatic toolchain download (a
  feature of the `go` command itself) if the installed toolchain is older
  and network access is available - or it'll fail with an explicit error
  telling you to upgrade if `GOTOOLCHAIN=local` is set. Either way, don't
  silently "fix" this by downgrading `go.mod` - the newer directive is
  load-bearing (see below).
- **Why `go.mod` says 1.26, not 1.22**: lessons 02, 03, and 10 use
  language/stdlib features that flatly require it - range-over-func
  iterators need `go 1.23`+ in `go.mod` to even compile, and `go vet`'s
  `stdversion` analyzer (1.23+) flags stdlib symbols newer than the
  declared version (e.g. `sync.WaitGroup.Go`, added in 1.25). Don't lower
  this without checking what breaks.
- Version-specific features actually used in lesson content: range-over-func
  / `iter.Seq`/`Seq2` / iterator-returning `slices`/`maps` functions (1.23),
  `testing.B.Loop` (1.24), `sync.WaitGroup.Go`, container-aware
  `GOMAXPROCS`, `testing/synctest` (1.25), experimental `goroutineleak`
  pprof profile (1.26, mentioned not demoed).

## Standard verification loop

Run this after any change, before considering it done (matches what every
prior session in this repo has done):

```sh
go build ./...                  # must be completely clean
go vet ./...                    # must show EXACTLY ONE finding: 00-setup/katas/vet/vetty.go (intentional)
gofmt -l .                      # must show EXACTLY ONE finding: 00-setup/katas/gofmt/messy.go (intentional)
go test -race ./...             # every kata/exercise STUB fails with "panic: TODO..." (expected);
                                 # every solutions/ and challenges/*/solution package must pass;
                                 # 10-testing-and-race's TestRacyAdder fails on purpose (demonstrates -race)
```

If `go vet`/`gofmt` show anything beyond those two known findings, something
regressed - fix it before moving on.

## State as of the last session (2026-07-03)

Everything above reflects the current, finished state:

- Lessons 00-11 all exist, each with README + demo file(s) + katas +
  (where applicable) exercise, all with expected-output comments.
- `challenges/` has 4 problems (parallel sum, worker pool, pipeline,
  concurrent cache), each with a passing reference solution.
- Every kata across every lesson has a solution in `katas/solutions/`.
- Go toolchain is 1.26.4; `go.mod` targets `go 1.26`; lessons 02/03/09/10
  call out Go 1.23-1.26 features explicitly where relevant.
- Full repo verification (build/vet/gofmt/test -race) passes cleanly per
  the loop above.

## Open threads / things the user may come back to

- **Lesson 11 (standard library tour) has no katas or exercise yet** -
  every other lesson does. If asked to "finish it up" or "make it
  consistent," that's the gap: 1-2 katas (e.g. a small JSON round-trip
  problem, a `strings`/`strconv` parsing kata) plus `solutions/`, in the
  same shape as every other lesson's `katas/`.
- The user originally floated wanting to explore "fundamental standard
  libraries" as a maybe-later item; it got scoped and built as lesson 11
  in the same session it was raised. If they want it deeper (more
  packages - `regexp`, `context` cross-links, `bytes`, `math/rand/v2`,
  etc.), that's an additive extension to 11, not a restructure.
- No lesson renumbering is anticipated to be needed - 11 was deliberately
  appended (not inserted) to avoid another cross-file renumbering pass
  like the one that happened when 02-core-fundamentals was added. If a
  future ask wants stdlib content to appear *earlier* in the path, that
  would require the same kind of renumbering + cross-reference sweep
  02-core-fundamentals's insertion needed - flag that cost to the user
  before doing it, don't just do it silently.
