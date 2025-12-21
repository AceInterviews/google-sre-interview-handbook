# 🧾 Interviewer Scorecards: How Google SRE Candidates Are Actually Evaluated

> **“You are not scored on answers.  
You are scored on signals emitted while answering.”**

This document describes the **evaluation dimensions** used internally during Google SRE interviews.

It is *not* a checklist.  
It is *not* a rubric you can memorize.  

It is a **map of attention** — what interviewers are trained to notice while you speak, reason, and react.

Most candidates fail because they optimize for *correctness*.  
Successful candidates optimize for *trust*.

---

## What This Is (and Is Not)

**This is:**
- A high-level view of *how interviewers think*
- A description of *what gets written down*
- A way to understand *why feedback feels vague*

**This is not:**
- A scoring formula
- A guaranteed playbook
- A substitute for practice under pressure

Interviewers are explicitly trained **not** to share this level of detail with candidates.

---

## The 5 Core Scorecard Dimensions

Every Google SRE interview — regardless of round — rolls up into **five evaluation dimensions**.

Not every round tests all five equally.  
But **every hire decision considers all five**.

---

## 1️⃣ Operational Judgment

**What is being evaluated**

Can this person make *safe decisions* under ambiguity?

**Signals interviewers note**

- Do they prioritize mitigation over explanation?
- Do they recognize irreversible actions?
- Do they slow down when stakes increase?

**Common negative signal**

Jumping into debugging while users are still impacted.

**What interviewers write (paraphrased)**

> “Good technical depth, but unclear judgment under pressure.”

---

## 2️⃣ Sequencing & Prioritization

**What is being evaluated**

Does the candidate do the *right things in the right order*?

**Signals interviewers note**

- Clear phase separation (stabilize → investigate → prevent)
- Avoiding parallel changes
- Explicit tradeoff narration

**Common negative signal**

Solving the “interesting” part first instead of the “important” part.

**This dimension quietly decides seniority leveling.**

---

## 3️⃣ Systems & Kernel Intuition

**What is being evaluated**

Does the candidate understand *where failures actually live*?

**Signals interviewers note**

- Awareness of kernel states (I/O wait, D-state, memory reclaim)
- Ability to reason below dashboards
- Comfort operating without logs

**Common negative signal**

Treating metrics as ground truth rather than clues.

This dimension has increased in weight post-2024.

---

## 4️⃣ Communication Under Load

**What is being evaluated**

Can this person be trusted in a real incident?

**Signals interviewers note**

- Calm narration
- Structured explanations
- Explicit uncertainty handling

**Common negative signal**

Sounding confident while skipping reasoning.

Interviewers are trained to penalize *false certainty* more than ignorance.

---

## 5️⃣ Reliability Mindset

**What is being evaluated**

Does this person think in terms of **risk, not features**?

**Signals interviewers note**

- Use of SLOs and error budgets
- Blast radius awareness
- Failure-mode thinking

**Common negative signal**

Optimizing for performance or elegance at the cost of safety.

---

## How Scorecards Are Actually Used

Each interviewer submits:

- A written narrative
- Signal highlights (positive and negative)
- A hire recommendation

The **hiring committee does not re-evaluate your answers**.  
They evaluate **patterns across scorecards**.

One strong negative signal repeated across rounds is usually decisive.

---

## Why Feedback Feels Vague

Candidates often receive feedback like:

> “Strong technical skills, but concerns around execution.”

This maps directly to scorecard dimensions — but those mappings are never shared.

This document exists to make that gap visible.

---

## A Critical Warning

Understanding scorecards improves awareness.  
It does **not** improve performance by itself.

Under interview pressure:
- cognition narrows
- sequencing collapses
- habits dominate

Only **repeated simulation** builds the reflexes these scorecards reward.

---

## Final Thought

Google SRE interviews are not designed to find the smartest engineer.

They are designed to find the engineer who:
- reduces risk
- preserves observability
- communicates clearly
- and can be trusted when information is incomplete

Everything else is secondary.
