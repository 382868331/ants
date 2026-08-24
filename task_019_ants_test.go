package ants

import (
	"testing"
	"time"
)

func TestTask019MultiPoolFreeAggregatesCapacity(t *testing.T) {
	mp, err := NewMultiPool(2, 3, RoundRobin, WithDisablePurge(true))
	if err != nil {
		t.Fatal(err)
	}
	defer mp.ReleaseTimeout(time.Second)
	if got := mp.Free(); got != 6 {
		t.Fatalf("idle multipool free=%d, want 6", got)
	}
}
