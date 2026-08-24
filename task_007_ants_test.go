package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts007Primary(t *testing.T) {
 var o Options; WithNonblocking(true)(&o); if !o.Nonblocking { t.Fatal("nonblocking disabled") }
}

func TestTaskAnts007Boundary(t *testing.T) {
 var o Options; WithNonblocking(false)(&o); if o.Nonblocking { t.Fatal("nonblocking enabled") }
}
