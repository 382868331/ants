package ants_test

import "testing"

// TestTask018 isolates the worker缓存 regression.
func TestTask018(t *testing.T) {
	TestPurgePool(t)
}
