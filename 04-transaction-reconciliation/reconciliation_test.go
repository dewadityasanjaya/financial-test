package reconciliation

import (
	"reflect"
	"testing"
)

func TestReconcile(t *testing.T) {
	internal := []Record{
		{Reference: "a-100", AmountCents: 5000, Currency: "USD"},
		{Reference: "b-200", AmountCents: 7500, Currency: "USD"},
		{Reference: "c-300", AmountCents: 1200, Currency: "IDR"},
	}
	external := []Record{
		{Reference: "a-100", AmountCents: 5000, Currency: "USD"},
		{Reference: "b-200", AmountCents: 7600, Currency: "USD"},
		{Reference: "d-400", AmountCents: 9900, Currency: "USD"},
	}

	got := Reconcile(internal, external)
	want := []Result{
		{Reference: "a-100", Status: Matched},
		{Reference: "b-200", Status: Mismatched},
		{Reference: "c-300", Status: MissingExternal},
		{Reference: "d-400", Status: MissingInternal},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("results mismatch\ngot:  %#v\nwant: %#v", got, want)
	}
}
