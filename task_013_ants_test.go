package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts013Primary(t *testing.T) {
 mp,err:=NewMultiPool(2,1,RoundRobin); if err!=nil { t.Fatal(err) }; defer mp.ReleaseTimeout(time.Second); if _,err=mp.RunningByIndex(2); err!=ErrInvalidPoolIndex { t.Fatalf("err=%v",err) }
}
