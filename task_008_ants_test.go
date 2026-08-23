package ants_test

import "testing"

// TestTask008 isolates the 容量上界 regression.
func TestTask008(t *testing.T) {
	TestAntsPoolWithFuncGetWorkerFromCache(t)
}
