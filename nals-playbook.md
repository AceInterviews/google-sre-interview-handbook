# 📉 The NALS Diagnostic Flowchart  
**Non-Abstract Large System Design (NALSD) — How Google Actually Evaluates You**

> **"NALSD is not system design. It is operational physics under constraints."**

Most candidates fail this round because they treat it like a generic system design interview  
(e.g., “Design Twitter”, “Add a cache”, “Add replication”).

**NALSD is different.**

You are given an **existing system** that is:
- already in production
- partially broken, or
- facing a scale / migration / reliability constraint

The interviewer is evaluating whether you can **reason like a production SRE**, not an architect.

---

## 🚨 The Core NALSD Failure Mode

Most candidates:
- start drawing boxes too early
- assume software can solve physical constraints
- ignore feasibility math
- design before stabilizing

Senior candidates fail here *more often* than juniors.

---

## 1️⃣ The Correct NALSD Execution Order (Non-Negotiable)

> **Do not draw anything until steps 1–4 are complete.**

### 1. USER GOAL (Not the System)
What is the *human* actually trying to do?

Examples:
- ❌ “Write to Blobstore”
- ✅ “Upload a photo and see it appear within 1 second”

Interview signal: You think in outcomes, not components.

---

### 2. CONSTRAINTS (Before Architecture)
What is *fixed* and cannot be hand-waved?

Examples:
- budget limits
- hardware counts
- failure rates
- latency bounds
- regulatory or data locality constraints

> **If a requirement violates physics or math, you must say so.**

This is a **pass signal**, not a failure.

---

### 3. SLIs / SLOs (Before Design)
Define how success is measured.

Examples:
- Availability: 99.95%
- Latency: p99 < 200ms
- Error rate: < 0.1%

Interviewers listen for:
> “Before designing, I want to agree on how we’ll measure success.”

---

### 4. FEASIBILITY MATH (The Silent Filter)
This is where many strong candidates fail.

Examples:
- HDD annual failure rate vs replication factor
- Network bandwidth vs data replication size
- Recovery time vs SLO window

> **If the math doesn’t work, stop. Say it out loud.**

Continuing to design after infeasible math is an automatic down-level.

---

### 5. DATA FLOW (Only Now)
Trace the request lifecycle:
- entry point
- dependencies
- storage
- fan-out
- fan-in
- failure points

No boxes yet — just flow.

---

### 6. RELIABILITY MECHANISMS
Only after feasibility is proven:
- retries (bounded)
- backoff
- circuit breakers
- dead-letter queues
- graceful degradation

---

### 7. SCALING STRATEGY
Think in *failure domains*, not throughput:
- sharding boundaries
- regional isolation
- blast radius containment

---

### 8. OBSERVABILITY
What tells you this system is failing *before users complain*?

Signals > dashboards.

Strong candidates mention:
- burn-rate alerts
- saturation signals
- queue depth
- tail latency

---

### 9. FAILURE & MITIGATION FIRST
Given a failure:
1. Clarify scope
2. Mitigate immediately
3. Restore safety
4. Investigate later
5. Prevent recurrence

> Root cause comes **after** stability.

---

### 10. TRADE-OFFS (Explicitly)
What did you sacrifice and why?

Examples:
- consistency for availability
- freshness for safety
- cost for resilience

Interviewers expect you to **name losses**, not hide them.

---

## 2️⃣ The “5-S Rule” (Updated for NALSD)

| S | What Interviewers Listen For |
|---|---|
| **Scope** | Are you narrowing the problem correctly? |
| **Scale** | Are you quantifying load, not guessing? |
| **SLIs** | Do you anchor decisions in metrics? |
| **Scarcity** | Do you respect physical limits? |
| **Safety** | Do you fail visibly and controllably? |

---

## 3️⃣ Power Phrases That Signal Senior SRE Thinking

Use sparingly. Overuse sounds scripted.

- “Before designing, I want to validate feasibility.”
- “If the math doesn’t work, we should stop here.”
- “Let’s stabilize impact before improving architecture.”
- “This requirement violates our failure budget.”
- “I’m intentionally not optimizing yet.”

---

## 4️⃣ The NALSD Stabilization Rule

In **any broken-system prompt**:

1. Clarify impact
2. Reduce blast radius
3. Restore safe state
4. Investigate root cause
5. Design prevention

Skipping step 2 is a red flag.

---

## 🚫 Common NALSD Red Flags (Immediate Down-Level)

Interviewers immediately notice when candidates:
- draw architecture before defining SLOs
- ignore feasibility math
- assume “more replicas” solves everything
- debug before mitigating
- optimize before stabilizing

---

## 🎯 Why This Matters

NALSD is where Google filters for:
- Staff+ judgment
- risk management instincts
- production realism
- decision sequencing

It is **not** about creativity.
It is about restraint.

---

## 🚀 Want to Practice This the Right Way?

This file teaches **what** interviewers expect.

The full Google SRE system trains:
- real NALSD scenarios
- math-first feasibility drills
- mitigation-before-design instincts
- interviewer scorecards
- execution sequencing under time pressure

👉 **The Complete Google SRE Interview Career Launchpad**  
https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer

Reading changes awareness.  
Simulation changes outcomes.
