package orderbook

import (
	"reflect"
	"testing"
)

func TestMatchOrders(t *testing.T) {
	orders := []Order{
		{ID: "s1", Side: Sell, PriceCents: 101, Quantity: 5},
		{ID: "s2", Side: Sell, PriceCents: 100, Quantity: 3},
		{ID: "b1", Side: Buy, PriceCents: 100, Quantity: 4},
		{ID: "b2", Side: Buy, PriceCents: 102, Quantity: 10},
		{ID: "s3", Side: Sell, PriceCents: 101, Quantity: 6},
	}

	got := MatchOrders(orders)
	want := []Trade{
		{BuyOrderID: "b1", SellOrderID: "s2", PriceCents: 100, Quantity: 3},
		{BuyOrderID: "b2", SellOrderID: "s1", PriceCents: 101, Quantity: 5},
		{BuyOrderID: "b2", SellOrderID: "s3", PriceCents: 102, Quantity: 5},
	}

	assertTrades(t, got, want)
}

func TestBuyMatchesLowestSellPriceThenEarlierOrder(t *testing.T) {
	orders := []Order{
		{ID: "s-expensive", Side: Sell, PriceCents: 105, Quantity: 2},
		{ID: "s-cheap-1", Side: Sell, PriceCents: 99, Quantity: 1},
		{ID: "s-cheap-2", Side: Sell, PriceCents: 99, Quantity: 2},
		{ID: "s-mid", Side: Sell, PriceCents: 101, Quantity: 5},
		{ID: "b1", Side: Buy, PriceCents: 105, Quantity: 5},
	}

	got := MatchOrders(orders)
	want := []Trade{
		{BuyOrderID: "b1", SellOrderID: "s-cheap-1", PriceCents: 99, Quantity: 1},
		{BuyOrderID: "b1", SellOrderID: "s-cheap-2", PriceCents: 99, Quantity: 2},
		{BuyOrderID: "b1", SellOrderID: "s-mid", PriceCents: 101, Quantity: 2},
	}

	assertTrades(t, got, want)
}

func TestSellMatchesHighestBuyPriceThenEarlierOrder(t *testing.T) {
	orders := []Order{
		{ID: "b-low", Side: Buy, PriceCents: 100, Quantity: 5},
		{ID: "b-high-1", Side: Buy, PriceCents: 105, Quantity: 2},
		{ID: "b-high-2", Side: Buy, PriceCents: 105, Quantity: 3},
		{ID: "s1", Side: Sell, PriceCents: 100, Quantity: 4},
	}

	got := MatchOrders(orders)
	want := []Trade{
		{BuyOrderID: "b-high-1", SellOrderID: "s1", PriceCents: 105, Quantity: 2},
		{BuyOrderID: "b-high-2", SellOrderID: "s1", PriceCents: 105, Quantity: 2},
	}

	assertTrades(t, got, want)
}

func TestDoesNotMatchFutureOrders(t *testing.T) {
	orders := []Order{
		{ID: "b1", Side: Buy, PriceCents: 100, Quantity: 2},
		{ID: "s1", Side: Sell, PriceCents: 99, Quantity: 1},
		{ID: "s2", Side: Sell, PriceCents: 98, Quantity: 1},
	}

	got := MatchOrders(orders)
	want := []Trade{
		{BuyOrderID: "b1", SellOrderID: "s1", PriceCents: 100, Quantity: 1},
		{BuyOrderID: "b1", SellOrderID: "s2", PriceCents: 100, Quantity: 1},
	}

	assertTrades(t, got, want)
}

func assertTrades(t *testing.T, got []Trade, want []Trade) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("trades mismatch\ngot:  %#v\nwant: %#v", got, want)
	}
}
