package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	cases := []struct {
		V    int
		A, B bool
	}{
		{0, true, true},
		{0, false, false},
		{-1, true, false},
		{1, false, true},
	}
	for _, c := range cases {
		assert.Equalf(t, c.V, Bool(c.A, c.B), "%v <=> %v", c.A, c.B)
	}
}

func TestBoolPtr(t *testing.T) {
	to := func(v bool) *bool {
		return &v
	}
	cases := []struct {
		V    int
		A, B *bool
	}{
		{0, nil, nil},
		{0, to(true), to(true)},
		{0, to(false), to(false)},
		{-1, to(true), to(false)},
		{1, to(false), to(true)},
	}
	for _, c := range cases {
		assert.NotPanicsf(t, func() {
			assert.Equalf(t, c.V, BoolPtr(NilFirst)(c.A, c.B), "%v <=> %v", c.A, c.B)
		}, "%v <=> %v", c.A, c.B)
	}
}
