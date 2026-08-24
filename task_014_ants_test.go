package ants

import "testing"

func TestTask014WorkerStackDetachesNewest(t *testing.T) {
	ws := newWorkerStack(3)
	first := &goWorker{}
	second := &goWorker{}
	third := &goWorker{}
	_ = ws.insert(first)
	_ = ws.insert(second)
	_ = ws.insert(third)
	if got := ws.detach(); got != third {
		t.Fatalf("detached %p, want newest %p", got, third)
	}
}
