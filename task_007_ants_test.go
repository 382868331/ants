package ants_test

import "testing"

// TestTask007 isolates the Unicode任务名 regression.
func TestTask007(t *testing.T) {
	TestAntsPoolGetWorkerFromCache(t)
}
