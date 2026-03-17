# 🧾 Interviewer Scorecards: How Google SRE Candidates Are Actually Evaluated

> **“You are not scored on your answers.  
> You are scored on the signals you emit while answering.”**

This document describes the **evaluation dimensions** used internally by Hiring Committees (HC) during Google SRE interviews. 

It is *not* a checklist. It is a **map of attention** — what interviewers are trained to notice while you speak, reason, and react under pressure.

Most candidates fail because they optimize for *correctness*. Successful candidates optimize for *trust*.

---

## 🚨 How Scorecards Are Actually Written

Interviewers do not grade you on a scale of 1 to 10 for "knowing Linux." They write a narrative justification for a **Hire** or **No Hire** recommendation, backed by specific signals.

Here is the difference between what you say and what they write:

```diff
- Candidate: "I will check the database CPU and query logs."
- Interviewer Note: "Candidate jumped straight to RCA (Root Cause Analysis). Failed to check user impact or stabilize the system. Weak operational judgment."

+ Candidate: "I'll drain traffic from the failing region to stop user impact, then I'll investigate the database."
+ Interviewer Note: "Candidate demonstrated excellent execution sequencing. Prioritized mitigation and blast-radius containment before debugging. Strong Hire."
```

The Hiring Committee does not re-evaluate your technical answers. They evaluate the **patterns across these scorecards**. One strong negative signal (like ignoring user impact) repeated across two rounds is usually decisive.

---

## 🧠 The 5 Core Scorecard Dimensions

Every Google SRE interview — regardless of the round — rolls up into **five evaluation dimensions**.

### 1️⃣ Operational Judgment
*Can this person make safe decisions under ambiguity?*

*   **Positive Signal:** Prioritizes mitigation over explanation. Recognizes irreversible actions (e.g., dropping a database table vs. failing over).
*   **Negative Signal:** Jumping into debugging while users are still impacted. 
*   **What HC Reads:** *"Good technical depth, but dangerous judgment under pressure."*

### 2️⃣ Sequencing & Prioritization
*Does the candidate do the right things in the right order?*

*   **Positive Signal:** Clear phase separation (Stabilize → Investigate → Prevent). Explicitly narrates trade-offs.
*   **Negative Signal:** Solving the “interesting” part of the system design first, instead of the “important” part (e.g., designing the cache before calculating the network bandwidth).
*   **The Reality:** This dimension quietly decides your seniority leveling (L4 vs L5).

### 3️⃣ Systems & Kernel Intuition
*Does the candidate understand where failures actually live?*

*   **Positive Signal:** Awareness of kernel states (I/O wait, D-state, memory reclaim, CFS quotas). Ability to reason *below* the dashboards.
*   **Negative Signal:** Treating metrics as ground truth rather than clues. Assuming "CPU is fine" means the system is healthy.
*   **The Reality:** This dimension has increased in weight post-2024. Cloud abstractions hide the easy bugs; SREs get the kernel bugs.

### 4️⃣ Communication Under Load
*Can this person be trusted to lead a War Room in a real incident?*

*   **Positive Signal:** Calm narration. Structured explanations. Explicitly stating what they *don't* know.
*   **Negative Signal:** Sounding confident while skipping reasoning. Bluffing command syntax.
*   **The Reality:** Interviewers are trained to penalize *false certainty* much more harshly than ignorance. 

### 5️⃣ Reliability Mindset
*Does this person think in terms of risk, not features?*

*   **Positive Signal:** Natural use of SLOs and error budgets to justify decisions. Constant blast-radius awareness. 
*   **Negative Signal:** Optimizing for performance or algorithmic elegance at the cost of safety and recoverability.

---

## 🧩 Why Feedback Feels Vague

Candidates who are rejected often receive feedback like:
> *"Strong technical skills, but concerns around execution."*

This maps directly to these scorecard dimensions. You solved the LeetCode problem, but your code wasn't safe for production. You designed the system, but your execution sequence was backwards. 

Because interviewers are explicitly trained **not** to share this level of detail with candidates, the feedback feels vague. This document exists to make that gap visible.

---

## 🚀 The Execution Gap: Knowing vs. Performing

Understanding scorecards improves your awareness. It does **not** improve your performance.

Under the pressure of a 45-minute Google interview:
*   Cognition narrows.
*   Execution sequencing collapses.
*   Old habits (like "guessing the bug") dominate.

Candidates rarely fail on one mistake. They fail due to **signal drift** they never notice during the interview.

### How to Train Against the Full Rubric

Strong candidates don't just read the rubric; they practice against it. They use:
*   Interviewer-style prompts with hidden constraints.
*   Real-time constraint changes ("*What if the network partition isolates the cache?*").
*   Partial information scenarios.

Only **repeated simulation** builds the reflexes these scorecards reward.

👉 **[Get The Complete Google SRE Interview Career Launchpad (Gumroad)](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

This file reveals the *shape* of the evaluation.  
**The full system trains you to score 'Exceptional' inside it.**
