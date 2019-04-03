package compare

func Int(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func IntPtr(order NilOrder) func(*int, *int) int {
	return func(a, b *int) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Int(*a, *b)
	}
}

func Int8(a, b int8) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Int8Ptr(order NilOrder) func(*int8, *int8) int {
	return func(a, b *int8) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Int8(*a, *b)
	}
}

func Int16(a, b int16) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Int16Ptr(order NilOrder) func(*int16, *int16) int {
	return func(a, b *int16) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Int16(*a, *b)
	}
}

func Int32(a, b int32) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Int32Ptr(order NilOrder) func(*int32, *int32) int {
	return func(a, b *int32) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Int32(*a, *b)
	}
}

func Int64(a, b int64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Int64Ptr(order NilOrder) func(*int64, *int64) int {
	return func(a, b *int64) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Int64(*a, *b)
	}
}

func Uint(a, b uint) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func UintPtr(order NilOrder) func(*uint, *uint) int {
	return func(a, b *uint) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Uint(*a, *b)
	}
}

func Uint8(a, b uint8) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Uint8Ptr(order NilOrder) func(*uint8, *uint8) int {
	return func(a, b *uint8) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Uint8(*a, *b)
	}
}

func Uint16(a, b uint16) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Uint16Ptr(order NilOrder) func(*uint16, *uint16) int {
	return func(a, b *uint16) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Uint16(*a, *b)
	}
}

func Uint32(a, b uint32) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Uint32Ptr(order NilOrder) func(*uint32, *uint32) int {
	return func(a, b *uint32) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Uint32(*a, *b)
	}
}

func Uint64(a, b uint64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Uint64Ptr(order NilOrder) func(*uint64, *uint64) int {
	return func(a, b *uint64) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Uint64(*a, *b)
	}
}

func Byte(a, b byte) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func BytePtr(order NilOrder) func(*byte, *byte) int {
	return func(a, b *byte) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Byte(*a, *b)
	}
}

func Rune(a, b rune) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func RunePtr(order NilOrder) func(*rune, *rune) int {
	return func(a, b *rune) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Rune(*a, *b)
	}
}

func Float32(a, b float32) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Float32Ptr(order NilOrder) func(*float32, *float32) int {
	return func(a, b *float32) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Float32(*a, *b)
	}
}

func Float64(a, b float64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}

func Float64Ptr(order NilOrder) func(*float64, *float64) int {
	return func(a, b *float64) int {
		v := order.compare(a == nil, b == nil)
		if v != 0 {
			return v
		} else if a == nil && b == nil {
			return 0
		}
		return Float64(*a, *b)
	}
}
