
# 🐍 🐹 Coding Patterns for Google SREs (Python & Go)

> **"Google SRE coding interviews are not about algorithms.  
They are about safety, restraint, and operational judgment under constraints."**

In a Google SRE coding round, interviewers are evaluating **Operational Maturity**, not cleverness.

They are listening for answers to questions like:

- Does this code **fail safely**?
- Will this script **hang forever** in production?
- Does it **crash on bad input**, or degrade gracefully?
- Does it **stream**, or does it blow up memory?
- Can another on-call engineer **read and trust this code at 3 AM**?

This file documents the **coding patterns that pass** — and the anti-patterns that silently fail candidates.

---

## 🧠 The SRE Coding Mental Model (Before You Write Code)

Strong candidates pause and reason before typing:

1. Is this **batch or streaming**?
2. What happens if **input is malformed**?
3. What happens if a **dependency hangs**?
4. What is the **maximum data size**?
5. How does this behave under **partial failure**?

If you narrate these out loud, you already signal seniority.

---

## 1️⃣ Pattern #1 — The Safety Wrapper (Go)
**Context, Timeouts, and Explicit Cancellation**

> Pattern: *Every external interaction must have a deadline.*

### Why this matters
Hanging tools are production hazards.  
Google SREs expect **bounded execution by default**.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func doTask(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := doTask(ctx)
	if err != nil {
		fmt.Printf("Task failed safely: %v\n", err)
		return
	}
	fmt.Println("Task completed")
}
````

### Interview signals

* Uses `context.WithTimeout`
* Defers `cancel()`
* Explicitly checks `ctx.Done()`
* Treats timeout as a **normal failure mode**, not an exception

---

## 2️⃣ Pattern #2 — Streaming Over Loading (Python)

**Never Load Unbounded Data Into Memory**

> Pattern: *Iterate, don’t accumulate.*

### Why this matters

Loading a 10–100GB log file into memory is an **instant down-level**.

```python
def parse_logs(path):
    try:
        with open(path) as f:
            for line in f:
                line = line.strip()
                if not line:
                    continue

                parts = line.split()
                if len(parts) < 9:
                    continue

                if parts[8] == "500":
                    yield line

    except FileNotFoundError:
        print("File not found")
```

Usage:

```python
for err in parse_logs("/var/log/nginx/access.log"):
    print(err)
```

### Interview signals

* Uses generator (`yield`)
* Defensive parsing
* Constant memory usage
* Safe failure on missing files

---

## 3️⃣ Pattern #3 — Concurrency With Limits (Go)

**Bounded Parallelism Beats Max Parallelism**

> Pattern: *Concurrency with backpressure.*

### Why this matters

Spawning 100,000 goroutines without limits is reckless.

```go
sem := make(chan struct{}, 10)

for _, host := range hosts {
	sem <- struct{}{}
	go func(h string) {
		defer func() { <-sem }()
		checkHealth(h)
	}(host)
}
```

### Interview signals

* Explicit concurrency cap
* Backpressure awareness
* Prevents resource exhaustion

---

## 4️⃣ Pattern #4 — Partial Failure Tolerance

**Do Not Abort the Whole Job**

> Pattern: *Continue on error, collect results.*

Bad candidates:

```python
# crashes on first failure
```

Strong candidates:

```python
errors = 0

for item in items:
    try:
        process(item)
    except Exception:
        errors += 1
        continue
```

### Interview signals

* Understands partial failure
* Prioritizes progress over perfection
* Resilient execution

---

## 5️⃣ Pattern #5 — Explicit Resource Cleanup

**Leaks Are Reliability Bugs**

> Pattern: *Always close, always defer.*

Python:

```python
with open(path) as f:
    ...
```

Go:

```go
resp, err := http.Get(url)
if err != nil { return }
defer resp.Body.Close()
```

### Interview signals

* Knows lifecycle management
* Prevents descriptor leaks
* Thinks beyond “happy path”

---

## 🧠 Execution Sequencing (The Hidden Scoring Axis)

Interviewers are silently scoring **order of operations**:

❌ Bad sequence:

* Parse everything
* Then validate
* Then add error handling

✅ Strong sequence:

1. Validate input
2. Bound execution
3. Stream data
4. Handle errors explicitly
5. Exit safely

This sequencing matters **more than syntax**.

---

## 🚫 Coding Anti-Patterns That Fail Candidates

Interviewers immediately notice:

* `readlines()` on unknown file sizes
* No timeout on network calls
* Unbounded goroutines / threads
* Silent error swallowing
* One-liner clever code with no clarity

These are **operational red flags**.

---

## 🧾 The Google SRE Coding Checklist

Before saying “I’m done”, mentally confirm:

1. Input validation exists
2. Timeouts are enforced
3. Memory use is bounded
4. Errors are handled explicitly
5. Code is readable under stress

If yes → you’re likely passing.

---

## 🚀 Want Real Practice (Not Toy Problems)?

This file shows **patterns**.

The full Google SRE system trains **execution under pressure**, including:

* 70+ Python & Go SRE-style coding drills
* Concurrency + timeout-heavy problems
* Streaming log processors
* Rate limiters, retry loops, DLQ handlers
* Interview-style evaluation rubrics

👉 **The Complete Google SRE Interview Career Launchpad**

[https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)

Reading patterns builds awareness.
Practicing execution builds confidence.

---

## 📌 Final Thought

Google doesn’t hire engineers who write clever code.

Google hires engineers who:

* write **safe** code
* think in **failure modes**
* respect **limits**
* explain **trade-offs clearly**

This file teaches the mindset.
The bundle trains the muscle.
