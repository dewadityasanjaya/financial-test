package pricefetcher

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestFetchQuotesFetchesConcurrently(t *testing.T) {
	ctx := context.Background()
	symbols := []string{"AAPL", "MSFT", "GOOG"}

	start := time.Now()
	got, err := FetchQuotes(ctx, symbols, func(ctx context.Context, symbol string) (Quote, error) {
		select {
		case <-time.After(50 * time.Millisecond):
			return Quote{Symbol: symbol, PriceCents: int64(len(symbol)) * 1000}, nil
		case <-ctx.Done():
			return Quote{}, ctx.Err()
		}
	})

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if elapsed := time.Since(start); elapsed >= 120*time.Millisecond {
		t.Fatalf("expected concurrent fetches, took %v", elapsed)
	}
	if len(got) != len(symbols) {
		t.Fatalf("expected %d quotes, got %d", len(symbols), len(got))
	}
	if got["AAPL"].PriceCents != 4000 {
		t.Fatalf("unexpected AAPL quote: %#v", got["AAPL"])
	}
}

func TestFetchQuotesReturnsFetchError(t *testing.T) {
	ctx := context.Background()
	wantErr := errors.New("upstream failed")

	_, err := FetchQuotes(ctx, []string{"AAPL", "FAIL", "MSFT"}, func(ctx context.Context, symbol string) (Quote, error) {
		if symbol == "FAIL" {
			return Quote{}, wantErr
		}
		return Quote{Symbol: symbol, PriceCents: 1000}, nil
	})

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}

func TestFetchQuotesRespectsContextCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := FetchQuotes(ctx, []string{"AAPL", "MSFT"}, func(ctx context.Context, symbol string) (Quote, error) {
		select {
		case <-time.After(time.Second):
			return Quote{Symbol: symbol, PriceCents: 1000}, nil
		case <-ctx.Done():
			return Quote{}, ctx.Err()
		}
	})

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context deadline exceeded, got %v", err)
	}
	if elapsed := time.Since(start); elapsed >= 200*time.Millisecond {
		t.Fatalf("expected quick cancellation, took %v", elapsed)
	}
}

