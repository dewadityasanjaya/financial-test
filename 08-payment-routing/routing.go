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
	// TODO: implement.
	return nil
}

