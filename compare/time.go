package compare

import (
	"time"
)

func Time(a, b time.Time) int {
	switch {
	case a.Before(b):
		return -1
	case a.After(b):
		return 1
	default:
		return 0
	}
}

func TimePtr(order NilOrder) func(*time.Time, *time.Time) int {
	return func(a, b *time.Time) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Time(*a, *b)
	}
}
