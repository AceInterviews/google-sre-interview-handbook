# 🌙 Cheat Sheet: The Night-Before-Onsite Checklist (The Anxiety Killer)

> **"When others panic, you prepare. When others cram trivia, you review patterns."**

It is 12 hours before your Google SRE loop. 

**Do not try to learn a new concept right now.** Do not open a 500-page textbook on Linux Kernel Development. Do not grind another LeetCode Hard. 

Cramming creates cognitive overload, which leads to freezing under pressure. Your goal tonight is **consolidation and confidence**. 

Close all your other tabs. Go through this 5-step pre-flight checklist in order. 

---

## [ ] Step 1: Review the 5-S Rule (NALSD Mental Framework)

Tomorrow, during the System Design (NALSD) round, you will feel the urge to start drawing boxes as soon as the interviewer finishes speaking. **Resist it.** 

Remind yourself to run through the 5-S Rule out loud *before* designing:

*   **Scope:** "Are we designing this for one region or globally?"
*   **Scale:** "What are the QPS and payload sizes?"
*   **SLIs:** "What is our target availability and latency?"
*   **Scarcity:** "What is our bottleneck? Network bandwidth, disk IOPS, or memory?"
*   **Safety:** "If the control plane fails, does the data plane fail open or closed?"

*Mental Trigger:* "I will not draw a database until I have defined the SLO."

---

## [ ] Step 2: The "Physics of SRE" Math Check

Google interviewers will test if your design violates the laws of physics. Memorize these three fundamental realities so you can do the math instantly on the whiteboard.

1.  **The BDP (Bandwidth-Delay Product) Formula:** 
    *   `Data in Flight = Bandwidth * Round Trip Time (RTT)`. 
    *   *Why you need it:* To prove whether a TCP connection can actually saturate a 10Gbps link over a high-latency cross-continent route.
2.  **The Speed of Light Limitation:** 
    *   US-East to EU-West RTT is roughly **~90ms**. US to APAC is roughly **~200ms**. 
    *   *Why you need it:* If you propose synchronous database replication across the globe, you must explicitly state that it adds 200ms to your p99 latency.
3.  **Disk IOPS Reality:** 
    *   A standard cloud HDD maxes out around **150-500 IOPS**. An SSD/NVMe does **3,000 to 100,000+ IOPS**. 
    *   *Why you need it:* 10,000 QPS of tiny 1KB random writes will destroy a standard disk's IOPS long before it hits the throughput limit. 

---

## [ ] Step 3: Run the 10 "System Doctor" Commands

Open a terminal right now. Type these 10 commands. Look at the output format. Burn the flags into your visual memory so you don't guess them tomorrow.

1.  `top -H` (Look for the thread count and `%wa` I/O wait).
2.  `iostat -xz 1` (Look for `%util` at 100% and high `await` times).
3.  `ss -plunt` (Look at the `Recv-Q` and `Send-Q` for backlog buildup).
4.  `df -i` (Check for Inode exhaustion, not just disk space).
5.  `dmesg -T | tail -n 20` (Check for `OOMKilled` or `Kernel panic`).
6.  `lsof -p <pid>` (Check for file descriptor leaks).
7.  `free -m` (Look at the `buff/cache` column, not just `free`).
8.  `vmstat 1` (Watch the `r` (run queue) and `b` (blocked) columns).
9.  `pidstat -d 1` (Find exactly which process is crushing the disk).
10. `strace -c -p <pid>` (Count which syscall is hanging).

*Mental Trigger:* "If I don't know the exact flag tomorrow, I will explicitly tell the interviewer what signal I am looking for before I 'man page' it."

---

## [ ] Step 4: Rehearse Your 3 STAR(M) Stories Aloud

Do not just read your notes. **Speak them out loud to an empty room.** If you stumble over the words now, you will stumble tomorrow. 

Ensure you have these three specific stories loaded in your short-term memory:
1.  **The "Stabilize First" Story:** A time you stopped a bleeding system *before* you found the root cause.
2.  **The "Toil Buster" Story:** A time you wrote code to eliminate a dangerous, repetitive manual task. 
3.  **The "Pushback" Story:** A time you used data (like error budgets or capacity limits) to push back on a risky deployment or feature launch.

**The (M)etrics Check:** Did you end every story with a hard number? ("Reduced MTTR by 40%", "Saved 10 hours of toil per week", "Contained blast radius to <1% of users").

---

## [ ] Step 5: Memorize the "Magic Question" (For the Recruiter)

Tomorrow isn't just about passing; it's about setting up your compensation negotiation. When the recruiter calls you after a successful loop to discuss "expectations" or an initial offer, do not negotiate against yourself. 

Memorize this exact phrase:

> *"Based on my research into compensation benchmarks for this level, and given my experience and the strong feedback from these interviews, I was hoping we could bring the offer closer to the [Target] range.* 
> 
> ***Is this the most competitive package we can put together?***"

---

## 🛌 Now, Go to Sleep. 

You have the mental models. You know the failure patterns. You have your stories ready. 

**This checklist is your safety net. But passing requires execution.**

If you are reading this weeks before your interview and realize you cannot fluently execute the math, the NALSD sequencing, or the Linux commands listed above, **you need to simulate the pressure.**

I built the **Complete SRE Career Launchpad** to provide that exact simulation. 

It acts as your ultimate insurance policy. It includes:
*   **20+ NALSD & Troubleshooting "War Room" Simulations.**
*   **70+ Production-Grade Coding Drills** (Python & Go).
*   **The Full Interviewer Scorecards** (so you know exactly how you are being graded).
*   **The Word-for-Word Negotiation Playbook.**

👉 **[Secure Your Offer: Get The Complete Google SRE Career Launchpad Here](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

Get some rest. You’ve trained for this.
