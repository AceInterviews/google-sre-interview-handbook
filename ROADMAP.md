
# 🗺️ The Google SRE Interview Roadmap (2026+ Edition)

> **"Mastering the Google SRE loop is not a matter of 'grinding.' It is a matter of sequencing."**

This roadmap provides a structured, 4-phase learning path to navigate this repository and prepare for a **Senior (L5) or Staff (L6) SRE offer.**

---

## ⏱️ Timeline Overview
*   **DevOps/SRE Background:** 2 Weeks (Focus on NALSD & Sequencing)
*   **SWE/Backend Background:** 4 Weeks (Focus on Linux Internals & Mindset)

---

## 🟢 Phase 1: Deconstructing the SRE Mindset (Days 1–5)
**Goal:** Stop thinking like a feature developer. Learn to prioritize risk and mitigation.

1.  **[Execution Sequencing](01-SRE-Mindset/execution-sequencing.md)**: Study the "Mitigate First" priority. This is the #1 reason L5+ candidates fail.
2.  **[Failure Patterns](01-SRE-Mindset/failure-patterns.md)**: Identify the habits you need to unlearn (e.g., Root Cause obsession).
3.  **[Counter-Patterns](01-SRE-Mindset/counter-patterns.md)**: Learn how to narrate your intent during a crisis.
4.  **[Interviewer Scorecards](interviewer-scorecards.md)**: Understand the 5 dimensions you are actually being graded on.

---

## 🔵 Phase 2: NALSD & Reliability Architecture (Days 6–15)
**Goal:** Master the "Physics of Scarcity." Move from drawing boxes to calculating limits.

1.  **[SRE vs. SWE Design](sre-vs-swe-design.md)**: Visualize the gap between a "functional" design and a "hardened" design.
2.  **[The NALSD Playbook](02-NALSD-System-Design/nals-playbook.md)**: Learn the 8-step framework for existing, broken systems.
3.  **[Math Traps](02-NALSD-System-Design/nalsd-math-traps.md)**: Practice calculating Bandwidth, RTO, and IOPS on the fly.
4.  **[Mock Transcript](mock-interview-transcript.md)**: Read how an L6 candidate handles a Black Friday latency spike.

---

## 🔴 Phase 3: Linux Internals & Troubleshooting (Days 16–25)
**Goal:** Develop kernel-level intuition. Learn to see through the dashboards.

1.  **[Linux Internals Cheat Sheet](03-Linux-Troubleshooting/linux-internals.md)**: Master the "Log Surgeon" and "System Doctor" toolkits.
2.  **[Incident Playbooks](03-Linux-Troubleshooting/)**: Walk through the "Standard Failure Library":
    *   [Kernel Panic & D-State](03-Linux-Troubleshooting/incident-playbook-kernel-panic.md)
    *   [BGP Route Leaks](03-Linux-Troubleshooting/incident-playbook-bgp-leak.md)
    *   [Disk I/O Saturation](03-Linux-Troubleshooting/incident-playbook-disk-pressure.md)
    *   [TLS Expiry Cascades](03-Linux-Troubleshooting/incident-playbook-tls-expiry.md)

---

## 🟡 Phase 4: Coding & Behavioral Signals (Days 26–30)
**Goal:** Finalize your "Identity Signal." Practice safe code and data-backed stories.

1.  **[Coding Patterns](04-Coding-Automation/coding-patterns.md)**: Learn to write streaming, concurrent code with timeouts.
2.  **[Reference Implementations](04-Coding-Automation/)**: Study the [Concurrent Health Checker](04-Coding-Automation/concurrent_health_checker.go) and [Safe Log Streamer](04-Coding-Automation/safe_log_streamer.py).
3.  **[The SRE-STAR(M) Method](05-Behavioral-Negotiation/behavioral-guide.md)**: Rebuild your career stories around **Mitigation** and **Metrics**.
4.  **[Negotiation Pocket Card](05-Behavioral-Negotiation/negotiation-tips.md)**: Prepare your scripts for the offer call.

---

## 🚀 Moving from "Understanding" to "Reflex"

This repository provides the **map**. However, reading a map is not the same as driving the car. In a 45-minute Google interview, you don't have time to "remember" these frameworks—they must be **reflexes**.

### **The Simulation-Based System**
If you want to train these reflexes under pressure, we built the **Complete SRE Career Launchpad**. It is a guided training program that turns these frameworks into muscle memory.

👉 **[Get the Full "Google SRE" Training Bundle Here](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

**The Bundle is Indispensable for:**
*   **Practice Scenarios:** 20+ more deep-dive incident simulations.
*   **Interactive Workbooks:** 70+ coding drills in Python and Go with SRE-specific grading.
*   **The 30-Day Blueprint:** A day-by-day checklist of exactly what to do to ensure you hit the "Exceptional" signal in every round.

---
**Good luck with your loop. Stop guessing. Start architecting.**
