package ants

import (
	"testing"
	"time"
)

func TestTask017RoundRobinStartsAtFirstPool(t *testing.T) {
	mp, err := NewMultiPool(3, 1, RoundRobin, WithDisablePurge(true))
	if err != nil {
		t.Fatal(err)
	}
	defer mp.ReleaseTimeout(time.Second)
	if got := mp.next(RoundRobin); got != 0 {
		t.Fatalf("first round-robin index=%d, want 0", got)
	}
	if got := mp.next(RoundRobin); got != 1 {
		t.Fatalf("second round-robin index=%d, want 1", got)
	}
}
