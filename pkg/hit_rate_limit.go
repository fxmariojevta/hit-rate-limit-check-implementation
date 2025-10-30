package pkg

import (
	"sync"
	"time"
)

// RateLimiter provides a mechanism to limit the number of allowed hits within a time window.
type RateLimiter struct {
	hitLimit   int           // The maximum number of allowed hits per window
	window     time.Duration // The duration of the rate limit window (in seconds)
	mu         sync.Mutex    // Mutex to ensure thread-safe access
	tokens     float64       // Current number of tokens
	lastRefill time.Time     // Last time tokens were refilled
	refilRate  float64       // Rate at which tokens are refilled
}

// NewRateLimiter creates and returns a new RateLimiter.
//
// Parameters:
//   - hitLimit: the maximum number of allowed hits per window
//   - window: the duration of the rate limit window (in seconds)
//
// Returns:
//   - pointer to a RateLimiter instance
func NewRateLimiter(hitLimit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		hitLimit:   hitLimit,
		window:     window,
		tokens:     float64(hitLimit),
		lastRefill: time.Now(),
		refilRate:  float64(hitLimit) / window.Seconds(),
	}
}

// IsHitRateLimit checks if the rate limit has been reached using token bucket algorithm.
//
// The token bucket algorithm works as follows:
// 1. Tokens are added to the bucket at a constant rate (refillRate)
// 2. The bucket has a maximum capacity (hitLimit)
// 3. Each request consumes one token
// 4. If a token is available, the request is allowed
// 5. If no tokens are available, the request is rate limited
//
// Returns:
//   - bool: true if the rate limit is reached, false otherwise
func (r *RateLimiter) IsHitRateLimit() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.lastRefill)

	r.tokens += elapsed.Seconds() * r.refilRate
	if r.tokens > float64(r.hitLimit) {
		r.tokens = float64(r.hitLimit)
	}

	r.lastRefill = now

	if r.tokens < 1 {
		return true
	}

	r.tokens--
	return false
}

// GetHitLimit returns the configured maximum number of allowed hits per window.
//
// No input parameters.
// Returns:
//   - int: the hit limit
func (r *RateLimiter) GetHitLimit() int {
	return r.hitLimit
}

// GetCurrentHitLimit returns the current remaining hits in the window.
//
// No input parameters.
// Returns:
//   - int: the remaining hit count
func (r *RateLimiter) GetCurrentHitLimit() int {
	return int(r.tokens)
}
