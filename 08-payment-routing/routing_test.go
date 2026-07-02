package routing

import (
	"reflect"
	"testing"
)

func TestRoutePayments(t *testing.T) {
	payments := []Payment{
		{ID: "p1", AmountCents: 10_000, Currency: "USD"},
		{ID: "p2", AmountCents: 250_000, Currency: "USD"},
		{ID: "p3", AmountCents: 50_000, Currency: "IDR"},
		{ID: "p4", AmountCents: 1_000, Currency: "EUR"},
	}

	providers := []Provider{
		{ID: "stripe", Currency: "USD", MinAmountCents: 100, MaxAmountCents: 500_000, FeeBps: 80, Priority: 2},
		{ID: "adyen", Currency: "USD", MinAmountCents: 100, MaxAmountCents: 300_000, FeeBps: 70, Priority: 3},
		{ID: "checkout", Currency: "USD", MinAmountCents: 100, MaxAmountCents: 300_000, FeeBps: 70, Priority: 1},
		{ID: "midtrans", Currency: "IDR", MinAmountCents: 10_000, MaxAmountCents: 100_000, FeeBps: 120, Priority: 1},
	}

	got := RoutePayments(payments, providers)
	want := []RouteResult{
		{PaymentID: "p1", ProviderID: "checkout", Routable: true},
		{PaymentID: "p2", ProviderID: "checkout", Routable: true},
		{PaymentID: "p3", ProviderID: "midtrans", Routable: true},
		{PaymentID: "p4", Routable: false},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("routes mismatch\ngot:  %#v\nwant: %#v", got, want)
	}
}

