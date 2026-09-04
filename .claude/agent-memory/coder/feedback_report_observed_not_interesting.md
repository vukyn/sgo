---
name: feedback-report-observed-not-interesting
description: Report measured results honestly even when they are uniform/boring; never tune thresholds or parameters to manufacture a more interesting spread
metadata:
  type: feedback
---

When asked to run something and report results, report exactly what was observed. Do
not adjust predicates, thresholds, or parameters to make the output distribution look
more varied or more interesting.

**Why:** the user's framing on the worldbit P3 classifier — "if everything lands on one
outcome, that is a real finding about the default parameter set". A uniform result is
data about the system under test; a tuned result is data about the tuner, and it
silently destroys the experiment the user actually wanted to run.

**How to apply:** whenever a task ends in "run X and report Y", especially where the
user states an expectation up front ("I expect most or all to come out STABLE"). Treat
the stated expectation as context, never as a target. If the result disagrees with the
expectation, say so plainly and let the user decide whether the code or the parameters
are wrong. Same rule for benchmark numbers and test-flakiness rates.

Related: [[worldbit-classifier-p3]]
