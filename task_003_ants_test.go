package ants_test

import "testing"

// TestTask003 isolates the 池容量解析 regression.
func TestTask003(t *testing.T) {
	TestAntsPoolWithFuncWaitToGetWorker(t)
}

// TestTask003Repeat guards deterministic behavior across repeated calls.
func TestTask003Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask003(t)
	}
}
