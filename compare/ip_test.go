package compare

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIP(t *testing.T) {
	to := func(s string) net.IP {
		return net.ParseIP(s)
	}
	cases := []struct {
		V    int
		A, B net.IP
	}{
		{0, nil, nil},
		{0, to("10.10.0.1"), to("10.10.0.1")},
		{-1, nil, to("10.10.0.1")},
		{1, to("10.10.0.1"), nil},
		{-1, to("10.10.1.10"), to("10.10.100.10")},
		{1, to("10.10.1.100"), to("10.10.1.20")},
	}
	for _, c := range cases {
		assert.Equalf(t, c.V, IP(c.A, c.B), "%v <=> %v", c.A, c.B)
	}
}

func TestIPPtr(t *testing.T) {
	to := func(s string) *net.IP {
		v := net.ParseIP(s)
		return &v
	}
	cases := []struct {
		V    int
		A, B *net.IP
	}{
		{0, nil, nil},
		{0, to("10.10.0.1"), to("10.10.0.1")},
		{-1, nil, to("10.10.0.1")},
		{1, to("10.10.0.1"), nil},
		{-1, to("10.10.1.10"), to("10.10.100.10")},
		{1, to("10.10.1.100"), to("10.10.1.20")},
	}
	for _, c := range cases {
		assert.NotPanicsf(t, func() {
			assert.Equalf(t, c.V, IPPtr(NilFirst)(c.A, c.B), "%v <=> %v", c.A, c.B)
		}, "%v <=> %v", c.A, c.B)
	}
}

func TestIPNet(t *testing.T) {
	to := func(s string) net.IPNet {
		_, v, err := net.ParseCIDR(s)
		if err != nil {
			panic(err)
		}
		return *v
	}
	cases := []struct {
		V    int
		A, B net.IPNet
	}{
		{0, to("10.10.0.0/16"), to("10.10.0.0/16")},
		{-1, to("10.10.0.0/17"), to("10.10.0.0/16")},
		{1, to("10.10.0.0/16"), to("10.10.0.0/17")},
	}
	for _, c := range cases {
		assert.Equalf(t, c.V, IPNet(c.A, c.B), "%v <=> %v", c.A, c.B)
	}
}

func TestIPNetPtr(t *testing.T) {
	to := func(s string) *net.IPNet {
		_, v, err := net.ParseCIDR(s)
		if err != nil {
			panic(err)
		}
		return v
	}
	cases := []struct {
		V    int
		A, B *net.IPNet
	}{
		{0, nil, nil},
		{0, to("10.10.0.0/16"), to("10.10.0.0/16")},
		{-1, to("10.10.0.0/17"), to("10.10.0.0/16")},
		{1, to("10.10.0.0/16"), to("10.10.0.0/17")},
	}
	for _, c := range cases {
		assert.NotPanicsf(t, func() {
			assert.Equalf(t, c.V, IPNetPtr(NilFirst)(c.A, c.B), "%v <=> %v", c.A, c.B)
		}, "%v <=> %v", c.A, c.B)
	}
}
