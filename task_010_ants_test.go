package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts010Primary(t *testing.T) {
 var o Options; WithDisablePurge(true)(&o); if !o.DisablePurge { t.Fatal("purge still enabled") }
}

func TestTaskAnts010Boundary(t *testing.T) {
 var o Options; WithDisablePurge(false)(&o); if o.DisablePurge { t.Fatal("purge disabled") }
}
