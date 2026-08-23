package ants_test

import "testing"

// TestTask010 isolates the 任务遍历 regression.
func TestTask010(t *testing.T) {
	TestAntsPoolWithFuncGetWorkerFromCachePreMalloc(t)
}
