# 04 - Transaction Reconciliation

Finance teams often compare internal ledger records with processor or bank statements.

## Task

Implement `Reconcile`.

Given internal records and external records:

- Match records by `Reference`.
- If the reference exists in both lists and amount/currency are equal, mark it `Matched`.
- If the reference exists in both lists but amount or currency differs, mark it `Mismatched`.
- If only internal has the reference, mark it `MissingExternal`.
- If only external has the reference, mark it `MissingInternal`.
- Return results sorted by reference ascending.

Assume references are unique within each input slice.

## Try It

```powershell
go test ./04-transaction-reconciliation
```

