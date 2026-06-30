# 01 - Idempotent Ledger

Backend finance systems often need to apply money movements exactly once, even when clients retry requests.

## Task

Implement `ApplyTransactions`.

Given starting account balances in cents and a list of transfer transactions:

- Ignore duplicate transaction IDs after the first successful or failed attempt.
- Move `AmountCents` from `From` to `To`.
- Reject a transaction if the amount is not positive.
- Reject a transaction if either account is unknown.
- Reject a transaction if the source account has insufficient funds.
- Return the final balances and one result per unique transaction ID, in first-seen order.

Use integer cents. Do not use floating point for money.

## Try It

```powershell
go test ./01-idempotent-ledger
```

