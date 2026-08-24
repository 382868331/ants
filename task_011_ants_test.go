package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts011Primary(t *testing.T) {
 if mp,err:=NewMultiPool(0,1,RoundRobin); err!=ErrInvalidMultiPoolSize || mp!=nil { t.Fatalf("mp=%v err=%v",mp,err) }
}

func TestTaskAnts011Boundary(t *testing.T) {
 if mp,err:=NewMultiPool(-1,1,RoundRobin); err!=ErrInvalidMultiPoolSize || mp!=nil { t.Fatalf("mp=%v err=%v",mp,err) }
}
