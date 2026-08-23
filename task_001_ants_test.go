package ants_test

import "testing"

// TestTask001 isolates the 空任务提交 regression.
func TestTask001(t *testing.T) {
	TestAntsPoolWaitToGetWorker(t)
}

// TestTask001Repeat guards deterministic behavior across repeated calls.
func TestTask001Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask001(t)
	}
}
