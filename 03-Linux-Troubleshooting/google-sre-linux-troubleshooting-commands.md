# 🐧 The Google SRE’s Tactical Linux Command Cheat Sheet

> **"The 20 Commands That Solve 80% of Production Incidents"**

In a Google SRE interview (especially Troubleshooting and NALS), you aren’t judged on how many flags you memorize. You are judged on your **command fluency** and **risk awareness**.

```diff
- The Junior Approach
- Runs `cat /var/log/syslog | grep error` on a 50GB file, maxes out memory, and crashes the production node.

+ The Senior Approach
+ Runs `tail -n 5000 /var/log/syslog | grep -i error` to safely extract recent signals without threatening node stability.
```

Can you manipulate text streams, diagnose system state, and inspect logs without leaving the terminal? This cheat sheet covers the "Bread and Butter" commands you need to have at your fingertips.

---

## 1. The "Log Surgeon" Toolkit (Text Processing)
*Google SREs live in logs. These commands turn raw text into data.*

### `tail` (The Pulse Check)
*   **Command:** `tail -f /var/log/syslog`
*   **Why use it:** Watch a log file in real-time as an incident is happening.
*   **Pro Tip:** `tail -f access.log | grep 500` (Watch only the errors live).
*   **Interview Signal:** *"I’d start by tailing the logs to see if the errors are spiking right now, establishing a timeline."*

### `head` & `tail -n` (Safe Inspection)
*   **Command:** `head -n 20 large_file.log`
*   **Why use it:** Inspecting a huge file without loading it all into memory.
*   **Interview Signal:** *"I won't open the full file in vim; I'll inspect the first few lines to understand the schema and avoid OOMing the box."*

### `awk` (The Column Extractor)
*   **Command:** `awk '{print $9}' access.log`
*   **Why use it:** Logs are often structured in columns. This grabs just the 9th column (usually the HTTP status code in Nginx logs).
*   **Interview Signal:** *"I’ll use `awk` to strip away the timestamps and IPs so I can focus purely on the distribution of HTTP status codes."*

### `sort | uniq -c | sort -nr` (The "Poor Man's MapReduce")
*   **Command:** `cat access.log | awk '{print $9}' | sort | uniq -c | sort -nr`
*   **What it does:** Counts occurrences of unique lines and sorts them descending.
*   **Why use it:** Instantly answers: "What is the top error code?" or "Which IP is hitting us the most?"
*   **Interview Signal:** This combination is the absolute hallmark of a senior sysadmin.

---

## 2. The "System Doctor" Toolkit (Resource Diagnosis)
*Is the box sick? These commands tell you why.*

### `top` / `htop` (The Dashboard)
```diff
- The Failing Diagnosis
- "CPU usage is at 90%, we need to scale up."

+ The Passing Diagnosis
+ "CPU is high, but looking at `%wa` (I/O Wait), the CPU is actually just waiting on a slow disk. We have a storage bottleneck, not a compute bottleneck."
```

### `lsof` (The Leak Detector)
*   **Command:** `lsof -p <PID>`
*   **Why use it:** "List Open Files." Crucial for debugging "Too Many Open Files" errors.
*   **Interview Signal:** *"I suspect a file descriptor leak. I'll run `lsof` to see if the process is holding onto deleted files or hanging network sockets."*

### `df -h` vs. `df -i` (The Silent Killer)
```diff
- The Failing Diagnosis
- Runs `df -h`. Sees 50% disk space free. Concludes the disk is fine.

+ The Passing Diagnosis
+ Runs `df -i`. Sees 100% Inode exhaustion because an app wrote 10 million tiny 1KB files. Concludes the disk cannot accept new writes.
```

### `ps` (The Process Investigator)
*   **Command:** `ps aux | grep <process_name>`
*   **Pro Tip:** Look for processes in `D` state (Uninterruptible Sleep) or `Z` state (Zombie).
*   **Interview Signal:** *"A 'D-state' process usually means it's frozen waiting on hardware I/O. `kill -9` won't work here. We need to investigate the storage array."*

---

## 3. The "Network Plumber" Toolkit
*Is the pipe broken?*

### `ss` or `netstat` (The Connection Checker)
*   **Command:** `ss -s` (Summary) or `ss -plunt` (Listening ports)
*   **Why use it:** Shows which ports are listening and the state of all TCP connections.
*   **Interview Signal:** *"I see 5,000 connections in `TIME_WAIT`. We might be exhausting our ephemeral port range."*

### `curl -vI` (The Handshake Inspector)
*   **Command:** `curl -vI https://google.com`
*   **Why use it:** The `-v` (verbose) flag shows the DNS resolution, TLS handshake, and headers without downloading the payload.
*   **Interview Signal:** *"I'll curl the local endpoint to see if the TLS certificate expired or if the issue is deeper in the application logic."*

---

## 4. Bonus: The "Emergency" Loop
*Use this in a coding interview when you need to monitor something but don't have a dashboard.*

**The "Watch" Command**
*   **Command:** `watch -n 1 "ss -s | grep estab"`
*   **What it does:** Runs the command every 1 second and updates the screen in real-time.
*   **Interview Signal:** *"I want to watch the active TCP connection count in real-time while I ramp up the synthetic load test."*

---

## 🚀 The Execution Gap: Knowing vs. Debugging

Knowing the commands is Step 1. 

Knowing **when to use them** to debug a Kernel Panic, CPU Throttling, or a BGP Route Leak under the pressure of a 45-minute Google interview is Step 2.

Most candidates can list these commands. Few can sequence them into a coherent **OODA Loop (Observe, Orient, Decide, Act)** during an active incident simulation.

The full **Linux Internals Playbook** (part of the Complete SRE Career Launchpad) trains this execution muscle. It covers:
*   Debugging `OOMKilled` events using `dmesg` and `cgroup` metrics.
*   Using `strace` to find why a process is hanging on a syscall.
*   Analyzing **CPU Throttling** in Kubernetes using `/proc/sched_debug`.
*   4 specific **Incident Playbooks** simulating real Google production outages.

👉 **[Get the Full Linux Internals Playbook Here](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**
