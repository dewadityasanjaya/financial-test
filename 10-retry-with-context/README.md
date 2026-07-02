# 10 - Retry With Context

Backend services often retry temporary upstream failures, but retries must stop when a request is canceled or times out.

## Task

Implement `Retry`.

Given a context, a maximum attempt count, a delay, and an operation:

- Call the operation until it succeeds or attempts are exhausted.
- Attempt count includes the first call.
- Wait `delay` between failed attempts.
- Stop immediately if the context is canceled.
- Return the operation result on success.
- Return the last operation error when attempts are exhausted.
- Return the context error if cancellation happens before the next attempt.

## Try It

```powershell
go test ./10-retry-with-context
```

