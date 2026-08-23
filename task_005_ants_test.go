package ants_test

import "testing"

// TestTask005 isolates the worker收尾 regression.
func TestTask005(t *testing.T) {
	TestAntsPoolWithFuncWaitToGetWorkerPreMalloc(t)
}
