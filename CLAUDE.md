# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

`sgo` is a small CLI (module `github.com/vukyn/sgo`, urfave/cli/v2) that analyzes and visualizes the structure of a Go project. Given a project path (`--path`/`-p`, default `.`) and an output format (`--output`/`-o`, `text` (default) or `json`), it scans the tree concurrently and reports code metrics, used frameworks, dependency packages, TODO comments, potential secret keys, empty files, project size, and an overall health summary.

It is a **standalone CLI tool, not a platform service** — like `gobuild/`, it has no domains, no DI container, no UI, and no `mprocs`/`hosts` entries. It is a self-contained binary, so the kuery shared-pkg rule does not apply and the local `pkg/` directory stays.

### Flags

- `--path`/`-p` — path to the Go project to analyze (default `.`, env `SGO_PATH`).
- `--output`/`-o` — output format `text|json` (default `text`, env `SGO_OUTPUT`).

## What it reports

For the scanned project, `sgo` collects: detected Go version, used Go frameworks, potential secret keys, TODO comments, empty `.go` and other files, total/comment/empty line counts, direct + indirect dependency packages (from `go.mod`), total project size, total `.go` file count, and warnings. Cache/config folders and files are ignored via an ignore-pattern list. Scanning is concurrent (worker goroutines fanning out over a path channel) with a `schollz/progressbar` progress indicator.

### Score / summary system

Each run ends with a `SummaryStatus` (`status`, `score`, `notes`) computed in `pkg/analyzer/summary.go`. Scoring starts from a base score and deducts points for code-smell signals — large project size (> 100MB), high/minor TODO counts, single/multiple potential secret keys, single/multiple empty files, and warnings. The resulting score maps to a status:

- `PERFECT` — score at or above the perfect threshold.
- `NEED REVIEW` — score at or above the need-review threshold.
- `BAD` — below the need-review threshold.

All thresholds and deduction weights are named constants in `pkg/analyzer/config.go` — tune scoring there, not inline.

## Structure

```
main.go                  # urfave/cli/v2 entrypoint: flags, runs analyzer, prints JSON or text
pkg/analyzer/
  config.go              # scan defaults, ignore patterns, score thresholds + deduction constants
  analyzer.go            # NewAnalyzer + Analyze(): concurrent scan, per-file analysis, result assembly
  helper.go              # formatting + progress-message helpers (uses kuery/query)
  summary.go             # SummaryStatus(): score computation and status mapping
```

## Commands

```bash
make build PRJ=     # go build -o bin/ ./$(PRJ)  (PRJ from .env / overridable on the CLI)
make install        # go install ./$(PRJ)
make tag            # git tag -a v$(VERSION) + push (VERSION from .env)

# CLI usage
sgo -p <path> -o json|text   # analyze <path>, print JSON or text (default text)

go build ./...      # verify
go vet ./...
```

The Makefile sources `.env` for `PRJ`/`VERSION`.

## Dependencies

- `github.com/urfave/cli/v2` — CLI framework / flag parsing.
- `github.com/vukyn/kuery` (`kuery/query` package only) — shared helpers.
- `github.com/schollz/progressbar/v3` — scan progress indicator.

## Conventions

- Keep scoring thresholds and deduction weights as named constants in `pkg/analyzer/config.go`; never hardcode them in `summary.go`.
- `os.ReadFile` over arbitrary scanned paths (gosec G304) is by-design for a directory-analysis CLI — those findings are intentionally left unsuppressed.
- The cosmetic `math/rand` progress-message picker carries a `//nolint:gosec` annotation (no security impact); keep it when touching `helper.go`.
- Bump `Version` in `main.go` when cutting a release; tag via `make tag`.
