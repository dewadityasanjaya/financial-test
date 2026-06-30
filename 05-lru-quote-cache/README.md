# 05 - LRU Quote Cache

LRU cache is a classic backend interview problem. In financial systems it can show up around exchange rates, market quotes, or account metadata.

## Task

Implement an LRU cache for quote prices.

`Cache` should support:

- `Put(symbol, priceCents)` in O(1) average time.
- `Get(symbol)` in O(1) average time.
- When capacity is exceeded, evict the least recently used symbol.
- Both `Get` and `Put` count as usage.
- Updating an existing symbol should overwrite the value and make it most recently used.

You may use Go's standard library.

## Try It

```powershell
go test ./05-lru-quote-cache
```

