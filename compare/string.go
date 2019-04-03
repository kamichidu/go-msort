package compare

import (
	"strings"
)

func String(a string, b string) int {
	return strings.Compare(a, b)
}

func StringPtr(order NilOrder) func(*string, *string) int {
	return func(a, b *string) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return String(*a, *b)
	}
}
