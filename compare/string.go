package compare

import (
	"strings"
)

func String(a string, b string) int {
	return strings.Compare(a, b)
}

func StringPtr(order NilOrder) func(*string, *string) int {
	return func(a, b *string) int {
		v := order.compare(a, b)
		if v != 0 {
			return v
		}
		return String(*a, *b)
	}
}
