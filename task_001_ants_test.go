package ants_test

import "testing"

// TestTask001 isolates the 空任务提交 regression.
func TestTask001(t *testing.T) {
	TestAntsPoolWaitToGetWorker(t)
}
