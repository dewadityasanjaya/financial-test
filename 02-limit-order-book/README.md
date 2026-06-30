# 02 - Limit Order Book Matching

This is a simplified version of a matching-engine interview problem.

## Task

Implement `MatchOrders`.

Orders arrive in input order. Each order has:

- `Side`: `Buy` or `Sell`
- `PriceCents`
- `Quantity`

Matching rules:

- A buy can match the lowest-price sell where `sell.PriceCents <= buy.PriceCents`.
- A sell can match the highest-price buy where `buy.PriceCents >= sell.PriceCents`.
- Better price wins first.
- For equal prices, earlier order wins first.
- Trade price is the resting order price.
- Partially filled orders keep their remaining quantity in the book.
- Return all trades in execution order.

## Try It

```powershell
go test ./02-limit-order-book
```

