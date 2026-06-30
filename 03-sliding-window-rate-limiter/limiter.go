package ratelimiter

import "time"

type Limiter struct {
	limit    int
	window   time.Duration
	requests map[string][]time.Time
}

func NewLimiter(limit int, window time.Duration) *Limiter {
	return &Limiter{
		limit:    limit,
		window:   window,
		requests: make(map[string][]time.Time),
	}
}

func (l *Limiter) Allow(key string, now time.Time) bool {
	cutoff := now.Add(-l.window)
	timestamps := l.requests[key]

	valid := timestamps[:0]
	for _, ts := range timestamps {
		if ts.After(cutoff) {
			valid = append(valid, ts)
		}
	}

	if len(valid) >= l.limit {
		l.requests[key] = valid
		return false
	}

	valid = append(valid, now)
	l.requests[key] = valid
	return true
}
