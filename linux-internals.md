# 🐧 The Google SRE’s Tactical Linux Command Cheat Sheet

> **"The 20 Commands That Solve 80% of Production Incidents"**

In a Google SRE interview (especially Troubleshooting and NALS), you aren’t judged on how many flags you memorize. You are judged on your **command fluency**.

Can you manipulate text streams, diagnose system state, and inspect logs without leaving the terminal? This cheat sheet covers the "Bread and Butter" commands you need to have at your fingertips.

---

## 1. The "Log Surgeon" Toolkit (Text Processing)
*Google SREs live in logs. These commands turn raw text into data.*

### `tail` (The Pulse Check)
*   **Command:** `tail -f /var/log/syslog`
*   **Why use it:** Watch a log file in real-time as an incident is happening.
*   **Pro Tip:** `tail -f access.log | grep 500` (Watch only the errors live).
*   **Interview Signal:** "I’d start by tailing the logs to see if the errors are spiking right now."

### `head` & `tail -n` (Safe Inspection)
*   **Command:** `head -n 20 large_file.log`
*   **Why use it:** Inspecting a 10GB file without loading it all into memory (which would crash the server).
*   **Interview Signal:** "I won't open the full file in vim; I'll inspect the first few lines to understand the schema."

### `grep` (The Filter)
*   **Command:** `grep -C 5 "ERROR" app.log`
*   **Why use it:** Finds the error *and* the 5 lines of context before and after it.
*   **Command:** `grep -v "healthz" access.log`
*   **Why use it:** "Invert" match. Hides the noise (like health checks) to see the real traffic.

### `awk` (The Column Extractor)
*   **Command:** `awk '{print $9}' access.log`
*   **Why use it:** Logs are often structured in columns. This grabs just the 9th column (usually the HTTP status code in Nginx logs).
*   **Interview Signal:** "I’ll use awk to strip away the timestamps and IPs so I can focus just on the status codes."

### `sort | uniq -c | sort -nr` (The "Poor Man's MapReduce")
*   **Command:** `cat access.log | awk '{print $9}' | sort | uniq -c | sort -nr`
*   **What it does:** Counts occurrences of unique lines and sorts them descending.
*   **Why use it:** Instantly answers: "What is the top error code?" or "Which IP is hitting us the most?"
*   **Interview Signal:** This combination is the hallmark of a senior sysadmin.

---

## 2. The "System Doctor" Toolkit (Resource Diagnosis)
*Is the box sick? These commands tell you why.*

### `top` (The Dashboard)
*   **Command:** `top` (then press `1` to see per-CPU usage).
*   **Why use it:** Instant view of CPU, Memory, and Load Average.
*   **Pro Tip:** Look at `%wa` (I/O Wait). If high, it’s a disk problem, not a code problem.

### `lsof` (The Leak Detector)
*   **Command:** `lsof -p <PID>`
*   **Why use it:** Lists "List Open Files." Crucial for debugging "Too Many Open Files" errors.
*   **Interview Signal:** "I suspect a file descriptor leak. I'll run `lsof` to see if the process is holding onto deleted files or old sockets."

### `df -h` & `du -sh` (The Disk Space Hunter)
*   **Command:** `df -h` (System-wide disk usage).
*   **Command:** `du -sh *` (Size of current directory contents).
*   **Why use it:** When the database stops writing, check disk space first.

### `ps` (The Process Investigator)
*   **Command:** `ps aux | grep <process_name>`
*   **Why use it:** Check if a process is running, who owns it, and its PID.
*   **Pro Tip:** Look for processes in `D` state (Uninterruptible Sleep) or `Z` state (Zombie).

---

## 3. The "Network Plumber" Toolkit
*Is the pipe broken?*

### `netstat` or `ss` (The Connection Checker)
*   **Command:** `netstat -plunt` or `ss -plunt`
*   **Why use it:** Shows which ports are listening (`-l`) and which process (`-p`) owns them.
*   **Interview Signal:** "I'll check if the webserver is actually listening on port 80."

### `curl -v` (The Handshake Inspector)
*   **Command:** `curl -v https://google.com`
*   **Why use it:** The `-v` (verbose) flag shows the TLS handshake, headers, and IP resolution.
*   **Interview Signal:** "I'll curl the local endpoint to see if the issue is inside the app or at the load balancer level."

### `dig` (The DNS Detective)
*   **Command:** `dig +short google.com`
*   **Why use it:** Verifies if DNS resolution is working and what IP it returns.

---

## 4. Bonus: The "Emergency" Loop
*Use this in a coding interview when you need to monitor something but don't have a monitoring system.*

**The "Watch" Command**
*   **Command:** `watch -n 1 "netstat -an | grep ESTABLISHED | wc -l"`
*   **What it does:** Runs the command every 1 second and updates the screen.
*   **Why use it:** "I want to watch the active connection count in real-time while I ramp up the load test."

---

## 🚀 Ready to Go Deeper?

Knowing the commands is Step 1. Knowing **how to use them to debug a Kernel Panic** or **CPU Throttling** is Step 2.

The full **Linux Internals Playbook** (part of the Complete SRE Career Launchpad) covers:
*   debugging `OOMKilled` events using `dmesg` and `cgroup` metrics.
*   Using `strace` to find why a process is hanging.
*   Analyzing **CPU Throttling** in Kubernetes using `/proc/sched_debug`.

👉 **[Get the Full Linux Internals Playbook Here](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**
