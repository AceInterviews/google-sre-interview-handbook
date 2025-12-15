
# 🐍 🐹 Coding Patterns for Google SREs (Python & Go)

> **"Google SRE code isn't about complex algorithms. It's about safety, concurrency, and handling failure."**

In a Google SRE coding interview, you are judged on **Operational Maturity**. 
*   Do you handle timeouts?
*   Do you crash on bad input?
*   Does your script load the whole file into memory (bad) or stream it (good)?

Here are the two most common patterns you need to master.

---

## 1. The "Safety Wrapper" (Go)
*Pattern: Running a task with a Timeout and Context.*

**Why it matters:** SRE tools must never hang forever. You must use `context` to handle cancellation and timeouts.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

// Mock function that does work
func doTask(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second): // Simulate work taking 2s
		return nil
	case <-ctx.Done(): // Handle cancellation
		return ctx.Err()
	}
}

func main() {
	// 1. Create a context with a 1-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Always clean up

	// 2. Run the task
	err := doTask(ctx)

	// 3. Handle the error explicitly
	if err != nil {
		fmt.Printf("Task failed: %v\n", err) // Output: context deadline exceeded
	} else {
		fmt.Println("Task finished successfully")
	}
}
```

**Interview Signal:** "I see you used `defer cancel()` and checked `ctx.Done()`. That shows you write production-ready code."

---

## 2. The "Log Streamer" (Python)
*Pattern: Processing a 100GB log file without crashing RAM.*

**Why it matters:** Loading a file with `readlines()` is an instant fail if the file is huge. You must iterate line-by-line.

```python
import sys

def parse_logs(file_path):
    try:
        # 1. Open safely (with statement)
        with open(file_path, 'r') as f:
            # 2. Iterate line by line (Lazy Loading)
            for line in f:
                line = line.strip()
                if not line:
                    continue
                
                # 3. Defensive Parsing (Don't crash on bad lines)
                parts = line.split(" ")
                if len(parts) < 3:
                    continue # Skip malformed lines
                
                status_code = parts[8]
                if status_code == "500":
                    yield line

    except FileNotFoundError:
        print("Error: File not found.")
    except Exception as e:
        print(f"Unexpected error: {e}")

# Usage: This handles a 1TB file easily because it only holds one line in RAM.
for error_line in parse_logs("/var/log/nginx/access.log"):
    print(error_line)
```

**Interview Signal:** "Good job using a generator (`yield`). That keeps memory usage constant regardless of file size."

---

## 3. The SRE Coding Checklist
Before you say "I'm done," check your code for these 5 things:

1.  **Input Validation:** Did you check if the file exists?
2.  **Timeouts:** Will this hang if the network is down?
3.  **Memory Safety:** Are you reading 10GB into RAM?
4.  **Error Handling:** Do you catch exceptions or let the program crash?
5.  **Standard Libs:** Did you stick to standard libraries (os, sys, re, time)?

---

## 🚀 Practice the Real Drills

Reading code is easy. Writing it under pressure is hard.

The **Google SRE Practice Workbooks (Python & Go)** include **70+ production scenarios** to practice:
*   **Problem 27:** Write a Dead-Letter Queue Reprocessor.
*   **Problem 28:** Build a Rate Limiter from scratch using Token Buckets.
*   **Problem 35:** Implement a Log Rotator that is concurrency-safe.

👉 **[Get the Full Coding Workbooks Here](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**
