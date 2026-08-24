package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts008Primary(t *testing.T) {
 var o Options; h:=func(any){}; WithPanicHandler(h)(&o); if o.PanicHandler==nil { t.Fatal("handler lost") }
}
