package pool

import (
	"math"
	"runtime"
)

// Priority represents the maximum number of goroutine resources that can be running
// in parallel in a goroutine worker. Some predefined constants are available, such
// as HIGH, NORMAL, LOW, etc.
type Priority uint

var (
	MAXIMUM = Priority(math.MaxUint32)
	HIGH    = PriorityRational(8, 1)
	NORMAL  = PriorityRational(1, 1)
	LOW     = PriorityRational(1, 4)
)

// PriorityRational creates a priority value from a rational numerator and denominator.
// The rational value is multiplied by the number of CPU threads on the machine.
func PriorityRational(num, den uint) Priority {
	max := uint(runtime.NumCPU()) * num / den
	if max < 1 {
		max = 1
	}
	return Priority(max)
}
