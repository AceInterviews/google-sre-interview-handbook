# 🎧 Mock Interview Transcripts: NALSD & Troubleshooting

> **“In an SRE interview, the silence of the interviewer is the loudest signal.”**

If you want to know why Senior Software Engineers routinely fail Google SRE loops, read these transcripts. 

This file contains two simulated interviews. They expose the hidden **Interviewer Scorecards** in real-time, showing exactly how candidates are graded during Google's two most notorious rounds.

**Contents:**
1. **[Simulation 1: The NALSD Round](#-simulation-1-the-nalsd-round-global-llm-platform)** (Testing Math, Physics, and Architecture)
2. **[Simulation 2: The Troubleshooting Round](#-simulation-2-the-troubleshooting-round-black-friday-outage)** (Testing Execution Sequencing and Crisis Leadership)

---

# 🏗️ Simulation 1: Google SRE NALSD Round Mock Interview (Global LLM Inference Platform)

> **Why include this:** This is our strongest flagship NALSD scenario for this entire Google SRE bundle. Google's public AI infrastructure direction makes it a credible forward-looking practice problem, although the exact interview prompt is unverified. Senior-level Google SRE-style NALSD mock interview for 2027+ preparation: global LLM inference capacity planning, token throughput, KV-cache sizing, accelerator provisioning, regional failure reserves, and infrastructure BOM.

> **The math that matters:** Input/output token throughput, concurrent streams, accelerator throughput, KV-cache footprint, prefill/decode capacity, and regional failure reserves.

**Interview simulation** · **Senior infrastructure engineering** · **Capacity sizing and BOM**

We will treat this as a live, interviewer-led NALSD round. The candidate will reason aloud, calculate resource requirements, justify architectural decisions, and respond to follow-up questions as they arise.

One distinction upfront: this is a realistic simulated interview, not a transcript of an actual Google interview. The prompt and constraints are supplied; accelerator throughput, model configuration details, and hardware prices are not. Where those missing inputs affect the BOM, the candidate will state assumptions rather than invent benchmark results.

The objective is to demonstrate how a strong candidate approaches the problem in the interview—not to recite a prewritten architecture.

---

## Phase 1 — Clarify the problem (Minutes 0–5)

### INTERVIEWER
> "Design a globally distributed LLM inference platform with hierarchical KV-cache management for Gemini-class models. We need 100,000 QPS globally, 4,000 input tokens on average, 500 output tokens, P99 TTFT under 300 milliseconds, P99 TPOT under 25 milliseconds, and 99.95% availability. How would you design it?"

### CANDIDATE
> "I would start by establishing the workload and the resource constraints. At this scale, the critical question is not just how many requests we can accept, but how much prefill computation, decode throughput, accelerator memory, and network capacity those requests consume.
> 
> I have three clarifications before sizing the system.
> 
> **First**, is 100,000 QPS the peak aggregate across all regions, and should I assume that traffic can be unevenly distributed?
> 
> **Second**, are the input and output token counts averages for initial and subsequent chat turns combined, or should I model separate workload classes?
> 
> **Third**, should I assume a particular accelerator and benchmarked model throughput, or can I use explicit planning assumptions and derive a conditional BOM?"

### INTERVIEWER — SIMULATED RESPONSE
> "Treat 100,000 QPS as peak global traffic. Traffic is uneven, and multi-turn sessions reuse prefixes. Use a representative accelerator configuration, but state your assumptions. We expect regional failure tolerance."

**Candidate's interpretation:** I will design for peak traffic rather than average daily traffic. I will distinguish logical conversation history from reusable KV-cache state. I will also avoid assuming that all requests consume the same amount of accelerator time.

The SLOs are requirements to validate, not performance guarantees that can be claimed before sizing and testing.

---

## Phase 2 — Workload math and the first capacity boundary (Minutes 5–12)

The first calculations establish the amount of work entering the platform.

### 1. Input and output token throughput

**Given:**
- Peak request rate: $100,000\text{ requests/second}$.
- Average input: $4,000\text{ tokens/request}$.
- Average output: $500\text{ tokens/request}$.

**Input-token rate:**
$$100,000 \times 4,000 = 400,000,000$$
That is **400 million input tokens per second** before accounting for cache hits.

**Output-token rate:**
$$100,000 \times 500 = 50,000,000$$
That is **50 million generated tokens per second**.

| Input throughput | Output throughput |
| :--- | :--- |
| **400M** <br>tokens/second | **50M** <br>tokens/second |

*These are logical workload rates before prefix-cache savings, not measured accelerator throughput.*

---

### 2. Concurrent generation streams

If we assume an average request generates 500 tokens at an average of 25 milliseconds per token, the estimated generation duration is:
$$500 \times 0.025 = 12.5\text{ seconds}$$

Using Little's Law, the approximate number of concurrent generation streams is:
$$100,000 \times 12.5 = 1,250,000$$

So the initial planning estimate is **1.25 million concurrent generation streams**.

But I would explicitly tell the interviewer that this estimate assumes the stated 25 milliseconds is representative of average token-generation time. The actual requirement is a P99 TPOT bound, so the average cannot be derived from the P99 alone. We need a measured or separately specified average to finalize concurrency.

---

### 3. The first interviewer challenge

### INTERVIEWER
> "You have calculated 1.25 million streams. Does that mean you need 1.25 million independent accelerator allocations?"

### CANDIDATE
> "No. That is the concurrency across the serving fleet, not the number of accelerators. Continuous batching allows an accelerator to serve multiple sequences concurrently, subject to memory capacity, memory bandwidth, and the latency target.
> 
> I need measured per-accelerator throughput and a validated batch configuration to translate these streams into an accelerator count. I will calculate memory demand next, because the KV cache can limit the feasible concurrency even when compute capacity appears sufficient."

## Phase 3 — KV-cache sizing: the first major physical constraint (Minutes 12–20)

### Candidate's whiteboard calculation
> "I'll use an explicit example KV-cache configuration so the estimate is reproducible. Assume 64 transformer layers, 8 KV heads, a head dimension of 128, and one byte per KV value. These are illustrative assumptions, not a claim about Gemini's internal architecture."

For each token, the cache holds both keys and values:

$$KV_{bytes} = 2 \times L \times H_{KV} \times D \times B$$
$$= 2 \times 64 \times 8 \times 128 \times 1$$
$$= 131,072\text{ bytes}$$

That is **128 KiB per token** before implementation-specific overhead.

For an active request containing approximately 4,250 tokens:

$$4,250 \times 131,072 = 557,056,000\text{ bytes}$$

That is approximately **557 MB per request** under this particular configuration.

This number is large, but it follows directly from the stated assumptions. It is not a universal KV-cache footprint.

---

### Estimate aggregate cache demand

Using the illustrative concurrency estimate of 1.25 million active requests:

$$KV_{total} = 1.25 \times 10^6 \times 557.056\text{ MB}$$
$$\approx 696\text{ PB}$$

| Illustrative uncached aggregate KV footprint |
| :---: |
| # ≈ 696 PB |

This is a workload-derived theoretical footprint, not a proposed HBM allocation.

The estimate assumes the example KV format and dimensions apply to all requests, all estimated token positions remain resident, and prefix sharing provides no reduction. Actual cache residency can differ substantially.

---

### Interviewer challenge

### INTERVIEWER
> "You have just calculated hundreds of petabytes of KV state. Are you proposing to buy that much accelerator memory?"

### CANDIDATE
> "No. The calculation exposes why we cannot retain every request's full KV state in accelerator HBM. It also tells me that I must validate the workload assumptions before committing to the hardware budget.
> 
> I would separate three things: active decode state that must be immediately accessible, reusable prefix state that can be cached, and conversation history that can be persisted and recomputed when necessary.
> 
> I would also verify the actual model's KV dimensions and precision. If the interviewer expects a feasible production BOM, I need benchmarked per-accelerator throughput and a realistic cache-residency target. Otherwise, the aggregate figure is a useful upper-bound scenario, not a purchase order."

*(Interviewer Note: This is an important NALSD moment. The candidate has used math to expose a physical constraint instead of jumping to a familiar architecture).*

### The interviewer pushes further

### INTERVIEWER
> "Fine. The platform must still support multi-turn conversations. How will you manage the cache, how many accelerator nodes do you need across prefill and decode, and what happens to our latency budget when a primary region goes dark?"

---

## 🔒 End of Preview: Sizing the Architecture

You have reached the end of the public preview (Phases 1–3). 

Most candidates fail NALSD interviews at this exact moment: they freeze when faced with physical hardware ceilings like the 696 PB KV-cache footprint, or they throw arbitrary GPU counts at the board without deriving capacity. Passing an L5/L6 Google SRE round requires using mathematical constraints to dictate the architectural boundary before placing a single server.

---

### 🗺️ What’s Inside the Full Walkthrough: Phases 4 through 6

In the complete interview session included in our bundle, the candidate steps through the rest of the 45-minute loop:

* **Phase 4 — Compute, Accelerator, and Network Sizing (Minutes 20–30)**
  * **Quantized Model Footprint:** Sharding a 70B parameter model across 8-accelerator hosts (640 GB aggregate physical HBM) and calculating real usable capacity.
  * **Throughput Sizing Formulas:** Deriving exact prefill ($N_{\text{prefill}}$) and decode ($N_{\text{decode}}$) node counts at target utilization ($u = 60\%$).
  * **Network Fabric Math:** Calculating the 1.6 Gbps+ outbound payload rate and isolating the prefill-to-decode KV transfer bandwidth to prevent congestion-driven tail latency.

* **Phase 5 — Architecture Derived from Sizing Constraints (Minutes 30–37)**
  * **End-to-End Request Path:** Designing the global routing layer, admission control, and prefix-aware scheduling across prefill and decode pools.
  * **Hierarchical Cache Retention:** Engineering the HBM $\rightarrow$ Host DRAM $\rightarrow$ Durable conversation history storage pipeline.
  * **Regional Failover Math:** Sizing surviving regions for the exact +33.3% load shift during an $N-1$ failure without exceeding TTFT/TPOT latency bounds.

* **Phase 6 — Final BOM, SLO Validation, and Conclusion (Minutes 37–45)**
  * **The Complete Production BOM:** Sizing criteria across accelerator hosts, HBM, host DRAM, regional gateways, and the control plane.
  * **Tail-Latency Validation Strategy:** Empirical testing plans for P99 TTFT, inter-token generation latency distributions, and retry amplification safeguards.
  * **Scoring Breakdown:** The exact rubrics Google interviewers use to distinguish an L5 (Senior) from an L6 (Staff) candidate.

---

## 🚀 Get the Full Walkthrough & Complete Interview Bundle

This scenario is just one part of the end-to-end curriculum designed to build your operational reflexes under real interview pressure.

👉 **[Unlock the Full Walkthrough in the Google SRE Interview Career Launchpad](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

### What is inside the complete system:
* **Full-Length NALSD & Troubleshooting Transcripts:** The complete, unabridged walkthroughs, mathematical proofs, and architectural blueprints for this global LLM platform and other flagship scenarios.
* **20+ "War Room" Interview Simulations:** Realistic mock scenarios equipped with the hidden **Interviewer Scorecards**, follow-up traps, and rubric criteria.
* **Linux Internals & Reliability Deep Dives:** Practical drills covering D-state processes, TCP connection backlogs, kernel memory pressure, and distributed consensus.
* **70+ Production-Grade Coding Drills:** Systems-focused Python & Go drills emphasizing concurrency, streaming, backpressure, and defensive design.
* **The Offer Maximizer:** Negotiation frameworks and word-for-word scripts to secure top-of-band L5/L6 compensation packages.

**Stop guessing the math under pressure. Master the simulation.**

---
---

## 🚨 Simulation 2: The Troubleshooting Round (Black Friday Outage)

> **Why include this:** This exposes the hidden "Execution Sequencing" rubric. It proves why finding the bug is less important than stabilizing the system.

**🗣️ INTERVIEWER:** 
"You are the on-call SRE for a global e-commerce checkout service. It is Black Friday. At 14:00 UTC, the p99 latency for the `/checkout` endpoint spikes from 150ms to 4.5 seconds. Global error rates climb to 15%.
You check your dashboards. CPU and Memory on the API pods are at 40% (completely normal). 
What do you do?"

---

### ❌ Act 1: Candidate A (The "Software Engineer")
*Candidate A is a brilliant Senior Backend Developer. They know Kubernetes and Linux. They are optimizing for 'Finding the Bug'.*

**Candidate A:** "Okay, latency is high but CPU is normal. This sounds like a downstream dependency issue or a bad database query holding up the threads. I'm going to `exec` into one of the API pods and run `curl` to test the local endpoint."

**Interviewer:** "The curl command hangs for 4 seconds, then returns a 500 Internal Server Error."

**Candidate A:** "Alright, I want to look at the application logs. I'll search Elasticsearch for the last 5 minutes of errors. What do I see?"

**Interviewer:** "You see thousands of `context deadline exceeded` errors pointing to the downstream Payment Gateway."

**Candidate A:** "Got it. The Payment Gateway is the root cause. I'm going to check the deployment history for the Payment team. Did they push a bad config?"

> ### 🛑 [Interviewer's Internal Notes - Candidate A]
> *   **Time Elapsed:** 8 Minutes.
> *   **User Impact:** 15% of Black Friday checkouts are still dropping. Revenue is burning.
> *   **Execution Sequencing:** FAILED. Candidate jumped straight to Root Cause Analysis (RCA). Did not attempt to stabilize the system.
> *   **Blast Radius:** FAILED. Candidate did not ask if this was a regional or global issue before diving into logs.
> *   **Current Score:** L3 (Junior) / No Hire.

---

### ✅ Act 2: Candidate B (The "Reliability Architect")
*Candidate B has the exact same technical knowledge, but they are optimizing for 'Operational Survivability'.*

**Candidate B:** "15% error rate on Black Friday is a SEV-1. Before I look at any logs, I need to understand the blast radius. Is this happening in all regions, or just one specific availability zone?"

**Interviewer:** "It appears to be happening globally across all clusters."

**Candidate B:** "Okay, since it's global and our compute resources are normal, our worker threads are likely getting blocked by a downstream bottleneck. My immediate priority is to stop the bleeding and restore the user experience. I am **not** going to investigate the root cause yet."

**Interviewer:** "What is your mitigation?"

**Candidate B:** "I will trigger our **Load Shedding** playbook. I want to drop all non-critical background traffic (like analytics batching) to free up connection pools. Simultaneously, I will enable **Circuit Breakers** on the downstream Payment API. We need to fail fast. Returning a clean 'Please try again' error in 50ms is infinitely better than holding a user's connection open for 4.5 seconds and exhausting our ephemeral ports."

> ### 🟢 [Interviewer's Internal Notes - Candidate B]
> *   **Time Elapsed:** 3 Minutes.
> *   **User Impact:** Halted. System is stabilized into a degraded, but safe, state.
> *   **Execution Sequencing:** EXCEPTIONAL. Demonstrated "Mitigate First, Debug Second" reflex without being prompted. 
> *   **Operational Empathy:** Recognized that hanging for 4.5 seconds is worse for users (and infrastructure) than a fast failure.
> *   **Current Score:** L5 / L6 (Senior/Staff) Strong Hire trajectory.

**Interviewer:** "The circuit breaker trips. Your p99 latency drops back to 150ms, but 20% of checkouts are now failing fast. What next?"

**Candidate B:** "Users are out of the 'hanging' state and our infrastructure is protected from a connection-pool collapse. Now that we are stable, I will initiate the root cause investigation. Since CPU is fine, I suspect a hidden kernel or network bottleneck. I'll check the TCP connection backlog (`ss -s`) and look for `D-state` (uninterruptible sleep) processes to see if the network stack is choked."

> ### 🟢 [Interviewer's Internal Notes - Candidate B (Continued)]
> *   **Systems Intuition:** EXCEPTIONAL. Candidate knows that "healthy CPU" during an outage often masks I/O or Kernel queue exhaustion. 
> *   **Hypothesis Discipline:** Candidate states exactly *why* they are running a command before running it.

---

## 🎯 The Core Takeaway

Candidate A and Candidate B both knew how to use logs. Both knew how to check CPU. 

**Candidate A failed because they acted like a developer fixing a bug.**  
**Candidate B passed because they acted like an Incident Commander managing a crisis.**

Google does not grade you on whether you find the broken server. They grade you on the sequence of decisions you make while the system is burning.

---

## 🚀 The Simulation Gap

Reading a transcript creates an "Aha!" moment. But doing this live, while a Google engineer is intentionally withholding information and watching a timer, is a completely different game.

If you jump to `grep` logs under pressure, you will fail the round. You must train the **Stabilize First** reflex until it is automatic.

I built a complete **Simulation-Based Training System** designed to build this exact muscle memory.

👉 **[The Complete Google SRE Interview Career Launchpad (Gumroad)](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

**What is inside the full system:**
*   **20+ "War Room" NALSD and Troubleshooting Simulations** (Just like the prompts above, complete with the hidden interviewer rubrics).
*   **The Linux Internals Deep Dives:** How to answer the "D-state" and "TCP Backlog" questions Candidate B used.
*   **70+ Production-Grade Coding Drills** in Python & Go (Mastering Concurrency, Streaming, and Safety).
*   **The Offer Maximizer:** Word-for-word negotiation scripts to turn your "Strong Hire" score into a Top-of-Band L5/L6 compensation package.

Don't let a sequencing error cost you a $300k+ offer. **Stop guessing. Start simulating.**
