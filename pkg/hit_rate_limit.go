package pkg

import "time"

// RateLimiter provides a mechanism to limit the number of allowed hits within a time window.
type RateLimiter struct {
	hitLimit    int           // The maximum number of allowed hits per window
	maxHitLimit int           // The current remaining hits in the window
	window      time.Duration // The duration of the rate limit window (in seconds)
	startTime   time.Time
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
		hitLimit:    hitLimit,
		maxHitLimit: hitLimit,
		window:      window,
	}
}

// ResetHitRateLimit resets the current hit counter to the maximum hit limit.
//
// No input parameters. No return value.
func (r *RateLimiter) ResetHitRateLimit() {
	r.maxHitLimit = r.hitLimit
	r.startTime = time.Time{}
}

// IsHitRateLimit checks if the rate limit has been reached.
//
// If the reset timer is not set, it schedules a reset after the window duration.
// Each call decrements the remaining hit count. When the count is below zero, the rate limit is considered reached.
//
// No input parameters.
// Returns:
//   - bool: true if the rate limit is reached, false otherwise
func (r *RateLimiter) IsHitRateLimit() bool {
	if r.startTime.IsZero() {
		r.startTime = time.Now()
	} else if time.Since(r.startTime) >= r.window*time.Second {
		r.ResetHitRateLimit()
	}

	if r.maxHitLimit-1 >= 0 {
		r.maxHitLimit = r.maxHitLimit - 1
		return false
	}

	return true
}

// GetHitLimit returns the configured maximum number of allowed hits per window.
//
// No input parameters.
// Returns:
//   - int: the hit limit
func (r *RateLimiter) GetHitLimit() int {
	return r.hitLimit
}

// GetMaxHitLimit returns the current remaining hits in the window.
//
// No input parameters.
// Returns:
//   - int: the remaining hit count
func (r *RateLimiter) GetMaxHitLimit() int {
	return r.maxHitLimit
}
