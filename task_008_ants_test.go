package ants_test

import "testing"

// TestTask008 isolates the 容量上界 regression.
func TestTask008(t *testing.T) {
	TestAntsPoolWithFuncGetWorkerFromCache(t)
}

// TestTask008Repeat guards deterministic behavior across repeated calls.
func TestTask008Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask008(t)
	}
}
