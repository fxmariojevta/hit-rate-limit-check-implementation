package tests

import (
	"testing"

	"github.com/fxmariojevta/hit-rate-limit-check-implementation/pkg"
)

func TestNewRateLimiter(t *testing.T) {
	rl := pkg.NewRateLimiter(5, 3)
	if rl != nil {
		t.Log("NewRateLimiter test passed")
	}
}

func TestResetHitRateLimit(t *testing.T) {
	rl := pkg.NewRateLimiter(3, 1)
	// Decrement maxHitLimit to simulate usage
	rl.IsHitRateLimit()
	rl.IsHitRateLimit()
	if rl.GetMaxHitLimit() != 1 {
		t.Errorf("Setup failed: expected maxHitLimit to be 1, got %d", rl.GetMaxHitLimit())
	}
	rl.ResetHitRateLimit()
	if rl.GetMaxHitLimit() != rl.GetHitLimit() {
		t.Errorf("ResetHitRateLimit failed: expected %d, got %d", rl.GetHitLimit(), rl.GetMaxHitLimit())
	}
}

func TestIsHitRateLimit(t *testing.T) {
	rl := pkg.NewRateLimiter(2, 1) // window=0 for immediate reset
	// First call, should not hit limit
	if rl.IsHitRateLimit() {
		t.Error("IsHitRateLimit failed: should not hit limit on first call")
	}
	// Second call, should not hit limit
	if rl.IsHitRateLimit() {
		t.Error("IsHitRateLimit failed: should not hit limit on second call")
	}
	// Third call, should hit limit
	if !rl.IsHitRateLimit() {
		t.Error("IsHitRateLimit failed: should hit limit on third call")
	}
}
