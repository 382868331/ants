package ants

import "testing"

func TestTask018RunningByIndexRejectsUpperBound(t *testing.T) {
	mp := &MultiPool{pools: []*Pool{{}, {}}}
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("upper-bound index panicked: %v", r)
		}
	}()
	if got, err := mp.RunningByIndex(len(mp.pools)); err != ErrInvalidPoolIndex || got != -1 {
		t.Fatalf("got (%d, %v), want (-1, ErrInvalidPoolIndex)", got, err)
	}
}
