# 07 - Context Worker Pool

Worker pools are common in backend systems for processing jobs with controlled concurrency.

## Task

Implement `ProcessPayments`.

Given a `context.Context`, a list of payments, a worker count, and a `Processor` function:

- Process payments using at most `workerCount` goroutines.
- Preserve result order so result `i` belongs to payment `i`.
- Stop scheduling new work when the context is canceled.
- Return an error if any processor call fails.
- Return quickly when the context is canceled.
- Do not leak goroutines.

## Try It

```powershell
go test ./07-context-worker-pool
```

