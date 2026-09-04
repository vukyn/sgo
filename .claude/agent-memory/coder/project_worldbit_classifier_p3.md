---
name: worldbit-classifier-p3
description: worldbit P3 (internal/stats classifier) delivered on feat/p2-agents; default params make nearly every seed STABLE, so P4 gains a parameter sweep
metadata:
  type: project
---

worldbit P3 (`internal/stats`: intcmp / window / classify / recorder) was implemented
on branch `feat/p2-agents` on 2026-07-31, uncommitted, handed to the committer agent.
Seeds 1-8 at 12 000 ticks classify as 7 x STABLE + 1 x OSCILLATING (seed 5) under the
default parameter set.

**Why:** the default parameters produce near-identical runs (boom to ~2100-3000 around
tick 760-820, settling to ~700-820, no extinctions), so seed variation alone cannot
exercise the six outcomes. The user's response was NOT to retune the classifier but to
extend **P4 (batch runner) with a parameter sweep** — sweeping config, not just seeds,
is how the interesting regions get found. The plan doc still files a sweep under a
hypothetical P7; the P4 extension supersedes that.

**How to apply:** when P4 lands, expect `--set`-style config-grid sweeping in scope
alongside the seed worker pool, and expect the runner to stop a run early on a terminal
outcome (STABLE included — it is terminal, so a "STABLE" record means three consecutive
stable windows were reached, not that the run stayed stable to the tick limit; seeds
1/2/6/7 resolve at tick 4800, seed 3 at 7200, seeds 4/8 at 9600). Do not read a STABLE
record as a statement about the whole run.

Related: [[feedback-report-observed-not-interesting]]
