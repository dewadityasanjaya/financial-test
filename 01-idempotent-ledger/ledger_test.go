package ledger

import (
	"reflect"
	"testing"
)

func TestApplyTransactions(t *testing.T) {
	opening := map[string]int64{
		"alice": 10_000,
		"bob":   2_500,
		"bank":  1_000_000,
	}

	transactions := []Transaction{
		{ID: "t1", From: "alice", To: "bob", AmountCents: 1500},
		{ID: "t1", From: "alice", To: "bob", AmountCents: 1500},
		{ID: "t2", From: "bob", To: "alice", AmountCents: 0},
		{ID: "t3", From: "bob", To: "carol", AmountCents: 100},
		{ID: "t4", From: "bob", To: "alice", AmountCents: 10_000},
		{ID: "t5", From: "bank", To: "alice", AmountCents: 50_000},
	}

	gotBalances, gotResults := ApplyTransactions(opening, transactions)

	wantBalances := map[string]int64{
		"alice": 58_500,
		"bob":   4_000,
		"bank":  950_000,
	}
	wantResults := []Result{
		{ID: "t1", Applied: true, Reason: ReasonApplied},
		{ID: "t2", Applied: false, Reason: ReasonInvalidAmount},
		{ID: "t3", Applied: false, Reason: ReasonUnknownAccount},
		{ID: "t4", Applied: false, Reason: ReasonInsufficientFunds},
		{ID: "t5", Applied: true, Reason: ReasonApplied},
	}

	if !reflect.DeepEqual(gotBalances, wantBalances) {
		t.Fatalf("balances mismatch\ngot:  %#v\nwant: %#v", gotBalances, wantBalances)
	}
	if !reflect.DeepEqual(gotResults, wantResults) {
		t.Fatalf("results mismatch\ngot:  %#v\nwant: %#v", gotResults, wantResults)
	}
	if opening["alice"] != 10_000 {
		t.Fatalf("opening balances should not be mutated")
	}
}
