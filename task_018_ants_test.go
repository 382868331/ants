package ants_test

import "testing"

// TestTask018 isolates the worker缓存 regression.
func TestTask018(t *testing.T) {
	TestPurgePool(t)
}

// TestTask018Repeat guards deterministic behavior across repeated calls.
func TestTask018Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask018(t)
	}
}
