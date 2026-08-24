package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts004Primary(t *testing.T) {
 var o Options; want:=2*time.Second; WithExpiryDuration(want)(&o); if o.ExpiryDuration!=want { t.Fatalf("expiry=%v",o.ExpiryDuration) }
}
