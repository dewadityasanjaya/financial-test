package workerpool

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestProcessPaymentsPreservesOrderAndLimitsConcurrency(t *testing.T) {
	payments := []Payment{
		{ID: "p1", AmountCents: 1000},
		{ID: "p2", AmountCents: 2000},
		{ID: "p3", AmountCents: 3000},
		{ID: "p4", AmountCents: 4000},
	}

	var active int64
	var maxActive int64

	got, err := ProcessPayments(context.Background(), payments, 2, func(ctx context.Context, payment Payment) (Receipt, error) {
		current := atomic.AddInt64(&active, 1)
		for {
			highest := atomic.LoadInt64(&maxActive)
			if current <= highest || atomic.CompareAndSwapInt64(&maxActive, highest, current) {
				break
			}
		}
		defer atomic.AddInt64(&active, -1)

		time.Sleep(20 * time.Millisecond)
		return Receipt{PaymentID: payment.ID, Approved: payment.AmountCents <= 3000}, nil
	})

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if maxActive > 2 {
		t.Fatalf("expected at most 2 active workers, got %d", maxActive)
	}
	for idx, receipt := range got {
		if receipt.PaymentID != payments[idx].ID {
			t.Fatalf("result order mismatch at %d: got %s want %s", idx, receipt.PaymentID, payments[idx].ID)
		}
	}
	if got[3].Approved {
		t.Fatal("expected p4 to be declined")
	}
}

func TestProcessPaymentsReturnsProcessorError(t *testing.T) {
	wantErr := errors.New("processor failed")
	payments := []Payment{{ID: "p1"}, {ID: "p2"}}

	_, err := ProcessPayments(context.Background(), payments, 2, func(ctx context.Context, payment Payment) (Receipt, error) {
		if payment.ID == "p2" {
			return Receipt{}, wantErr
		}
		return Receipt{PaymentID: payment.ID, Approved: true}, nil
	})

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected %v, got %v", wantErr, err)
	}
}

func TestProcessPaymentsRespectsContextCancellation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	payments := []Payment{{ID: "p1"}, {ID: "p2"}, {ID: "p3"}}

	start := time.Now()
	_, err := ProcessPayments(ctx, payments, 2, func(ctx context.Context, payment Payment) (Receipt, error) {
		select {
		case <-time.After(time.Second):
			return Receipt{PaymentID: payment.ID, Approved: true}, nil
		case <-ctx.Done():
			return Receipt{}, ctx.Err()
		}
	})

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline exceeded, got %v", err)
	}
	if elapsed := time.Since(start); elapsed >= 200*time.Millisecond {
		t.Fatalf("expected quick cancellation, took %v", elapsed)
	}
}

