# 🚀 Getting Started: How to Use This Repository

> **"If you're reading random guides, you're already behind.**  
> **Google SRE interviews reward *how you think*, not just what you know."**

Welcome to the **Google SRE Interview Handbook (2026 Edition)**. 

This repository is not a "link dump" or a list of LeetCode questions. It is a **mental model training system** designed to upgrade you from a standard Software Engineer to a Google-caliber Reliability Architect.

Before you click into the folders, you need to understand how to consume this information.

---

## 🛑 The Mental Shift (Read This First)

Most candidates approach SRE interviews the wrong way. If you try to study for this like a standard FAANG Software Engineering loop, you will fail.

```diff
- The Failing Approach (The Developer)
- 1. Grind 200 LeetCode Mediums.
- 2. Memorize a list of 50 Linux commands.
- 3. Draw a global load balancer for every system design question.

+ The Passing Approach (The Reliability Architect)
+ 1. Practice writing safe, streaming, bounded-memory automation.
+ 2. Learn to extract kernel-level signals (I/O wait, D-state, CFS quotas).
+ 3. Do the physical math (Bandwidth-Delay Product) before drawing any architecture.
```

This repository is designed to teach you the **green path**.

---

## 🗺️ How This Repository is Structured

To get the most out of this handbook, do not read it randomly. We have organized the knowledge into specific phases. 

### 🌟 The "Shock and Awe" Root Files (Start Here)
If you only have 15 minutes today, read these files located in the root directory. They expose the "meta-game" of the interview:
*   👉 **[Real Interview Patterns](REAL-INTERVIEW-PATTERNS_2.md):** How interviewers use twists and silence to break you.
*   👉 **[The Mock Interview Transcript](mock-interview-transcript.md):** A fly-on-the-wall look at how a candidate is graded in real-time.
*   👉 **[Interviewer Scorecards](interviewer-scorecards.md):** The 5 hidden dimensions you are actually graded on.
*   👉 **[SRE vs. SWE System Design](sre-vs-swe-design.md):** The exact difference between an L4 and an L6 architecture.

### 📁 Folder 01: The SRE Mindset
*   **What's inside:** `execution-sequencing.md`, `failure-patterns.md`, `counter-patterns.md`, `REAL-INTERVIEW-PATTERNS.md`
*   **How to use it:** Read these to understand *why* doing the right thing in the wrong order causes you to fail.

### 📁 Folder 02: NALSD & System Design
*   **What's inside:** `nals-playbook.md`, `nalsd-math-traps.md`
*   **How to use it:** Use the 5-S Rule and learn the physical math (IOPS, Network Bandwidth) that catches 90% of candidates off guard.

### 📁 Folder 03: Linux & Troubleshooting
*   **What's inside:** `linux-internals.md` (Cheat Sheet), plus **4 Real-World Incident Playbooks** (Kernel Panics, BGP Leaks, TLS Expiry, Disk Pressure).
*   **How to use it:** Stop memorizing flags. Learn how to map a symptom to a subsystem to a root cause.

### 📁 Folder 04: Coding & Automation
*   **What's inside:** `coding-patterns.md`, plus real, heavily-commented production code (`token_bucket.go`, `safe_log_streamer.py`).
*   **How to use it:** Learn why streaming, contexts, and exponential backoff with jitter are mandatory for passing.

### 📁 Folder 05: Behavioral & Negotiation
*   **What's inside:** `behavioral-guide.md`, `negotiation-tips.md`.
*   **How to use it:** Learn the SRE-STAR(M) method to prove your impact, and use the exact scripts to maximize your final offer.

---

## ⚠️ The Execution Gap: Knowing vs. Doing

This repository gives you the **Map**. 

It provides the mental models, the failure patterns, and the execution frameworks. But reading a framework does not build a reflex. 

Understanding that you need to "stabilize first" is easy. Actually remembering to do it when an interviewer tells you *"a backhoe just cut the primary trans-Atlantic fiber line and 15% of checkouts are failing"* is incredibly hard.

**You need simulation.**

If you want the **Complete End-to-End Preparation System**—the engine that actually builds your reflexes—check out the full bundle:

👉 **[The Complete Google SRE Career Launchpad](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

**What you get in the full system:**
*   📘 **20+ Deep-Dive Scenarios:** Full walkthroughs of NALS and Troubleshooting simulations.
*   🐍 **70+ Coding Workbooks (Python & Go):** Practice problems focusing on concurrency, automation, and production safety.
*   💼 **The Offer Maximizer:** Word-for-word negotiation scripts that reflect real compensation committee logic.
*   📅 **The 30-Day Prep Schedule:** A structured, day-by-day roadmap to ensure you are ready.

Use the open-source repo to build your awareness.  
**Use the full bundle to secure the offer.**
