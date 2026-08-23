package ants_test

import "testing"

// TestTask013 isolates the Windows工作路径 regression.
func TestTask013(t *testing.T) {
	TestAntsPool(t)
}

// TestTask013Repeat guards deterministic behavior across repeated calls.
func TestTask013Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask013(t)
	}
}
