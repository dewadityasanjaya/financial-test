package velocity

import "time"

type Transaction struct {
	ID          string
	CardID      string
	AmountCents int64
	CreatedAt   time.Time
}

type cardWindow struct {
	transactions []Transaction
	totalCents   int64
}

func FlagVelocity(transactions []Transaction, window time.Duration, thresholdCents int64) []string {
	windows := make(map[string]cardWindow)
	flagged := []string{}

	for _, tx := range transactions {
		currentWindow := windows[tx.CardID]
		cutoff := tx.CreatedAt.Add(-window)

		valid := currentWindow.transactions[:0]
		for _, oldTx := range currentWindow.transactions {
			if oldTx.CreatedAt.After(cutoff) {
				valid = append(valid, oldTx)
			} else {
				currentWindow.totalCents -= oldTx.AmountCents
			}
		}

		currentWindow.transactions = valid
		currentWindow.transactions = append(currentWindow.transactions, tx)
		currentWindow.totalCents += tx.AmountCents

		if currentWindow.totalCents > thresholdCents {
			flagged = append(flagged, tx.ID)
		}

		windows[tx.CardID] = currentWindow
	}

	return flagged
}
