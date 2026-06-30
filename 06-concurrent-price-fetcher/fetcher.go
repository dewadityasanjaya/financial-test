package pricefetcher

import "context"

type Quote struct {
	Symbol     string
	PriceCents int64
}

type Fetcher func(ctx context.Context, symbol string) (Quote, error)

func FetchQuotes(ctx context.Context, symbols []string, fetcher Fetcher) (map[string]Quote, error) {
	type result struct {
		symbol string
		quote  Quote
		err    error
	}

	resultCh := make(chan result, len(symbols))

	for _, symbol := range symbols {
		go func(symbol string) {
			quote, err := fetcher(ctx, symbol)
			resultCh <- result{
				symbol: symbol,
				quote:  quote,
				err:    err,
			}
		}(symbol)
	}

	quotes := make(map[string]Quote, len(symbols))

	for range symbols {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case res := <-resultCh:
			if res.err != nil {
				return nil, res.err
			}
			quotes[res.symbol] = res.quote
		}
	}

	return quotes, nil
}
