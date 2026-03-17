# 🚨 Incident Playbook: Disk I/O Saturation & The Silent Outage

> **"A full disk is easy to detect. A slow disk is a silent killer."**

In a Google SRE troubleshooting interview, you will often be handed a scenario where the application CPU is low, memory is stable, the network is uncongested, but **p99 latency has spiked from 50ms to 2 seconds**. 

If you start debugging the application code, you fail. The database is choking on disk I/O, and the application threads are piling up waiting for the write to finish.

---

## 🛑 The "Throughput vs. IOPS" Trap

The most common mistake candidates make when debugging storage is looking at the wrong metric. 

```diff
- The Failing Move (The "Dashboard Watcher")
- "I see our cloud volume is rated for 500 MB/s, and we are only writing 40 MB/s. The disk is fine. I'm going to check the application logic for a deadlock."

+ The Passing Move (The "Reliability Architect")
+ "We are only doing 40 MB/s, but our workload is logging thousands of tiny 4KB events per second. I need to check our IOPS (Input/Output Operations Per Second) limits. If our volume is capped at 10,000 IOPS, we are hitting the ceiling, causing I/O wait times to skyrocket."
```

Throughput is how much water flows through the pipe. IOPS is how many buckets of water you can carry at once. **Small, random writes exhaust IOPS long before they exhaust throughput.**

---

## 🔍 1. Symptoms & Initial Triage
*   **User Impact:** Intermittent timeouts on write-heavy APIs (e.g., checkout, logging, uploads).
*   **Metrics:** Application thread pools are exhausted (hanging). Database replication lag is increasing. 
*   **The Trap:** Because CPU utilization looks low (~20%), candidates assume the node has plenty of capacity. They miss that the CPU is low because it is stuck in `iowait`.

## 🛠️ 2. First 5 Commands (Localization)

You need to prove the disk is the bottleneck.

**1. Check I/O Wait (The Smoking Gun)**
*   *Command:* `top` or `vmstat 1`
*   *Why:* Look at the `%wa` (I/O wait) column. If this number is consistently high (e.g., >20%), your CPU is spending its time waiting for the disk to respond.

**2. Check Disk Saturation and Queue Depth**
*   *Command:* `iostat -xz 1`
*   *Why:* This is the most important command. Look at `%util` (is the disk 100% busy?) and `await` (how long are requests waiting in the queue?). If `await` jumps from 2ms to 200ms, your disk is saturated.

**3. Identify the Offending Process**
*   *Command:* `sudo iotop -oPa`
*   *Why:* Which process is actually doing the writing? Is it the database (`postgres`), or is it a runaway logging agent (`fluentd`) competing for the same disk?

**4. Check the Dirty Page Cache**
*   *Command:* `cat /proc/meminfo | grep -i dirty`
*   *Why:* Linux doesn't write to disk immediately; it writes to memory (the page cache) and flushes it later. If the "Dirty" pages are huge, the kernel is desperately trying to flush data to a slow disk.

**5. Check Inode Exhaustion (The "Fake" Full Disk)**
*   *Command:* `df -i`
*   *Why:* If an application wrote 5 million 1KB files, you might have plenty of storage space (`df -h` looks fine), but you have run out of "file pointers" (inodes). The disk refuses new writes.

## 🛡️ 3. Mitigation Sequence (Stabilize)

You cannot "magically" make the disk faster during an incident. You must reduce the load.

1.  **Shed Load (Drop Non-Critical I/O):** Temporarily disable verbose debug logging, or pause background batch jobs (like database backups or compactions) that are competing for IOPS.
2.  **Increase IOPS (Cloud-Native):** If you are on AWS EBS or GCP Persistent Disk, dynamically modify the volume to provision higher IOPS (e.g., switch from `gp2` to `io1`/`io2`).
3.  **Route Traffic Away:** If this is a database primary, failover to a replica in a different zone that isn't experiencing noisy neighbor disk contention.

## 🔬 4. Root Cause Investigation

Once the system is stable, find out *why* the disk saturated.
*   **The "Fsync" Storm:** A developer changed a config so the database forces a synchronous write to disk (`fsync()`) on *every single transaction* instead of batching them. This instantly murders IOPS.
*   **The Noisy Neighbor:** Another container on the same physical host (or a shared cloud volume) started a massive data dump, starving your application of disk time.
*   **Log Sprawl:** An error loop caused the application to write 100x more logs than usual, filling the page cache.

## 🧱 5. Prevention (The Senior Signal)

To score "Exceptional" (L5/L6), you must engineer the system to decouple application latency from disk latency.

*   **Asynchronous Logging:** Never block the main application thread to write a log. Use memory-buffered, asynchronous log shippers.
*   **Tiered Storage:** Put write-heavy sequential logs (like WALs) on dedicated, high-throughput disks. Put random-access database tables on high-IOPS NVMe drives. Do not mix them.
*   **I/O Cgroups (Resource Control):** In Kubernetes, use `blkio` cgroup limits to ensure background tasks (like backups) are throttled and cannot consume more than 10% of the disk's IOPS, protecting the primary database.

---

## 🚀 The "Execution Sequencing" Gap

Knowing `iostat` is table stakes. 

Knowing **when** to look at `iowait` instead of CPU, **how** to throttle background jobs to save the database, and **why** you must batch writes to preserve IOPS is what separates the hires from the rejects.

Google SRE interviews test your **Execution Sequencing** under pressure. If your sequence is wrong, your technical knowledge won't save you.

I built **The Complete Google SRE Career Launchpad** to simulate these exact, high-stakes infrastructure failures. 

👉 **[Get The Complete Google SRE Interview Career Launchpad (Gumroad)](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

**The Full Training System Includes:**
*   **20+ NALSD and Troubleshooting Simulations:** Practice routing around I/O Stalls, BGP Leaks, and Cascading Quota Failures.
*   **The Interviewer Scorecards:** See exactly how the Hiring Committee grades your stabilization strategies.
*   **70+ Production-Grade Coding Drills** in Python & Go.
*   **The Negotiation Playbook:** Word-for-word scripts to secure Top-of-Band compensation.

Don't let a slow disk ruin your loop. **Train your reflexes.**
