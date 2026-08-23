package ants_test

import "testing"

// TestTask015 isolates the worker句柄 regression.
func TestTask015(t *testing.T) {
	TestPanicHandlerPreMalloc(t)
}
