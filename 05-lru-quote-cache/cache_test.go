package lru

import "testing"

func TestCache(t *testing.T) {
	cache := NewCache(2)

	cache.Put("AAPL", 19_000)
	cache.Put("MSFT", 42_000)

	if got, ok := cache.Get("AAPL"); !ok || got != 19_000 {
		t.Fatalf("expected AAPL=19000, got value=%d ok=%v", got, ok)
	}

	cache.Put("GOOG", 28_000)

	if _, ok := cache.Get("MSFT"); ok {
		t.Fatal("MSFT should have been evicted as least recently used")
	}
	if got, ok := cache.Get("AAPL"); !ok || got != 19_000 {
		t.Fatalf("AAPL should still be cached, got value=%d ok=%v", got, ok)
	}

	cache.Put("AAPL", 20_000)
	cache.Put("TSLA", 25_000)

	if _, ok := cache.Get("GOOG"); ok {
		t.Fatal("GOOG should have been evicted after AAPL was updated")
	}
	if got, ok := cache.Get("AAPL"); !ok || got != 20_000 {
		t.Fatalf("expected updated AAPL=20000, got value=%d ok=%v", got, ok)
	}
}
