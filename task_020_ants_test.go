package ants

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestTask020RebootRestartsRoundRobinAtZero(t *testing.T) {
	mp, err := NewMultiPool(3, 1, RoundRobin, WithDisablePurge(true))
	if err != nil {
		t.Fatal(err)
	}
	if err := mp.ReleaseTimeout(time.Second); err != nil {
		t.Fatal(err)
	}
	mp.Reboot()
	defer mp.ReleaseTimeout(time.Second)
	if atomic.LoadInt32(&mp.state) != OPENED {
		t.Fatal("reboot did not reopen multipool")
	}
	if got := mp.next(RoundRobin); got != 0 {
		t.Fatalf("first index after reboot=%d, want 0", got)
	}
}
