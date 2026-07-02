package retry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetrySucceedsAfterTemporaryFailures(t *testing.T) {
	var calls int
	temporaryErr := errors.New("temporary")

	got, err := Retry(context.Background(), 3, time.Millisecond, func(ctx context.Context) (string, error) {
		calls++
		if calls < 3 {
			return "", temporaryErr
		}
		return "approved", nil
	})

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if got != "approved" {
		t.Fatalf("expected approved, got %q", got)
	}
	if calls != 3 {
		t.Fatalf("expected 3 calls, got %d", calls)
	}
}

func TestRetryReturnsLastError(t *testing.T) {
	wantErr := errors.New("still failing")

	_, err := Retry(context.Background(), 2, time.Millisecond, func(ctx context.Context) (string, error) {
		return "", wantErr
	})

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}

func TestRetryStopsOnContextCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	var calls int
	start := time.Now()
	_, err := Retry(ctx, 5, time.Second, func(ctx context.Context) (string, error) {
		calls++
		return "", errors.New("temporary")
	})

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline exceeded, got %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected only 1 call before long delay cancellation, got %d", calls)
	}
	if elapsed := time.Since(start); elapsed >= 200*time.Millisecond {
		t.Fatalf("expected quick cancellation, took %v", elapsed)
	}
}

