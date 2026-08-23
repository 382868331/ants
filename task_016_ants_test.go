package ants_test

import "testing"

// TestTask016 isolates the 池状态恢复 regression.
func TestTask016(t *testing.T) {
	TestPoolPanicWithoutHandler(t)
}
