package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	cases := []struct {
		V    int
		A, B string
	}{
		{0, "", ""},
		{0, "000", "000"},
		{-1, "000", "010"},
		{1, "010", "001"},
	}
	for _, c := range cases {
		assert.Equalf(t, c.V, String(c.A, c.B), "%v <=> %v", c.A, c.B)
	}
}

func TestStringPtr(t *testing.T) {
	to := func(v string) *string {
		return &v
	}
	cases := []struct {
		V    int
		A, B *string
	}{
		{0, nil, nil},
		{0, to(""), to("")},
		{0, to("000"), to("000")},
		{-1, to("000"), to("010")},
		{1, to("010"), to("001")},
	}
	for _, c := range cases {
		assert.NotPanicsf(t, func() {
			assert.Equalf(t, c.V, StringPtr(NilFirst)(c.A, c.B), "%v <=> %v", c.A, c.B)
		}, "%v <=> %v", c.A, c.B)
	}
}
