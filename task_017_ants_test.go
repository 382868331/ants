package ants_test

import "testing"

// TestTask017 isolates the 任务标签匹配 regression.
func TestTask017(t *testing.T) {
	TestPoolPanicWithoutHandlerPreMalloc(t)
}
