package ants

import "testing"

func TestTask015WorkerExpiryIncludesCutoff(t *testing.T) {
	ws := newWorkerStack(3)
	for _, ts := range []int64{10, 20, 30} {
		_ = ws.insert(&goWorker{lastUsed: ts})
	}
	if got := ws.binarySearch(0, ws.len()-1, 20); got != 1 {
		t.Fatalf("cutoff index=%d, want 1", got)
	}
	if got := ws.binarySearch(0, ws.len()-1, 9); got != -1 {
		t.Fatalf("pre-cutoff index=%d, want -1", got)
	}
	if got := ws.binarySearch(0, ws.len()-1, 30); got != 2 {
		t.Fatalf("latest cutoff index=%d, want 2", got)
	}
}
