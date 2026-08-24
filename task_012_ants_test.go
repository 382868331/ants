package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts012Primary(t *testing.T) {
 if mp,err:=NewMultiPool(1,1,LoadBalancingStrategy(99)); err!=ErrInvalidLoadBalancingStrategy || mp!=nil { t.Fatalf("mp=%v err=%v",mp,err) }
}
