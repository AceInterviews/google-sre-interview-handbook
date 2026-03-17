# ✅ Counter-Patterns: What Passing Google SRE Candidates Do Differently

> **“Strong candidates don’t do different things.  
> They do the same things — but in a different order and with a different reflex.”**

This document describes the **behavioral counter-patterns** consistently observed in *passing* Google SRE interviews (L4, L5, L6).

These are not tricks. They are **habits of sequencing, narration, and restraint**.

Most candidates know *what* to do. Passing candidates know **when** to do it.

---

## 🛑 Counter-Pattern #1 — Stabilize First, Always

```diff
- Failing Instinct (The Detective)
- "Let me understand the problem fully before acting. I'll check the logs."

+ Passing Behavior (The Commander)
+ "Let me reduce user impact before understanding the problem fully. I'll drain traffic."
```

**What passing candidates do:**
*   Explicitly separate *mitigation* from *investigation*.
*   Treat every live issue as time-sensitive.
*   Buy time before spending thought.

**The Golden Phrase:** *"Before diagnosing the root cause, my priority is to stop the bleeding."*

---

## 🛑 Counter-Pattern #2 — Declare Constraints Before Solutions

```diff
- Failing Instinct (The Builder)
- "I’ll design a robust, multi-region active-active architecture."

+ Passing Behavior (The Architect)
+ "I want to confirm the constraints (RTO/RPO, budget, bandwidth) before proposing anything."
```

**What passing candidates do:**
*   Ask about: scale, failure rates, recovery windows, and economic budgets.
*   Delay design until constraints are mathematically explicit.
*   Are comfortable proving a requirement is *impossible* (e.g., 5PB over 10Gbps).

---

## 🛑 Counter-Pattern #3 — One Hypothesis at a Time

```diff
- Failing Instinct (The Guesser)
- "Let me check logs, metrics, network, and CPU to see what looks weird."

+ Passing Behavior (The Scientist)
+ "My hypothesis is a network bottleneck. I will check the TCP retransmit queue to validate."
```

**Why this matters:**
Interviewers score **debugging discipline**, not tool breadth. Restraint > Speed.
*   State a hypothesis out loud.
*   Run *exactly one* validating action.
*   Observe and Adjust.

---

## 🛑 Counter-Pattern #4 — Narrate Intent, Not Commands

```diff
- Failing Instinct (The Terminal Jockey)
- "I’ll run `top`, then `lsof`, then `netstat`."

+ Passing Behavior (The Diagnostician)
+ "I want to understand whether this is CPU-bound, I/O-bound, or blocked on the kernel."
```

**Interviewers evaluate:**
*   Clarity of reasoning.
*   Prioritization.
*   Calmness under ambiguity.
*   *Not* command memorization. Commands are implementation details; intent is the signal.

---

## 🛑 Counter-Pattern #5 — Prefer Leading Indicators Over Dashboards

```diff
- Failing Instinct (The Dashboard Watcher)
- "CPU is fine, so it’s not compute-related."

+ Passing Behavior (The Kernel Whisperer)
+ "CPU metrics can hide stalls. I want to check wait states and run-queue latency."
```

**What passing candidates look for:**
*   `D-state` processes (Uninterruptible Sleep).
*   Memory pressure (Page Cache thrashing vs. RSS).
*   File descriptor exhaustion.
*   Connection backlogs (SYN drops).

---

## 🛑 Counter-Pattern #6 — Say “I Don’t Know” Precisely

```diff
- Failing Instinct (The Bluffer)
- *Starts guessing at kernel flags or making up API endpoints.*

+ Passing Behavior (The Epistemically Humble)
+ "I don’t know the exact command syntax, but I know the signal I need is the TCP window size."
```

**Why this scores higher:**
*   Shows epistemic discipline (you know what you don't know).
*   Preserves trust (crucial for on-call teams).
*   Aligns with Google's blameless, ego-less culture.

---

## 🛑 Counter-Pattern #7 — Code Defensively, Not Cleverly

```diff
- Failing Instinct (The LeetCoder)
- "I’ll write the shortest, most algorithmically elegant solution using `readlines()`."

+ Passing Behavior (The SRE)
+ "I’ll write streaming code an on-call engineer can trust to not OOM-crash at 3 AM."
```

**What passing code looks like:**
*   Streaming input (O(1) memory).
*   Explicit error handling (No silent failures).
*   Readable control flow over clever one-liners.

---

## 🛑 Counter-Pattern #8 — Use SRE Language Deliberately

Passing candidates don't just use vocabulary; they use **identity signaling**. They naturally frame answers using:

*   [ ] "Blast radius"
*   [ ] "Fail safely / Fail open"
*   [ ] "Error budget"
*   [ ] "Graceful degradation"
*   [ ] "Toil reduction"

Interviewers notice this immediately. It proves you are already one of them.

---

## 🛑 Counter-Pattern #9 — Treat Silence as Thinking Time

```diff
- Failing Instinct (The Panicker)
- *Filling silence with rambling explanations and backtracking.*

+ Passing Behavior (The Commander)
+ *Pausing. Thinking. Then speaking with deliberate intent.*
```
Silence signals composure, not confusion.

---

## 🎯 The Core Shift

Passing candidates stop asking:
> *"How do I solve this problem?"*

They start asking:
> *"What is the safest next move right now?"*

That shift changes everything.

---

## 🚀 The "Knowing vs. Doing" Gap

Reading what strong candidates do differently creates clarity. **Executing those behaviors under pressure is harder.**

In real Google SRE interviews:
- These counter-patterns must appear **consistently**.
- They must surface **without prompting**.
- They must hold even when the interviewer introduces mid-scenario twists.

Most candidates can describe the right behavior. Very few can **default to it automatically**. That gap—between recognition and reflex—is where interviews are decided.

If you want to train these reflexes, I built a **simulation-based preparation system**. It forces you to practice execution sequencing, partial-information debugging, and recovery from wrong decisions.

👉 **[Get The Complete Google SRE Interview Career Launchpad (Gumroad)](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

This repository sharpens your awareness. The full system builds your interview-grade instinct.
