# 📉 The NALS Diagnostic Flowchart  

> **"NALSD is not System Design. It is operational physics under constraints."**

Most candidates fail Google's Non-Abstract Large System Design (NALSD) round because they treat it like a generic whiteboarding interview (e.g., “Design Twitter”, “Add a Redis cache”, “Use Kafka”).

**NALSD is fundamentally different.** 

You are usually given an **existing system** that is:
1. Already in production.
2. Handling massive, Google-scale traffic.
3. Currently broken, hitting a physical limit, or requiring a zero-downtime migration.

The interviewer is evaluating whether you can **reason like a Production SRE**, not a Cloud Architect.

---

## 🚨 The "Physics vs. Magic" Trap

The fastest way to fail NALSD is to assume software can solve physical constraints. 

```diff
- The L4 "Cloud Architect" Mindset
- "To survive a region failure, I will synchronously replicate our 5 Petabyte database to Europe."

+ The L6 "Reliability Architect" Mindset
+ "Wait. 5 Petabytes over a dedicated 10Gbps link takes ~46 days to transfer. Synchronous replication violates the laws of physics for our 200ms latency SLO. We must use async replication and accept data staleness, or change the requirement."
```

If you draw a box on the whiteboard before doing the **Capacity Math**, you have already failed.

---

## 🧠 The NALSD Mental Flowchart (Execution Order)

Do not jump around. Top-tier candidates execute this exact sequence. **Memorize this order.**

### 🟢 PHASE 1: Interrogate the Reality
*Do not draw anything yet.*

1. **USER GOAL:** What is the *human* actually trying to do?
   * *Don't say:* "Write to Blobstore." 
   * *Do say:* "Upload a 5MB photo and see it on their feed in <2 seconds."
2. **CONSTRAINTS:** What is fixed? (Budget, hardware counts, legacy code, speed of light).
3. **FEASIBILITY MATH:** Calculate Bandwidth, IOPS, and Storage. 
   * *The Signal:* **If the math doesn't work, stop.** Tell the interviewer the requirement is physically impossible. This is a pass signal, not a failure.

### 🔵 PHASE 2: Establish the Contract
4. **SLIs / SLOs:** Define how success is measured *before* designing the system.
   * *The Signal:* "Before we discuss architecture, I want to agree that our target is 99.95% availability and p99 latency < 200ms."
5. **DATA FLOW:** Verbally trace the request lifecycle (Entry -> Dependencies -> Storage -> Fan-out).

### 🔴 PHASE 3: Architect for Failure
*Now you can design. But design defensively.*

6. **RELIABILITY MECHANISMS:** Where will this break? Add bounded retries, exponential backoff, and circuit breakers.
7. **SCALING STRATEGY:** Think in *failure domains*, not just throughput. How do we isolate a bad deploy to one cell instead of taking down the global fleet?
8. **OBSERVABILITY:** How do we know it's failing *before* users complain? (Burn-rate alerts, queue depth saturation).

---

## 💥 The "War Room" Protocol (If the prompt is a broken system)

If the interviewer says: *"Latency is spiking in South America,"* you must immediately shift into Incident Command mode.

```diff
- The Failing Sequence (The Debugger)
- 1. "Let me check the application logs."
- 2. "I'll look at the database CPU."
- 3. "Let's find the root cause."

+ The Passing Sequence (The Commander)
+ 1. "Is this a global outage or just South America? (Clarify Blast Radius)"
+ 2. "I am draining traffic from South America to US-East immediately. (Mitigate & Stabilize)"
+ 3. "Now that users are safe, I will investigate the root cause. (Debug)"
```
**Skipping Step 2 (Stabilization) is an automatic down-level.** Root cause analysis comes *after* user safety.

---

## 📐 The “5-S Rule” for NALSD

Keep this checklist in your head. Have you addressed all five?

| S | The NALSD Check | What Interviewers Listen For |
|---|---|---|
| **Scope** | What exactly are we fixing? | Are you narrowing the problem to avoid boiling the ocean? |
| **Scale** | What is the math? | Are you quantifying load (QPS, IOPS) instead of guessing? |
| **SLIs** | How do we measure it? | Do you anchor your design trade-offs in metrics? |
| **Scarcity**| What are the limits? | Do you respect physical limits (Network, Disk, Memory)? |
| **Safety** | How does it break? | Do you fail visibly and cleanly (Graceful Degradation)? |

---

## 🎙️ Power Phrases That Signal Senior SRE Thinking

Use these sparingly to signal operational maturity to the hiring committee:

*   *"Before drawing the architecture, I want to validate the feasibility math."*
*   *"If the math doesn’t work, we should stop here and renegotiate the product requirement."*
*   *"Let’s stabilize user impact before we start hunting for the root cause."*
*   *"I’m intentionally not optimizing this yet; I want to ensure it fails safely first."*

---

## 🚀 The Execution Gap: Knowing vs. Doing

This file teaches you **what** interviewers expect. But NALSD is a high-pressure, 45-minute verbal sprint. 

Knowing the flowchart won't help you if you freeze when the interviewer says: *"Actually, your fallback database just ran out of inodes. Now what?"*

To pass, you must train your reflexes.

I built **The Complete Google SRE Career Launchpad** to simulate these exact conditions. It includes:
*   **10+ Deep-Dive NALSD Scenarios** (e.g., Control Plane failures, Regional Latency Spikes, BGP Leaks).
*   **Math-First Feasibility Drills** (Train your brain to calculate BDP and IOPS instantly).
*   **Interviewer Scorecards** (See exactly how your answers are graded behind closed doors).

👉 **[Get the Complete Google SRE Career Launchpad Here](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

Free resources create awareness.  
**Simulation changes outcomes.**

