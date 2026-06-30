package ratelimiter

import (
	"testing"
	"time"
)

func TestLimiterAllow(t *testing.T) {
	start := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	limiter := NewLimiter(3, time.Minute)

	if !limiter.Allow("acct-1", start) {
		t.Fatal("first request should be allowed")
	}
	if !limiter.Allow("acct-1", start.Add(10*time.Second)) {
		t.Fatal("second request should be allowed")
	}
	if !limiter.Allow("acct-1", start.Add(20*time.Second)) {
		t.Fatal("third request should be allowed")
	}
	if limiter.Allow("acct-1", start.Add(30*time.Second)) {
		t.Fatal("fourth request inside the same window should be rejected")
	}
	if !limiter.Allow("acct-2", start.Add(30*time.Second)) {
		t.Fatal("different keys should be isolated")
	}
	if !limiter.Allow("acct-1", start.Add(61*time.Second)) {
		t.Fatal("request after earliest timestamp expires should be allowed")
	}
}
