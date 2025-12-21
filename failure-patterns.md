# ❌ Failure Patterns in Google SRE Interviews (2025–2026)

> **“Most candidates don’t fail because they lack knowledge.  
> They fail because they apply the right knowledge at the wrong time.”**

This document captures **recurring failure patterns** observed across modern Google SRE interview loops.

These are **not mistakes**.
They are **misalignments between candidate behavior and what the interview is actually simulating**.

If you recognize yourself in more than one section below, that is normal.
Most strong engineers do.

---

## Pattern #1 — Treating the Interview Like a Technical Test

**What the candidate thinks**

> “If I get the correct solution, I’ll pass.”

**What the interview actually is**

A **live operational simulation** under uncertainty.

**How this failure shows up**

- Jumping straight to architecture before clarifying constraints
- Optimizing design before stabilizing impact
- Solving the problem instead of managing the situation

**Example**

Prompt:
> “Latency is high in one region after a deploy.”

Candidate response:
> “I’ll redesign the caching layer to reduce latency.”

Why this fails:
- No mitigation
- No rollback
- No blast-radius reasoning
- No acknowledgment of user impact

**What interviewers expected instead**

> “Before investigating root cause, I want to stop user impact by draining traffic or rolling back.”

**Signal sent**

❌ Feature builder  
✅ Risk manager (expected)

---

## Pattern #2 — Root Cause Obsession During Active Incidents

**What the candidate thinks**

> “Finding the root cause quickly shows seniority.”

**Hidden rubric reality**

Google explicitly scores **Time to Mitigation**, not Time to Root Cause.

**How this failure shows up**

- Diving into logs immediately
- Running multiple commands without stopping impact
- Explaining *why* instead of *what you’ll do right now*

**Example**

Candidate:
> “Let me grep logs to see what error occurred.”

Interviewer (silently thinking):
> “Users are still down.”

**Correct framing**

> “I’m not investigating yet. I’m stabilizing the system first.”

**Why this matters**

An SRE who finds the root cause after 30 minutes but mitigates in 2 is safer than one who finds it in 5 and mitigates in 25.

---

## Pattern #3 — Architecture Before Arithmetic (The NALS Trap)

**What the candidate thinks**

> “NALS is system design with more constraints.”

**What NALS actually tests**

Whether you can detect **physical impossibility**.

**How this failure shows up**

- Proposing replication without calculating bandwidth
- Suggesting redundancy without failure-rate math
- Ignoring storage, power, rack, or time constraints

**Classic fail**

> “We’ll replicate 5PB of data across regions.”

Without calculating:
- transfer time
- link capacity
- recovery windows

**What passes**

> “Before proposing architecture, I want to validate whether the constraints make the requirement achievable.”

**Signal sent**

❌ Diagrammer  
✅ Custodian of scarce resources

---

## Pattern #4 — Metric Worship (Post-2024 Anti-Pattern)

**What the candidate thinks**

> “Metrics tell me what’s wrong.”

**Modern reality**

Metrics are often **lagging indicators**.

Many failures now happen:
- in kernel space
- in I/O wait
- in networking queues
- between dashboards

**How this failure shows up**

- “CPU looks fine” → stops investigating
- Over-reliance on dashboards
- No kernel-level hypotheses

**What strong candidates do**

They ask:
- Are processes stuck in D-state?
- Is there memory reclaim pressure?
- Are file descriptors exhausted?
- Is conntrack saturated?

**Interview ceiling**

Candidates who never leave dashboards are rarely rated above L4.

---

## Pattern #5 — The Scattershot Debugger

**What the candidate thinks**

> “Running more commands shows thoroughness.”

**What it actually signals**

Lack of hypothesis discipline.

**How this failure shows up**

- Running `top`, then `lsof`, then `netstat`, then logs
- Changing multiple variables at once
- Restarting services without explanation

**Why this is dangerous**

It destroys observability and masks causality.

**What interviewers expect**

One change.
One hypothesis.
One observation.

Over and over.

---

## Pattern #6 — False Certainty Penalty

**What the candidate thinks**

> “I need to sound confident.”

**What Google values instead**

Epistemic humility.

**How this failure shows up**

- Guessing commands
- Bluffing kernel behavior
- Asserting without data

**What scores higher**

> “I don’t know the exact command, but I know I need to inspect TCP retransmissions.”

**Why**

SREs are trusted because they know **what they don’t know**.

---

## Pattern #7 — Coding Like a Competitive Programmer

**What the candidate thinks**

> “This is LeetCode Easy/Medium.”

**What the coding round actually tests**

Automation safety under production constraints.

**How this failure shows up**

- Loading full datasets into memory
- Writing dense one-liners
- Ignoring error handling
- No narration of trade-offs

**What passes**

- Streaming input
- Defensive checks
- Clear variable naming
- Calm explanation

**Key insight**

Readable code > clever code.

Always.

---

## Pattern #8 — Talking Like a Developer, Not an SRE

**Developer language**

- “Fix the bug”
- “Optimize performance”
- “Improve code quality”

**SRE language**

- “Reduce blast radius”
- “Protect the SLO”
- “Fail safely”
- “Buy time”

**Interviewers listen for this shift**

It signals identity change.

---

## The Real Divide

Most failed candidates:
- know the tools
- know the theory
- know the answers

They fail because they:
- act too early
- optimize too soon
- explain too much
- mitigate too late

Passing the Google SRE loop is not about brilliance.

It is about **sequence**.

---

## A Note on Preparation

Reading failure patterns creates awareness.

It does **not** create reflexes.

Execution under pressure requires:
- practice with interruption
- simulated ambiguity
- deliberate recovery from wrong first steps

That gap is intentional.

This repository documents the *map*.
Walking it is a separate skill.
