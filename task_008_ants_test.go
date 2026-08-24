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

func TestTaskAnts008Boundary(t *testing.T) {
 var o Options; WithPanicHandler(nil)(&o); if o.PanicHandler!=nil { t.Fatal("nil handler changed") }
}
