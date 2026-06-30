# 03 - Sliding Window Rate Limiter

API rate limiting is a common backend screen, especially for payment, trading, and banking APIs.

## Task

Implement `Limiter.Allow`.

The limiter should:

- Track requests per key.
- Allow at most `limit` requests within the previous `window` duration, including the current timestamp.
- Treat timestamps as non-decreasing for each key.
- Keep keys isolated from each other.
- Evict expired timestamps so memory does not grow forever.

## Try It

```powershell
go test ./03-sliding-window-rate-limiter
```

