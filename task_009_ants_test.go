package ants_test

import "testing"

// TestTask009 isolates the 任务结果顺序 regression.
func TestTask009(t *testing.T) {
	TestAntsPoolWithFuncGenericGetWorkerFromCache(t)
}
