package main

import (
	"fmt"
	"time"

	"github.com/fxmariojevta/hit-rate-limit-check-implementation/pkg"
)

func main() {
	const Window = 4
	const HitLimit = 5
	const BlockedLimit = 4
	blockedCount := 0
	i := 0
	startTime := time.Now()

	rateLimiter := pkg.NewRateLimiter(HitLimit, Window)
	for {
		i++
		limited := rateLimiter.IsHitRateLimit()
		status := "allowed"
		if limited {
			status = "blocked (rate limit reached)"
		}
		fmt.Printf("Hit %d: %s | elapsed time: %.2f seconds\n", i, status, time.Since(startTime).Seconds())
		if limited {
			blockedCount++
		}
		if blockedCount >= BlockedLimit {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}
}
