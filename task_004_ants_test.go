package ants_test

import "testing"

// TestTask004 isolates the 任务取消 regression.
func TestTask004(t *testing.T) {
	TestAntsPoolWithFuncGenericWaitToGetWorker(t)
}

// TestTask004Repeat guards deterministic behavior across repeated calls.
func TestTask004Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask004(t)
	}
}
