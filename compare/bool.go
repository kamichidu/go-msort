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
		v := order.compare(a, b)
		if v != 0 {
			return v
		}
		return Bool(*a, *b)
	}
}
