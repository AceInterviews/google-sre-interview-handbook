# ⚡ The Google SRE Interview Handbook (2026+ Edition)

> **"The definitive open-source playbook for the modern Google Site Reliability Engineer (SRE). Moving beyond LeetCode to NALS, Linux Internals, and Reliability Architecture."**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

## 🚨 The Problem
Most Google SRE interview guides are stuck in 2018. They focus on generic algorithms and high-level system design.

But the modern Google SRE interview loop has evolved. It now tests for **Operational Maturity**:
1.  **NALS (Non-Abstract Large System Design):** Can you diagnose a complex system you didn't build?
2.  **Linux Internals:** Can you debug a kernel panic or CPU throttle without guessing?
3.  **Reliability Architecture:** Can you design for 99.999% reliability using error budgets and SLOs?

This repository contains the **core mental models, checklists, and cheat sheets** required to pass these rounds.

---

## 📚 Table of Contents

*   **[The NALS Diagnostic Flowchart for Google SRE Interviews](nals-playbook.md)** - How to stabilize and debug large-scale outages.
*   **[The Linux Internals Cheat Sheet for Google SRE Interviews](linux-internals.md)** - The 20 commands that solve 80% of incidents.
*   **[Coding Patterns for Google SRE Interviews](coding-patterns.md)** - Python/Go patterns for concurrency, retries, and rate limiting.
*   **[Behavioral: The Google SRE-STAR(M) Method](behavioral-guide.md)** - How to answer "Tell me about a time..." with metrics.
*   **[Negotiation Pocket Card for Google SRE Interviews](negotiation-tips.md)** - Scripts to maximize your offer.

---

## 🧠 The Google SRE Mastery Pyramid

To pass the Staff/Senior rounds, you need to move beyond "using tools" to "understanding systems."

| Layer | Focus | Key Tools |
| :--- | :--- | :--- |
| **Culture** | Blamelessness, Postmortems, Psychological Safety | RCAs, Error Budgets |
| **Incident Eng** | "Observe → Localize → Fix → Prevent" | Live Debugging |
| **Observability** | eBPF, PSI, Kernel Telemetry | Prometheus, bcc-tools |
| **Linux Tools** | `strace`, `perf`, `tcpdump`, `lsof` | Reading queues/resources |
| **Kernel** | Scheduler, Memory, Network Stack | Syscalls, IRQs, OOM |

---

# 📘 Interview-Grade Linux Command Playbook

### How Google SREs Use Basic Commands Under Pressure

> **This is not a Linux commands list.**
> This is how Google evaluates *judgment, signal selection, and reasoning* through command usage.

Most candidates know Linux commands.
Very few know **which command to reach for first — and why**.

That difference decides interviews.


## 🎯 Why This Module Exists

In Google SRE interviews:

* You are **not evaluated on command memorization**
* You *are* evaluated on:

  * prioritization
  * signal selection
  * reasoning under uncertainty
  * calm narration of thought process

Interviewers listen closely to:

> *“Why did you choose that command?”*

This playbook teaches **command → intent → signal** mapping.


## 🧠 The Mental Model Interviewers Expect

Before touching the keyboard, strong SREs silently ask:

1. Is this a **live system or historical issue**?
2. Am I looking for **symptoms or root cause**?
3. Is the problem **CPU, memory, I/O, network, or application-level**?
4. What is the **lowest-cost, highest-signal command** to start with?

This module trains that reflex.



## 🔥 Section 1 — Log Inspection (The First 5 Minutes Matter)

### `tail -f` — Live Signal, Not Noise

**When to use**

* Incident is ongoing
* You want to correlate behavior with time

**Interview narration**

> “I’ll start with `tail -f` to observe real-time behavior before making assumptions.”

**Example**

```bash
tail -f /var/log/nginx/error.log
```

**Signal interviewers look for**

* You didn’t start grepping blindly
* You value *temporal correlation*



### `tail -n` — Context Before Panic

**When to use**

* Incident already occurred
* You want *recent history*, not the full log

```bash
tail -n 200 /var/log/app.log
```

**What this shows**

* You understand **blast radius**
* You don’t overload yourself with data



