# 🐍 🐹 Coding Patterns for Google SREs (Python & Go)

> **"Google SRE coding interviews are not about algorithms.  
> They are about safety, restraint, and operational judgment under constraints."**

In a Google SRE coding round (often called "Practical Scripting"), interviewers are evaluating **Operational Maturity**, not cleverness. Grinding LeetCode Dynamic Programming will not save you here.

They are listening for answers to questions like:
- Will this script **hang forever** in production?
- Does it **crash on bad input**, or degrade gracefully?
- Does it **stream**, or does it blow up memory?
- Can another on-call engineer **read and trust this code at 3 AM**?

This file documents the **coding patterns that pass** — and the LeetCode anti-patterns that silently fail candidates.

---

## 1️⃣ Pattern #1 — The Safety Wrapper (Go Contexts & Timeouts)

**The SRE Rule:** *Every external interaction must have a deadline.*

Hanging tools are production hazards. Google SREs expect **bounded execution by default**. If you make a network call without a timeout, you fail.

```diff
- The LeetCode Way (Unsafe)
- resp, err := http.Get(url) 
- // If the network partitions, this goroutine hangs forever.

+ The Google SRE Way (Safe)
+ ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
+ defer cancel() // Always clean up
+ req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
+ resp, err := client.Do(req)
```
**Interview Signal:** Treating a timeout as a **normal failure mode**, not an unexpected exception.

---

## 2️⃣ Pattern #2 — Streaming Over Loading (Python)

**The SRE Rule:** *Never load unbounded data into memory. Iterate, don’t accumulate.*

Loading a 50GB log file into memory using `.readlines()` is an **instant down-level** to L3. 

```diff
- The LeetCode Way (OOM Crash)
- def parse_logs(path):
-     lines = open(path).readlines() # Loads 50GB into RAM. Server crashes.
-     for line in lines:
-         process(line)

+ The Google SRE Way (O(1) Memory)
+ def parse_logs(path):
+     try:
+         with open(path) as f:
+             for line in f: # Lazy evaluation. Yields one line at a time.
+                 line = line.strip()
+                 if "500" in line:
+                     yield line
+     except FileNotFoundError:
+         logging.error("Log file missing, aborting cleanly.")
```
**Interview Signal:** Using a generator (`yield`), defensive parsing, and ensuring memory usage remains `O(1)` regardless of file size.

---

## 3️⃣ Pattern #3 — Bounded Concurrency (Go)

**The SRE Rule:** *Bounded parallelism beats maximum parallelism.*

Spawning 100,000 goroutines without limits is reckless and will exhaust the host's ephemeral ports or file descriptors.

```diff
- The LeetCode Way (Self-DDoS)
- for _, job := range jobs {
-     go process(job) // Spawns 100k goroutines instantly
- }

+ The Google SRE Way (Backpressure via Semaphore)
+ sem := make(chan struct{}, 10) // Bound concurrency to 10 workers max
+ for _, job := range jobs {
+     sem <- struct{}{} // Blocks if 10 are already running
+     go func(j Job) {
+         defer func() { <-sem }() // Release token
+         process(j)
+     }(job)
+ }
```
**Interview Signal:** Explicit concurrency caps and backpressure awareness.

---

## 4️⃣ Pattern #4 — Exponential Backoff + Jitter

**The SRE Rule:** *Never retry blindly. You will crush the recovering dependency.*

If an API returns a `503 Service Unavailable`, a simple `while True` retry loop creates a "Thundering Herd" that keeps the backend offline forever.

```diff
- The LeetCode Way (Retry Storm)
- while attempts < 3:
-     if call_api(): break
-     time.sleep(1) # All workers wake up at the exact same second

+ The Google SRE Way (Jittered Backoff)
+ delay = base_delay * (2 ** attempt)
+ jitter = random.uniform(0, 0.2 * delay) # Spread out the load
+ time.sleep(delay + jitter)
```
**Interview Signal:** Mentioning "Jitter" proves you have operated distributed systems at scale.

---

## 5️⃣ Pattern #5 — Explicit Resource Cleanup

**The SRE Rule:** *Leaks are reliability bugs.*

```diff
- The LeetCode Way (Resource Leak)
- resp, err := http.Get(url)
- if err != nil { return err }
- body, _ := ioutil.ReadAll(resp.Body)
- // Leaks the file descriptor

+ The Google SRE Way (Defensive Defer)
+ resp, err := client.Do(req)
+ if err != nil { return err }
+ defer resp.Body.Close() // Guaranteed execution even on panic
```

---

## 🧾 The Google SRE Coding Checklist

Before saying “I’m done” in the interview, mentally confirm:

1. [ ] **Input Validation:** Does it crash on empty lines or malformed JSON?
2. [ ] **Timeouts:** Will this hang if the network drops?
3. [ ] **Memory Constraints:** Am I reading 100GB into RAM?
4.[ ] **Error Handling:** Did I log the error, or just silently `continue`?
5. [ ] **Readability:** Can an on-call engineer read this at 3 AM?

If yes → you’re likely passing.

---

## 🚀 The Execution Gap: From Reading to Writing

It is easy to look at a `diff` block and nod your head. **It is incredibly hard to write defensive, concurrent, streaming code from scratch while a Google engineer is watching you on a timer.**

This repository teaches you the *patterns*. 
You still need to build the *muscle memory*.

I built a complete **Simulation-Based Training System** for the SRE coding rounds. It forces you to write production-safe code under constraint.

👉 **[The Complete Google SRE Interview Career Launchpad](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

**Included in the Full Bundle:**
*   **70+ Python & Go SRE-style coding drills.**
*   Concurrency + timeout-heavy mock scenarios (e.g., writing a Rate Limiter, a DLQ Reprocessor, and a safe Config Rollback script).
*   **Interviewer-style evaluation rubrics** so you can grade your own code.

Reading patterns builds awareness.  
**Practicing execution builds offers.**

