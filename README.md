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
