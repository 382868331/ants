package ants_test

import "testing"

// TestTask009 isolates the 任务结果顺序 regression.
func TestTask009(t *testing.T) {
	TestAntsPoolWithFuncGenericGetWorkerFromCache(t)
}

// TestTask009Repeat guards deterministic behavior across repeated calls.
func TestTask009Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask009(t)
	}
}
