package workerpool

import "context"

type Payment struct {
	ID          string
	AmountCents int64
}

type Receipt struct {
	PaymentID string
	Approved  bool
}

type Processor func(ctx context.Context, payment Payment) (Receipt, error)

func ProcessPayments(ctx context.Context, payments []Payment, workerCount int, processor Processor) ([]Receipt, error) {
	// TODO: implement.
	return nil, nil
}
