package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts005Primary(t *testing.T) {
 var o Options; WithPreAlloc(true)(&o); if !o.PreAlloc { t.Fatal("prealloc disabled") }
}

func TestTaskAnts005Boundary(t *testing.T) {
 var o Options; WithPreAlloc(false)(&o); if o.PreAlloc { t.Fatal("prealloc enabled") }
}
