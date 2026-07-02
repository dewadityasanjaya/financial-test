# 08 - Payment Routing

Payment systems often choose between multiple providers based on support, limits, fees, and priority.

## Task

Implement `RoutePayments`.

Given payments and provider rules:

- A provider can process a payment only when currency matches.
- Payment amount must be between `MinAmountCents` and `MaxAmountCents`, inclusive.
- Pick the provider with the lowest `FeeBps`.
- If fees are equal, pick the lower `Priority` number.
- If still tied, pick lexicographically smaller provider ID.
- Return one routing result per payment, preserving input order.
- If no provider can process a payment, return `Routable: false`.

`FeeBps` means basis points. `100` bps equals 1%.

## Try It

```powershell
go test ./08-payment-routing
```

