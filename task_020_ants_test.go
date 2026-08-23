package ants_test

import "testing"

// TestTask020 isolates the 协程池版本 regression.
func TestTask020(t *testing.T) {
	TestNonblockingSubmit(t)
}

// TestTask020Repeat guards deterministic behavior across repeated calls.
func TestTask020Repeat(t *testing.T) {
	for attempt := 0; attempt < 2; attempt++ {
		TestTask020(t)
	}
}
