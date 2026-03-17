# 🚨 Incident Playbook: The BGP Route Leak

> **"What do you do when your dashboards are perfectly green, but the Internet thinks you no longer exist?"**

In a Senior/Staff SRE interview, they will test your "Edge" and "Network" intuition. They will give you a scenario where the application is flawless, the databases are fast, but traffic has suddenly dropped off a cliff.

If your troubleshooting stops at the Kubernetes ingress controller, you fail. You must prove you can debug the global routing table.

---

## 🛑 The "Application Tunnel Vision" Trap

When traffic drops, junior engineers assume the application broke and stopped accepting connections. Senior SREs look at the edge.

```diff
- The Failing Move (The "App Developer")
- "Traffic dropped by 80% in South America! The API pods must be deadlocked. I'm going to exec into the containers, check the JVM heaps, and restart the ingress controllers."

+ The Passing Move (The "Reliability Architect")
+ "Our internal API health is 100%, but ingress traffic plummeted. We aren't down; the internet is misrouting our users. I'm going to check external Looking Glass servers to see if our BGP prefixes are being hijacked or leaked by a transit provider in South America."
```

If you try to fix a global routing leak by restarting a Docker container, you instantly broadcast a lack of operational maturity.

---

## 🔍 1. Symptoms & Initial Triage
*   **User Impact:** Massive, sudden drop in QPS (Queries Per Second) from specific geographic regions or specific ISPs. Users report timeouts.
*   **Metrics:** CPU, Memory, and application error rates are completely normal (or suspiciously low due to lack of traffic).
*   **The Trap:** Because no internal alerts fire (since the systems are "healthy"), this is often first reported by users on Twitter/Reddit or via external synthetic monitoring (e.g., Catchpoint, ThousandEyes).

## 🛠️ 2. First 5 Commands (Localization)

You cannot SSH into an ISP's router in another country. How do you get signal?

**1. Verify from the Outside In (External Probing)**
*   *Action:* Use `mtr` or `traceroute` from external jump boxes (or public looking glasses) targeting your service IP.
*   *Why:* You need to see exactly where the packets are dying. If they reach your edge, it's your problem. If they die 5 hops away inside `AS12345`, it's a transit issue.

**2. Check BGP Prefix Announcements**
*   *Action:* Use public BGP routing databases (e.g., bgp.he.net, RIPE stat, or routeviews).
*   *Why:* You are looking to see if a random Autonomous System (AS) suddenly started advertising a more specific route (e.g., a `/24` instead of your `/20`) for your IP space, sucking your traffic into a black hole.

**3. Inspect Edge Traffic at the Packet Level**
*   *Command:* `sudo tcpdump -nni eth0 port 443 and tcp[tcpflags] == tcp-syn`
*   *Why:* Are the initial `SYN` packets even reaching your load balancers? If the packets never arrive, the code cannot be the root cause.

**4. Check Edge Router Logs**
*   *Command:* Check your BGP daemon logs (BIRD, FRR, Quagga) or Cloud Provider edge logs.
*   *Why:* Check if your BGP sessions with your transit providers (peers) have flapped or dropped.

**5. Query Global Synthetic Monitors**
*   *Action:* Look at global ping/HTTP dashboards.
*   *Why:* A BGP leak usually only affects a portion of the internet (whoever accepts the bad route). You need to map the exact blast radius.

## 🛡️ 3. Mitigation Sequence (Stabilize)

You do not control the router that is leaking your routes. You must use **Traffic Engineering** to mitigate.

1.  **DNS Steering (Fastest):** If you use DNS-based global load balancing, immediately update your DNS records to point affected users to a different IP range/region that is not being hijacked. (Assumes you have low TTLs).
2.  **Anycast Withdrawal:** If using Anycast, withdraw the BGP announcements for the affected PoP (Point of Presence). This forces upstream routers to find the next best path to your other healthy datacenters.
3.  **Prepend AS Paths:** If a specific peering link is saturated by a leak, advertise your routes to that peer with a heavily prepended AS Path (e.g., `AS_PATH: 65000 65000 65000`). This makes the route look artificially "long" and ugly to the internet, forcing traffic to take a different path.
4.  **Engage the NOC:** Open an emergency ticket with your upstream transit provider (Tier 1 ISP) to filter out the leaked routes. 

## 🔬 4. Root Cause Investigation

Once traffic is restored via steering, you investigate the routing anomaly.
*   **The Classic "Fat Finger":** A small ISP in another country accidentally leaked their full routing table to a Tier-1 provider without proper filtering. Suddenly, the entire internet thinks the fastest way to Google is through a 10Gbps pipe in a tiny local ISP. The pipe saturates instantly, dropping all packets.

## 🧱 5. Prevention (The Senior Signal)

To score "Exceptional" (L5/L6), you must discuss cryptographic and policy-based routing defenses.

*   **RPKI (Resource Public Key Infrastructure):** Cryptographically sign your BGP Route Origin Authorizations (ROAs). This allows the rest of the internet to automatically reject fake announcements of your IPs because the cryptographic signature won't match.
*   **Strict IRR Filtering:** Ensure your upstream providers strictly filter accepted prefixes based on the Internet Routing Registry (IRR).
*   **Automated BGP Alerting:** Deploy tools like `BGPalerter` or Catchpoint to trigger PagerDuty the second your AS path changes globally or a new ASN advertises your prefix. Do not wait for user support tickets.

---

## 🚀 The "Execution Sequencing" Gap

In a high-pressure interview, it is incredibly tempting to focus on what you can control (your application code) rather than what you can't (the global internet routing table). 

But a Google Reliability Architect knows that **the network is not reliable**. 

If you don't sequence your debugging to verify the network edge *before* tearing apart your microservices, you will fail the Troubleshooting round.

I built **The Complete Google SRE Career Launchpad** to simulate these exact, "outside-in" infrastructure failures. 

👉 **[Get The Complete Google SRE Career Launchpad (Gumroad)](https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer)**

**The Full Training System Includes:**
*   **20+ NALSD and Troubleshooting Simulations:** Practice routing around BGP Leaks, Kernel Panics, and Cascading Quota Failures.
*   **The Interviewer Scorecards:** See exactly how the Hiring Committee grades your stabilization strategies.
*   **70+ Production-Grade Coding Drills** in Python & Go.
*   **The Negotiation Playbook:** Word-for-word scripts to secure Top-of-Band compensation.

Don't let a network trap ruin your loop. **Train your reflexes.**
