package ants_test

import "testing"

// TestTask003 isolates the 池容量解析 regression.
func TestTask003(t *testing.T) {
	TestAntsPoolWithFuncWaitToGetWorker(t)
}
