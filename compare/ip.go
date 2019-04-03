package compare

import (
	"bytes"
	"net"
)

func IP(a, b net.IP) int {
	return bytes.Compare(a.To16(), b.To16())
}

func IPPtr(order NilOrder) func(*net.IP, *net.IP) int {
	return func(a, b *net.IP) int {
		v := order.compare(a, b)
		if v != 0 {
			return v
		}
		return IP(*a, *b)
	}
}

func IPNet(a, b net.IPNet) int {
	v := bytes.Compare(a.Mask, b.Mask)
	if v != 0 {
		return v
	}
	return IP(a.IP, b.IP)
}

func IPNetPtr(order NilOrder) func(*net.IPNet, *net.IPNet) int {
	return func(a, b *net.IPNet) int {
		v := order.compare(a, b)
		if v != 0 {
			return v
		}
		return IPNet(*a, *b)
	}
}
