package ants_test

import "testing"

// TestTask020 isolates the 协程池版本 regression.
func TestTask020(t *testing.T) {
	TestNonblockingSubmit(t)
}
