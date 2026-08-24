package ants

import "testing"

func TestTask003MaxBlockingTasksKeepsLimit(t *testing.T) {
	for _, limit := range []int{0, 1, 7} {
		opts := loadOptions(WithMaxBlockingTasks(limit))
		if opts.MaxBlockingTasks != limit {
			t.Fatalf("limit %d stored as %d", limit, opts.MaxBlockingTasks)
		}
	}
}
