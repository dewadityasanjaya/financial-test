# 06 - Concurrent Price Fetcher

Backend services often call multiple upstream services at once, then combine the results.

## Task

Implement `FetchQuotes`.

Given a `context.Context`, a list of symbols, and a `Fetcher` function:

- Fetch all symbols concurrently using goroutines.
- Respect context cancellation and timeout.
- Return successful quotes by symbol.
- Return an error if any fetch fails.
- Return quickly when the context is canceled.
- Do not leak goroutines.

The result map should contain only successfully fetched quotes if there is no error.

## Try It

```powershell
go test ./06-concurrent-price-fetcher
```

