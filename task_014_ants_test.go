package ants_test

import "testing"

// TestTask014 isolates the worker尺寸 regression.
func TestTask014(t *testing.T) {
	TestPanicHandler(t)
}

// TestTask014Repeat guards deterministic behavior across repeated calls.
func TestTask014Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask014(t)
	}
}
