# 🚨 Incident Playbook: The Silent TLS Expiry Cascade

> **"A manual certificate renewal is not a fix. It is a scheduled outage waiting to happen."**

In a Senior SRE troubleshooting interview, the prompt will sound incredibly simple: *"At 00:00 UTC, traffic to our main API dropped by 90%. CPU and memory are completely idle. Go."*

If you check the database, the queues, or the application code, you are wasting time. The network edge has severed the connection.

---

## 🛑 The "Leaf vs. Chain" Trap

When candidates finally suspect a certificate issue, they make a fatal diagnostic mistake that reveals a lack of deep network experience.

```diff
- The Failing Move (The "Frontend Dev")
- "I will check if our SSL certificate expired today by looking at our internal dashboard or checking the expiry date of the domain's primary cert."

+ The Passing Move (The "Reliability Architect")
+ "I will immediately test the TLS handshake from an external network using `openssl`. I am not just checking if the leaf certificate expired; I am checking if an Intermediate CA in the chain was rotated or expired, which causes clients to silently drop the connection."
```

Checking internal metrics is useless if the issue is that external browsers no longer trust your certificate chain. You must test from the outside in.

---

## 🔍 1. Symptoms & Initial Triage
*   **User Impact:** Massive drop in successful requests. Users report "Your connection is not private" or `NET::ERR_CERT_DATE_INVALID` errors.
*   **Internal Impact:** If relying on mTLS (Mutual TLS), internal microservices suddenly refuse to talk to each other. `grpc` calls fail with transport security errors.
*   **Metrics:** CPU and memory on the backend drop to near zero. Load balancers show a massive spike in 502s or aborted connections.

## 🛠️ 2. First 5 Commands (Localization)

You must bypass the application and ask the network layer what certificate it is actually serving.

**1. The Ultimate Source of Truth (Inspect the Handshake)**
*   *Command:* `echo | openssl s_client -showcerts -servername api.example.com -connect api.example.com:443 2>/dev/null | openssl x509 -inform pem -noout -text | grep -E "Not After|Issuer"`
*   *Why:* This simulates a real client connection and prints the exact expiration date (`Not After`) and the Issuer. 

**2. The Quick HTTP Check**
*   *Command:* `curl -vI https://api.example.com`
*   *Why:* The `-v` (verbose) flag prints the TLS handshake process. If it fails here, you know it's a crypto/cert issue, not an application 500 error.

**3. Check Local Disk Certificates (If you are on the LB/Proxy)**
*   *Command:* `openssl x509 -enddate -noout -in /etc/ssl/certs/api.crt`
*   *Why:* Verifies if the file actually sitting on the server is expired. If this file is valid, but `s_client` shows an expired cert, your proxy (e.g., Nginx/Envoy) hasn't been reloaded to pick up the new file in memory.

**4. Check Load Balancer Logs**
*   *Command:* `grep "SSL_do_handshake() failed" /var/log/nginx/error.log | tail -n 20`
*   *Why:* Confirms the proxy is dropping connections due to TLS failures.

## 🛡️ 3. Mitigation Sequence (Stabilize)

You need to restore trust immediately. 

1.  **Route Traffic Away:** If this is a regional rotation failure, immediately use DNS or Global Load Balancing to steer traffic to a healthy PoP (Point of Presence) that has a valid certificate.
2.  **Emergency Manual Push:** If the automation failed, manually provision a new certificate (or rollback to the previous one if a bad cert was deployed) and **force reload** the edge proxies.
3.  **Degrade mTLS (Internal Only):** If an internal mTLS CA expired and the entire mesh is down, temporarily configure the service mesh to "Permissive Mode" (allow plaintext) to restore the data plane, *if* security policies allow for SEV-1 mitigation.

## 🔬 4. Root Cause Investigation

Why did the certificate expire in the first place?
*   **Automation Failure:** The `cert-manager` pod crashed, or the ACME (Let's Encrypt) challenge failed due to a recent DNS change, preventing auto-renewal.
*   **The "Hanging" Proxy:** The automated pipeline successfully fetched and wrote the new certificate to the disk, but failed to send a `SIGHUP` signal to Nginx/Envoy to reload the config into memory. It kept serving the old, now-expired cert.
*   **Rate Limits:** You spun up too many environments and hit the Let's Encrypt API rate limits, blocking production renewals.

## 🧱 5. Prevention (The Senior Signal)

To score "Exceptional" (L5/L6), you must prove you will never let a human track an expiration date again.

*   **Proactive Alerting:** Expose `cert_expiry_days_remaining` as a Prometheus metric. Alert at 30 days (Warning) and 14 days (Critical). 
*   **External Synthetic Probers:** Do not rely on internal cron jobs to check file dates. Deploy a "Blackbox Exporter" outside your network that continuously performs TLS handshakes against your public endpoints and alerts if the chain is invalid.
*   **Automated Rotation (Short TTLs):** Shift to short-lived certificates (e.g., 90 days via ACME). If you force the system to rotate certificates constantly, the automation is tested continuously, eliminating the "once-a-year" surprise failure.

---

## 🚀 The "Execution Sequencing" Gap

In an interview, identifying an expired certificate is easy. 

Knowing **how to verify the intermediate chain**, **why the proxy needs a reload**, and **how to architect synthetic probers** is what separates the hires from the rejects.

Google SRE interviews test your **Execution Sequencing** under pressure. If your sequence is wrong, your technical knowledge won't save you.

I built **The Complete Google SRE Career Launchpad** to simulate these exact, high-stakes infrastructure failures. 

👉 **[Get The Complete Google SRE Interview Career Launchpad (Gumroad)](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

**The Full Training System Includes:**
*   **20+ NALSD and Troubleshooting Simulations:** Practice routing around TLS Expiries, BGP Leaks, and Cascading Quota Failures.
*   **The Interviewer Scorecards:** See exactly how the Hiring Committee grades your stabilization strategies.
*   **70+ Production-Grade Coding Drills** in Python & Go.
*   **The Negotiation Playbook:** Word-for-word scripts to secure Top-of-Band compensation.

Don't let an expired certificate freeze your interview. **Train your reflexes.**
