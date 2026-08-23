package ants_test

import "testing"

// TestTask006 isolates the 任务去重 regression.
func TestTask006(t *testing.T) {
	TestAntsPoolWithFuncGenericWaitToGetWorkerPreMalloc(t)
}

// TestTask006Repeat guards deterministic behavior across repeated calls.
func TestTask006Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask006(t)
	}
}
