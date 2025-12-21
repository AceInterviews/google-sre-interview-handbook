# ✅ Counter-Patterns: What Passing Google SRE Candidates Do Differently (2026+)

> **“Strong candidates don’t do different things.  
They do the same things — but in a different order.”**

This document describes **behavioral counter-patterns** consistently observed in *passing* Google SRE interviews.

These are not tricks.
They are **habits of sequencing, narration, and restraint**.

Most candidates know *what* to do.
Passing candidates know **when** to do it.

---

## Counter-Pattern #1 — Stabilize First, Always

**Failing instinct**

> “Let me understand the problem fully before acting.”

**Passing behavior**

> “Let me reduce user impact before understanding the problem fully.”

**What passing candidates do**

- Explicitly separate *mitigation* from *investigation*
- Treat every live issue as time-sensitive
- Buy time before spending thought

**Typical phrasing**

> “Before diagnosing root cause, I want to stop the bleeding.”

This sentence alone signals seniority.

---

## Counter-Pattern #2 — Declare Constraints Before Solutions

**Failing instinct**

> “I’ll design a robust architecture.”

**Passing behavior**

> “I want to confirm the constraints before proposing anything.”

**What passing candidates do**

- Ask about:
  - scale
  - failure rates
  - recovery windows
  - budgets
- Delay design until constraints are explicit

**Key insight**

In NALS, architecture is often *downstream* of math.

Passing candidates are comfortable proving a requirement is impossible.

---

## Counter-Pattern #3 — One Hypothesis at a Time

**Failing instinct**

> “Let me check logs, metrics, network, and CPU.”

**Passing behavior**

> “I have one hypothesis. I’ll test it.”

**What passing candidates do**

- State a hypothesis out loud
- Run exactly one validating action
- Observe
- Adjust

**Why this matters**

Interviewers score **debugging discipline**, not tool breadth.

Restraint > speed.

---

## Counter-Pattern #4 — Narrate Intent, Not Commands

**Failing instinct**

> “I’ll run `top`, then `lsof`, then `netstat`.”

**Passing behavior**

> “I want to understand whether this is CPU-bound, I/O-bound, or blocked on the kernel.”

Commands are implementation details.
Intent is the signal.

**Interviewers evaluate**

- clarity of reasoning
- prioritization
- calmness under ambiguity

Not command memorization.

---

## Counter-Pattern #5 — Prefer Leading Indicators Over Dashboards

**Failing instinct**

> “CPU is fine, so it’s not compute-related.”

**Passing behavior**

> “CPU metrics can hide stalls. I want to check wait states.”

**What passing candidates do**

- Question surface-level metrics
- Look for:
  - D-state processes
  - memory pressure
  - file descriptor exhaustion
  - connection backlog

**Modern reality**

Post-2024 incidents often occur *between* dashboards.

---

## Counter-Pattern #6 — Say “I Don’t Know” Precisely

**Failing instinct**

> Bluffing commands or kernel behavior.

**Passing behavior**

> “I don’t know the exact command, but I know the signal I need.”

**Why this scores higher**

- Shows epistemic discipline
- Preserves trust
- Aligns with blameless culture

Google hires people who manage uncertainty, not hide it.

---

## Counter-Pattern #7 — Code Defensively, Not Cleverly

**Failing instinct**

> “I’ll write the shortest solution.”

**Passing behavior**

> “I’ll write code an on-call engineer can trust at 3 AM.”

**What passing code looks like**

- streaming input
- explicit error handling
- readable control flow
- clear variable names

**Key signal**

Correctness under constraints > algorithmic elegance.

---

## Counter-Pattern #8 — Use SRE Language Deliberately

Passing candidates naturally use phrases like:

- “blast radius”
- “fail safely”
- “rollback”
- “error budget”
- “guardrails”
- “buying time”

This is not vocabulary.
It is **identity signaling**.

Interviewers notice immediately.

---

## Counter-Pattern #9 — Treat Silence as Thinking Time

**Failing instinct**

> Filling silence with explanations.

**Passing behavior**

- Pausing
- Thinking
- Then speaking with intent

Silence signals composure, not confusion.

---

## The Core Shift

Passing candidates stop asking:

> “How do I solve this problem?”

They start asking:

> “What is the safest next move right now?”

That shift changes everything.

---

## Final Note

Knowing these counter-patterns does not guarantee success.

They must become **reflexes**:
- under time pressure
- with partial information
- while being interrupted

This repository makes the patterns explicit.

Execution is a separate skill.
