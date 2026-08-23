package ants_test

import "testing"

// TestTask007 isolates the Unicode任务名 regression.
func TestTask007(t *testing.T) {
	TestAntsPoolGetWorkerFromCache(t)
}

// TestTask007Repeat guards deterministic behavior across repeated calls.
func TestTask007Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask007(t)
	}
}
