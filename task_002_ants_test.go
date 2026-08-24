package ants

import "testing"

func TestTask002NonblockingOptionPreservesChoice(t *testing.T) {
	for _, want := range []bool{false, true} {
		opts := loadOptions(WithNonblocking(want))
		if opts.Nonblocking != want {
			t.Fatalf("WithNonblocking(%v) stored %v", want, opts.Nonblocking)
		}
	}
}
