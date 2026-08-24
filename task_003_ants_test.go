package ants

import "testing"

func TestTask003MaxBlockingTasksKeepsLimit(t *testing.T) {
	for _, limit := range []int{0, 1, 7} {
		opts := loadOptions(WithMaxBlockingTasks(limit))
		if opts.MaxBlockingTasks != limit {
			t.Fatalf("limit %d stored as %d", limit, opts.MaxBlockingTasks)
		}
	}
	if got := loadOptions(WithMaxBlockingTasks(3), WithMaxBlockingTasks(5)).MaxBlockingTasks; got != 5 {
		t.Fatalf("last option should win, got %d", got)
	}
}
