# 💬 The Google SRE-STAR(M) Method (Behavioral Interview Guide)

> **"Google doesn’t hire for what you know. It hires for how you think."**

Most candidates fail the behavioral round because they tell "hero stories" ("I stayed up all night and fixed the bug").

Google wants **Systemic stories** ("I fixed the bug, then I wrote a playbook and added a canary test so it never happens again").

To pass, you must upgrade the standard STAR method to **STAR(M)**.

---

## 1. The STAR(M) Framework

| Step | Description | Example |
| :--- | :--- | :--- |
| **Situation** | Context in 1 sentence | "We had a regional outage due to a bad config push." |
| **Task** | Your goal | "Stabilize traffic & restore service for 1M users." |
| **Action** | The key steps you took | "Drained traffic → rolled back config → verified error rates." |
| **Result** | The outcome | "Outage contained within 10 minutes, <2% user impact." |
| **(M)etrics** | **Data-backed proof** | "Reduced MTTR by 70%, saved $50k in SLA credits." |

**The (M)etric is the Senior Signal.** It proves you measure impact.

---

## 2. Power Phrases for Interviews
*Use these to signal "Googliness" (Psychological Safety, Consensus, Data).*

*   **Bad:** "I fixed it myself."
*   **Good:** "I coordinated with the DB and Network teams to isolate the issue."

*   **Bad:** "The other team messed up the code."
*   **Good:** "We identified a gap in the testing pipeline that allowed the bug through."

*   **Bad:** "I think we should do X."
*   **Good:** "The data suggests X will reduce latency by 20%."

---

## 3. The "Blameless" Culture Check
Google SREs practice **Blameless Postmortems**. You will be tested on this.

**If asked: "Tell me about a time you made a mistake."**
1.  **Own it immediately:** "I pushed a bad config that caused 500 errors."
2.  **Explain the fix:** "I rolled back instantly."
3.  **Explain the SYSTEM fix:** "I realized our CI pipeline didn't catch this, so I added a validator to prevent anyone else from making the same mistake."

---

## 🚀 Build Your Story Bank

You need 5-6 polished stories before the interview.

The **Behavioral & Reliability Scenario Vault** (part of the full bundle) provides **10+ Gold-Standard SRE Stories** covering:
*   "Tell me about a time you disagreed with a manager."
*   "Tell me about a time you had to compromise."
*   "Tell me about your biggest failure."

Each story is written in the **STAR(M)** format so you can adapt it to your own experience.

👉 **[Get the Full Behavioral Vault Here](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**
