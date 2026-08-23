package ants_test

import "testing"

// TestTask005 isolates the worker收尾 regression.
func TestTask005(t *testing.T) {
	TestAntsPoolWithFuncWaitToGetWorkerPreMalloc(t)
}

// TestTask005Repeat guards deterministic behavior across repeated calls.
func TestTask005Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask005(t)
	}
}
