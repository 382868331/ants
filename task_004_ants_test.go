package ants_test

import "testing"

// TestTask004 isolates the 任务取消 regression.
func TestTask004(t *testing.T) {
	TestAntsPoolWithFuncGenericWaitToGetWorker(t)
}