### `grep` + `tail` — Precision Over Volume

Bad candidates:

```bash
grep error /var/log/app.log
```

Strong candidates:

```bash
tail -n 5000 /var/log/app.log | grep -i timeout
```

**Why this matters**

* Shows scoped thinking
* Shows intent to reduce false positives



## ⚙️ Section 2 — Process & Resource Awareness

### `ps aux` — Snapshot Thinking

**When to use**

* Need a fast view of resource-heavy processes

```bash
ps aux --sort=-%cpu | head
```

**Interview signal**

* You understand *static snapshots vs dynamic metrics*



### `top` vs `htop` — Conscious Tool Choice

**Good explanation**

> “I’ll use `top` first for a quick system-wide view before drilling deeper.”

Interviewers care **why**, not **which**.



## 🧵 Section 3 — File Descriptors & Hidden Killers

### `lsof` — The Silent Outage Detector

**Use cases**

* Port binding failures
* File descriptor exhaustion
* Stuck deleted files consuming disk

```bash
lsof -i :8080
```

**Narration**

> “This helps confirm whether the service is actually listening or blocked by another process.”



## 💾 Section 4 — Disk & I/O (Where Seniors Stand Out)

### `df -h` vs `du -sh`

Strong candidates **never confuse these**.

```bash
df -h
du -sh /var/*
```

**Key insight**

* `df` = filesystem view
* `du` = directory view

Interviewers *love* this distinction.



## 🧠 Section 5 — Memory & Kernel Signals

### `free -m`

```bash
free -m
```

**Strong explanation**

> “I’m checking memory pressure and reclaim behavior before assuming a leak.”

Bonus points if candidate mentions:

* page cache
* swap behavior



## 🌐 Section 6 — Network Sanity Checks

### `netstat` / `ss`

```bash
ss -lntp
```

**Use when**

* Service unreachable
* Connection exhaustion suspected

**Narration**

> “This confirms whether the service is bound and accepting connections.”



## 🔁 Section 7 — Putting It Together (Interview Scenario)

**Question**

> “The site is slow. CPU is normal. What do you do?”

**Strong answer flow**

1. `tail -n` logs for recent anomalies
2. `ss` to check connection backlog
3. `iostat` / disk wait (if available)
4. Only then form hypotheses

This shows:
✔ restraint
✔ prioritization
✔ system thinking



## 🚫 Common Interview Red Flags

Interviewers immediately notice when candidates:

* start grepping entire logs
* run commands without explaining intent
* jump tools without hypothesis
* over-optimize too early

This module trains you *out* of those habits.


## ⭐ Why This Section Improves Coding & Scripting Rounds

Even in **coding interviews**, Google evaluates:

* how you reason about input size
* streaming vs batch thinking
* observability mindset

Candidates who think like SREs:

* choose iterators
* avoid loading everything into memory
* explain tradeoffs clearly

This command mindset **directly transfers** to code.


## 🎯 Final Takeaway

Google doesn’t hire people who know Linux commands.

Google hires people who:

* know **which signal matters first**
* stay calm under ambiguity
* explain their thinking clearly
* choose tools intentionally

This playbook teaches exactly that.




---

## 🚀 Want the Complete System?

This repository covers the **foundational frameworks**.

If you want the **complete end-to-end preparation system**—including practice workbooks, mock simulations, and deep-dive scenarios—check out the full bundle:

👉 **[The Complete Google SRE Career Launchpad (Gumroad)](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

**What's included in the full bundle:**
*   📘 **20+ Production Scenarios:** Deep dives into Kernel Panics, BGP Leaks, and connection storms.
*   🐍 **Coding Workbooks (Python & Go):** 70+ practice problems focusing on concurrency, automation, and safety.
*   💼 **The Negotiation Playbook:** Word-for-word scripts to increase your offer.
*   📅 **The 30-Day Prep Schedule:** A day-by-day roadmap to interview readiness.

---

## 🤝 Contributing
Contributions are welcome! Please open an issue or PR if you have additional commands, patterns, or interview insights to share.

## 📄 License
MIT License. Free to use and share.
