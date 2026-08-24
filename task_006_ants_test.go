package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts006Primary(t *testing.T) {
 var o Options; WithMaxBlockingTasks(3)(&o); if o.MaxBlockingTasks!=3 { t.Fatalf("max=%d",o.MaxBlockingTasks) }
}
