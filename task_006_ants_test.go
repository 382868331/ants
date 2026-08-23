package ants_test

import "testing"

// TestTask006 isolates the 任务去重 regression.
func TestTask006(t *testing.T) {
	TestAntsPoolWithFuncGenericWaitToGetWorkerPreMalloc(t)
}
