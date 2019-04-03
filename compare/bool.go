package compare

func Bool(a, b bool) int {
	switch {
	case a == b:
		return 0
	case a:
		return -1
	default:
		return 1
	}
}

func BoolPtr(order NilOrder) func(*bool, *bool) int {
	return func(a, b *bool) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Bool(*a, *b)
	}
}
