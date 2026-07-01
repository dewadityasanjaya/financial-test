package workerpool

import (
	"context"
	"sync"
)

type Payment struct {
	ID          string
	AmountCents int64
}

type Receipt struct {
	PaymentID string
	Approved  bool
}

type Processor func(ctx context.Context, payment Payment) (Receipt, error)

type job struct {
	index   int
	payment Payment
}

type result struct {
	index   int
	receipt Receipt
	err     error
}

func ProcessPayments(ctx context.Context, payments []Payment, workerCount int, processor Processor) ([]Receipt, error) {
	if workerCount <= 0 {
		workerCount = 1
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobs := make(chan job)
	results := make(chan result, len(payments))

	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := range jobs {
				receipt, err := processor(ctx, j.payment)

				select {
				case results <- result{index: j.index, receipt: receipt, err: err}:
				case <-ctx.Done():
					return
				}

				if err != nil {
					cancel()
					return
				}
			}
		}()
	}

	go func() {
		defer close(jobs)

		for i, payment := range payments {
			select {
			case jobs <- job{index: i, payment: payment}:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	receipts := make([]Receipt, len(payments))
	completed := 0

	for res := range results {
		if res.err != nil {
			return nil, res.err
		}

		receipts[res.index] = res.receipt
		completed++

		if completed == len(payments) {
			return receipts, nil
		}
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return receipts, nil
}
