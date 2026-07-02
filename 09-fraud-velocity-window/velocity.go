package velocity

import "time"

type Transaction struct {
	ID          string
	CardID      string
	AmountCents int64
	CreatedAt   time.Time
}

func FlagVelocity(transactions []Transaction, window time.Duration, thresholdCents int64) []string {
	// TODO: implement.
	return nil
}

