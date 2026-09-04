---
name: sgo-onboarded
description: sgo onboarded as standalone Go CLI (analyzes Go project structure); pkg/ exception like gobuild; scan baseline of accepted findings
metadata: 
  node_type: memory
  type: project
---

`sgo` (module `github.com/vukyn/sgo`, urfave/cli/v2) onboarded into platform root 2026-06-16. Standalone CLI that scans a Go project and reports metrics/frameworks/TODOs/secrets/size + scored health summary (`text`/`json`). `main.go` + `pkg/analyzer/{config,analyzer,helper,summary}.go`. NOT a service — no domains/DI/UI/mprocs/hosts. Like [[gobuild-preset-system]], the kuery shared-pkg rule does NOT apply; its local `pkg/` stays (self-contained binary). Imports only `kuery/query`.

**Onboarding PR**: vukyn/sgo#1 (branch `chore/platform-onboarding`) — added CLAUDE.md + LICENSE(MIT), `.gitignore`+=`.code-review-graph/`, go 1.24.1→1.26.4, kuery v1.5.0→v1.23.0, x/sys→v0.46.0. Root CLAUDE.md repo list updated (drift fix).

**Scan baseline (accepted — don't re-flag as new)**: gosec G304 `os.ReadFile` at analyzer.go:101 & :180 = by-design (dir-analysis CLI reads user-pointed paths). G404 helper.go:25 + G104 analyzer.go:112 = cosmetic progress-bar (silenced w/ `_ =` and `//nolint:gosec`). No secrets, no reachable code vuln.
