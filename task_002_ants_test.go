package ants_test

import "testing"

// TestTask002 isolates the 任务切片 regression.
func TestTask002(t *testing.T) {
	TestAntsPoolWaitToGetWorkerPreMalloc(t)
}
