# 🚨 Incident Playbook: The Kernel Panic Cascade

> **"At Google scale, hardware and kernels don't just fail. They fail simultaneously, trap your workloads, and take down your control plane."**

In a senior SRE troubleshooting interview, the interviewer will eventually strip away your application metrics. They will tell you the dashboards are blank, SSH is timing out, and Kubernetes is shedding pods.

They are testing your **Machine-Space Intuition**. Can you debug the Operating System itself?

---

## 🛑 The "Reboot" Trap

When faced with an unresponsive node, the instinct of a junior operator is to turn it off and on again. In a Google SRE loop, this is a fatal error.

```diff
- The Failing Move (The "Script Runner")
- "I can't SSH into the box. I will go into the AWS/GCP console and hard reboot the instance. Once it comes back up, I'll check the logs."

+ The Passing Move (The "Reliability Architect")
+ "A hard reboot destroys the forensic evidence in RAM. I will cordon the node to stop new traffic, trigger a crash dump (kdump) to preserve the memory state, and isolate the node for investigation. If this is a bad kernel patch, rebooting it will just put it back into a crash loop."
```

If you destroy the evidence before capturing the signal, you fail the round.

---

## 🔍 1. Symptoms & Initial Triage
*   **User Impact:** Elevated 5xx errors or timeouts.
*   **Metrics:** A sudden drop in active nodes. Kubernetes reports nodes as `NotReady` due to heartbeat timeouts.
*   **The Trap:** Because the application metrics suddenly stop emitting, candidates often assume the *monitoring pipeline* is broken. Always verify host liveness first.

## 🛠️ 2. First 5 Commands (Localization)

You cannot SSH into a panicked or frozen box. How do you get signal?

**1. Access the Out-of-Band Console**
*   *Command:* `gcloud compute connect-to-serial-port <instance-name>` (or AWS EC2 Serial Console).
*   *Why:* This bypasses the frozen network stack. Look for the legendary text: `Kernel panic - not syncing: Fatal exception`.

**2. Inspect the Crash Dump Directory**
*   *Command:* `ls -lh /var/crash/`
*   *Why:* If the system was configured correctly, `kdump` captured the kernel memory into a `vmcore` file right before it died.

**3. If partially alive: Check for D-State Processes**
*   *Command:* `ps -eo pid,stat,wchan,comm | grep D`
*   *Why:* "D-State" is Uninterruptible Sleep. The process is waiting on hardware (usually Disk I/O or a dead NFS mount). `kill -9` will not work. The kernel is stuck waiting for a hardware interrupt that is never coming.

**4. Check Kernel Logs for OOM Cascades**
*   *Command:* `dmesg -T | grep -i -E "panic|oom|killed"`
*   *Why:* A severe Out-Of-Memory event might have killed a critical system daemon (like `systemd` or `kubelet`), leaving the node technically powered on but operationally dead.

**5. Trigger a Manual Crash Dump (The "Magic SysRq" Key)**
*   *Command:* `echo c > /proc/sysrq-trigger`
*   *Why:* If the node is completely deadlocked but hasn't panicked yet, this forces a kernel crash and generates a dump so you can analyze what was holding the locks.

## 🛡️ 3. Mitigation Sequence (Stabilize)

1.  **Assess the Blast Radius:** Is this isolated to one node (likely a bad RAM stick or SSD), or is it 30% of the fleet in one Availability Zone?
2.  **Cordon and Drain:** Tell the orchestrator (Kubernetes) to stop scheduling new work to the affected instances.
3.  **Halt Rollouts:** If a new OS image, kernel version, or DaemonSet was recently deployed, hit the big red **Freeze** button immediately.
4.  **Auto-Remediation:** Ensure your autoscaler is spinning up fresh instances on the *last known good* kernel image to replace the dying ones.

## 🔬 4. Root Cause Investigation

Once the fleet is stable, you analyze the `vmcore` dump using tools like `crash` or `gdb`. 
Typical Google-scale culprits include:
*   **Null Pointer Dereferences** in a recently updated network driver (e.g., NVMe or NIC).
*   **Cgroup Memory Leaks:** A container exhausted its memory, but a kernel bug prevented the OOM killer from executing properly, causing a system-wide lockup.
*   **Hardware Faults:** ECC memory errors.

## 🧱 5. Prevention (The Senior Signal)

To score "Exceptional" (L5/L6), you must engineer the system to survive the next panic automatically.

*   **Automated Reboot on Panic:** Set the sysctl `kernel.panic = 10`. This tells the kernel: "If you panic, write the crash dump, wait 10 seconds, and reboot yourself." Don't wait for a human.
*   **Infrastructure Canaries:** Treat OS and kernel upgrades exactly like application code. Roll them out to 1% of the fleet, bake for 24 hours, and monitor for `node_kernel_panic_total` metrics before expanding.
*   **eBPF Observability:** Deploy eBPF agents to track high-latency syscalls that precede D-state lockups, catching the degradation *before* the panic occurs.

---

## 🚀 The "Execution Sequencing" Gap

Knowing what `kdump` does is easy. 

Knowing **when** to trigger it, **how** to isolate the node, and **why** you must stop the Kubernetes scheduler from sending traffic to it during a live 45-minute interview is what separates the hires from the rejects.

Google SRE interviews test your **Execution Sequencing** under pressure. If your sequence is wrong, your technical knowledge won't save you.

I built **The Complete Google SRE Career Launchpad** to simulate these exact, high-stakes infrastructure failures. 

👉 **[Get The Complete Google SRE Career Launchpad (Gumroad)](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

**The Full Training System Includes:**
*   **20+ NALSD and Troubleshooting Simulations:** Practice routing around Kernel Panics, BGP Leaks, and Cascading Quota Failures.
*   **The Interviewer Scorecards:** See exactly how the Hiring Committee grades your stabilization strategies.
*   **70+ Production-Grade Coding Drills** in Python & Go.
*   **The Negotiation Playbook:** Word-for-word scripts to secure Top-of-Band compensation.

Don't let a kernel panic freeze your interview. **Train your reflexes.**
