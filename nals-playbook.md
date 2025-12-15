# 📉 The NALS Diagnostic Flowchart
**Non-Abstract Large System Design (NALS) — How Google Evaluates You**

> **"NALS isn’t about drawing fancy boxes. It’s about reasoning like an SRE under real-world constraints."**

Most candidates fail this round because they treat it like a generic "System Design" interview (designing Twitter from scratch). 

**NALS is different.** You are usually given an *existing* system that is broken, scaling poorly, or needs a major migration. The interviewer wants to see if you can **diagnose, stabilize, and improve** it.

---

## 1. The "War Room" Mental Flow
*Do not jump straight to "Add a Cache." Follow this sequence to show seniority.*

1.  **USER GOAL:** What is the user actually trying to do? (e.g., "Upload a photo", not "Write to Blobstore").
2.  **REQUIREMENTS:** What are the constraints? (Consistency vs. Availability).
3.  **SLIs / SLOs:** How do we measure success? (e.g., "99.9% success rate", "<200ms latency").
4.  **DATA FLOW:** Trace the request life cycle.
5.  **RELIABILITY DESIGN:** How do we handle failure? (Retries, Dead Letter Queues).
6.  **SCALABILITY STRATEGY:** How do we handle 10x growth? (Sharding, Partitioning).
7.  **OBSERVABILITY:** How do we know it's working? (Metrics, Logs, Traces).
8.  **FAILURE / MITIGATION:** What happens when a region dies?
9.  **TRADE-OFFS:** What did we sacrifice to get here?

---

## 2. The “5-S Rule” for Any System Design
*Use this checklist to ensure you don't miss critical constraints.*

| S | Question | Example |
| :--- | :--- | :--- |
| **Scope** | What exactly are we designing? | "A feature flag service for 10M users." |
| **Scale** | What load, data size, or concurrency? | "1M QPS reads, 10K QPS writes." |
| **SLIs** | What does success mean? | "<200ms latency, 99.95% availability." |
| **Storage** | What data layer ensures durability? | "Cloud SQL + Redis cache." |
| **Safety** | What happens when it fails? | "Fail open with stale reads." |

---

## 3. Interview Power Phrases
*Use these to signal you are thinking about Reliability, not just Architecture.*

*   "I’ll start by defining **SLOs** before choosing architecture."
*   "Let’s **fail safely and visibly**, not silently."
*   "We’ll separate **control-plane** from **data-plane** to isolate failure domains."
*   "I’d introduce **circuit breakers** for dependency health."
*   "We can trade **consistency** for **availability** here since it’s read-heavy."

---

## 4. The "Stabilization First" Mindset
In a NALS scenario where a system is "broken" (e.g., "Latency is high in South America"):

1.  **Clarify:** Is it all users? Or just one region?
2.  **Stabilize:** Don't find the root cause yet. Stop the bleeding. (e.g., "Drain traffic from South America to US-East").
3.  **Investigate:** Now look at logs, DNS, and BGP.
4.  **Prevent:** How do we ensure this never happens again?

---

## 🚀 Want to Practice Real Scenarios?

Theory is good. Practice is better.

The **NALS Practice Playbook** (part of the full bundle) contains **10+ real-world simulated scenarios**, including:
*   **"South America is 500ms slower"** (Regional Latency Debugging)
*   **"The Config Service is Down"** (Control Plane Failure)
*   **"Packet Drops under Load"** (Kernel/Network Bottlenecks)

Each scenario includes a step-by-step "Interviewer Scorecard" so you know exactly what they are looking for.

👉 **[Get the Full NALS Practice Playbook Here](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**
