package ants_test

import "testing"

// TestTask019 isolates the 空任务集 regression.
func TestTask019(t *testing.T) {
	TestPurgePreMallocPool(t)
}
