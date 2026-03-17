***

```go
/*
=============================================================================
🟢 GOOGLE SRE CODING PATTERN: THE TOKEN BUCKET RATE LIMITER
=============================================================================

Why Google asks this in coding rounds:
Rate limiting is the ultimate shield against cascading failures and DDoS.
If you cannot write a concurrency-safe rate limiter, you cannot protect production.

The "Senior SRE" Signals demonstrated in this code:
1. Lazy Math over Tickers: Instead of a background goroutine waking up every
   millisecond (which burns CPU), we calculate tokens based on time elapsed 
   since the last request. O(1) efficiency.
2. Context Awareness: Respects downstream timeouts and client cancellations.
3. Thread Safety: Uses sync.Mutex to prevent race conditions across thousands
   of concurrent requests.
4. TryAcquire vs Acquire: Provides both fail-fast and blocking mechanisms.

👉 Want to master the other 70+ production coding patterns?
Get the Complete SRE Career Launchpad: https://aceinterviews.gumroad.com/l/Google_SRE_Interviews_Your_Secret_Bundle_to_Conquer
=============================================================================
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TokenBucket represents a thread-safe rate limiter.
type TokenBucket struct {
	mu         sync.Mutex
	rate       float64   // Tokens added per second
	capacity   float64   // Maximum burst capacity
	tokens     float64   // Current available tokens
	lastRefill time.Time // Last time tokens were updated
}

// NewTokenBucket initializes a new rate limiter.
// rate: How many requests per second to allow.
// burst: How many requests to allow instantly before enforcing the rate.
func NewTokenBucket(rate float64, burst float64) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		capacity:   burst,
		tokens:     burst, // Start full
		lastRefill: time.Now(),
	}
}

// refill is an internal helper to update the token count based on elapsed time.
// Caller MUST hold the mutex.
func (tb *TokenBucket) refill(now time.Time) {
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tokensToAdd := elapsed * tb.rate

	tb.tokens += tokensToAdd
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}

// TryAcquire attempts to take a token immediately (Fail-Fast).
// Returns true if a token was consumed, false if rate limited.
func (tb *TokenBucket) TryAcquire() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	tb.refill(now)

	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
		return true
	}
	return false
}

// Acquire blocks until a token is available or the context expires/cancels.
func (tb *TokenBucket) Acquire(ctx context.Context) error {
	for {
		tb.mu.Lock()
		now := time.Now()
		tb.refill(now)

		if tb.tokens >= 1.0 {
			tb.tokens -= 1.0
			tb.mu.Unlock()
			return nil
		}

		// Calculate exactly how long we need to wait for 1 token
		tokensNeeded := 1.0 - tb.tokens
		waitDuration := time.Duration((tokensNeeded / tb.rate) * float64(time.Second))
		tb.mu.Unlock()

		// Wait for the token to generate OR for the context to timeout
		select {
		case <-time.After(waitDuration):
			// Woke up, loop around and try to grab the token again
			continue
		case <-ctx.Done():
			// Client cancelled the request or timeout hit. Stop waiting!
			return ctx.Err()
		}
	}
}

// ============================================================================
// SIMULATION: How an SRE tests this locally
// ============================================================================

func main() {
	fmt.Println("Starting Rate Limiter Simulation...")
	fmt.Println("Config: 2 requests per second, Burst capacity of 3.")
	fmt.Println("---------------------------------------------------")

	// Allow 2 req/sec, but allow a sudden burst of up to 3.
	limiter := NewTokenBucket(2.0, 3.0)

	// 1. Simulate a massive, instant traffic spike (Flash Crowd)
	fmt.Println("[!] Flash crowd arriving (10 concurrent requests):")
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if limiter.TryAcquire() {
				fmt.Printf("   ✅ Req %02d: Accepted (Served from Burst)\n", id)
			} else {
				fmt.Printf("   ❌ Req %02d: 429 Too Many Requests (Load Shedded)\n", id)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("\n[zZz] Waiting 2 seconds for tokens to refill...\n")
	time.Sleep(2 * time.Second)

	// 2. Simulate a client willing to wait (Blocking with Timeout)
	fmt.Println("[!] Client willing to wait up to 1 second for capacity:")
	
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Drain the bucket first
	limiter.TryAcquire() 
	limiter.TryAcquire()
	limiter.TryAcquire()

	err := limiter.Acquire(ctx)
	if err != nil {
		fmt.Printf("   ❌ Client timed out waiting for capacity: %v\n", err)
	} else {
		fmt.Println("   ✅ Client acquired token after waiting.")
	}

	fmt.Println("\nSimulation Complete. Production APIs protected.")
}
```

***
