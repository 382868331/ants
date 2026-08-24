package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts009Primary(t *testing.T) {
 var o Options; l:=taskLogger{}; WithLogger(l)(&o); if o.Logger==nil { t.Fatal("logger lost") }
}
