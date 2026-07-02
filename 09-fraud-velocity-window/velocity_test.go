package velocity

import (
	"reflect"
	"testing"
	"time"
)

func TestFlagVelocity(t *testing.T) {
	start := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	transactions := []Transaction{
		{ID: "t1", CardID: "card-1", AmountCents: 4_000, CreatedAt: start},
		{ID: "t2", CardID: "card-1", AmountCents: 3_000, CreatedAt: start.Add(10 * time.Second)},
		{ID: "t3", CardID: "card-2", AmountCents: 9_000, CreatedAt: start.Add(20 * time.Second)},
		{ID: "t4", CardID: "card-1", AmountCents: 5_000, CreatedAt: start.Add(30 * time.Second)},
		{ID: "t5", CardID: "card-1", AmountCents: 7_000, CreatedAt: start.Add(2 * time.Minute)},
	}

	got := FlagVelocity(transactions, time.Minute, 10_000)
	want := []string{"t4"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("flagged mismatch\ngot:  %#v\nwant: %#v", got, want)
	}
}

