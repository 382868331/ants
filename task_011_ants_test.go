package ants_test

import "testing"

// TestTask011 isolates the 完成计数 regression.
func TestTask011(t *testing.T) {
	TestAntsPoolWithFuncGenericGetWorkerFromCachePreMalloc(t)
}
