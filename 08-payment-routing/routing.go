package routing

type Payment struct {
	ID          string
	AmountCents int64
	Currency    string
}

type Provider struct {
	ID             string
	Currency       string
	MinAmountCents int64
	MaxAmountCents int64
	FeeBps         int
	Priority       int
}

type RouteResult struct {
	PaymentID  string
	ProviderID string
	Routable   bool
}

func RoutePayments(payments []Payment, providers []Provider) []RouteResult {
	results := make([]RouteResult, 0, len(payments))

	for _, payment := range payments {
		var best Provider
		found := false

		for _, provider := range providers {
			if provider.Currency != payment.Currency {
				continue
			}

			if payment.AmountCents < provider.MinAmountCents ||
				payment.AmountCents > provider.MaxAmountCents {
				continue
			}

			if !found || isBetterProvider(provider, best) {
				best = provider
				found = true
			}
		}

		if !found {
			results = append(results, RouteResult{
				PaymentID: payment.ID,
				Routable:  false,
			})
			continue
		}

		results = append(results, RouteResult{
			PaymentID:  payment.ID,
			ProviderID: best.ID,
			Routable:   true,
		})
	}

	return results
}

func isBetterProvider(candidate Provider, current Provider) bool {
	if candidate.FeeBps != current.FeeBps {
		return candidate.FeeBps < current.FeeBps
	}

	if candidate.Priority != current.Priority {
		return candidate.Priority < current.Priority
	}

	return candidate.ID < current.ID
}
