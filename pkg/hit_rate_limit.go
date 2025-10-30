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

// IsHitRateLimit checks if the rate limit has been reached.
//
// Implements a fixed window rate limiter.
// Each window allows up to hitLimit requests.
//
// Returns:
//   - bool: true if the rate limit is reached, false otherwise
func (r *RateLimiter) IsHitRateLimit() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(r.lastRefill)

	// Reset window if enough time has passed
	if elapsed >= r.window {
		r.tokens = float64(r.hitLimit)
		r.lastRefill = now
	}

	// Check if we have tokens available
	if r.tokens < 1 {
		return true // Rate limit reached
	}

	// Consume a token
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
