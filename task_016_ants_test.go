package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts016Primary(t *testing.T) {
 q:=newWorkerStack(1); if !q.isEmpty() || q.len()!=0 { t.Fatalf("empty=%v len=%d",q.isEmpty(),q.len()) }
}
