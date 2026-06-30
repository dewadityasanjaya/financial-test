package ledger

type Transaction struct {
	ID          string
	From        string
	To          string
	AmountCents int64
}

type Result struct {
	ID      string
	Applied bool
	Reason  string
}

const (
	ReasonDuplicate         = "duplicate"
	ReasonInvalidAmount     = "invalid_amount"
	ReasonUnknownAccount    = "unknown_account"
	ReasonInsufficientFunds = "insufficient_funds"
	ReasonApplied           = "applied"
)

func ApplyTransactions(openingBalances map[string]int64, transactions []Transaction) (map[string]int64, []Result) {
	balances := make(map[string]int64)
	seen := make(map[string]bool)
	result := []Result{}

	for account, balance := range openingBalances {
		balances[account] = balance
	}

	for _, tx := range transactions {
		if seen[tx.ID] {
			continue
		}
		seen[tx.ID] = true

		if tx.AmountCents < 1 {
			result = append(result, Result{
				ID:      tx.ID,
				Applied: false,
				Reason:  ReasonInvalidAmount,
			})
			continue
		}

		fromBalance, fromOK := balances[tx.From]
		_, toOK := balances[tx.To]
		if !fromOK || !toOK {
			result = append(result, Result{
				ID:      tx.ID,
				Applied: false,
				Reason:  ReasonUnknownAccount,
			})
			continue
		}

		if fromBalance < tx.AmountCents {
			result = append(result, Result{
				ID:      tx.ID,
				Applied: false,
				Reason:  ReasonInsufficientFunds,
			})
			continue
		}

		balances[tx.From] -= tx.AmountCents
		balances[tx.To] += tx.AmountCents
		result = append(result, Result{
			ID:      tx.ID,
			Applied: true,
			Reason:  ReasonApplied,
		})
	}

	return balances, result
}
