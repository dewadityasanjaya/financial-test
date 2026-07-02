# 09 - Fraud Velocity Window

Fraud systems often flag cards or accounts that spend too much within a short time window.

## Task

Implement `FlagVelocity`.

Given transactions sorted by time:

- Track transactions per card.
- For each transaction, consider only previous transactions for the same card within `window`, plus the current transaction.
- Flag the transaction if the total amount in that window is greater than `thresholdCents`.
- Return flagged transaction IDs in input order.
- Expire old transactions so memory does not grow forever.

## Try It

```powershell
go test ./09-fraud-velocity-window
```

