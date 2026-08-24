package ants

import (
 "testing"
 "time"
)

var _ = time.Second

type taskLogger struct{}
func (taskLogger) Printf(string, ...any) {}

func TestTaskAnts001Primary(t *testing.T) {
 Reboot(); defer Release(); if got:=Running(); got!=0 { t.Fatalf("running=%d",got) }
}
