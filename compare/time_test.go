package compare

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	to := func(s string) time.Time {
		v, err := time.Parse(time.RFC3339, s)
		if err != nil {
			panic(err)
		}
		return v
	}
	cases := []struct {
		V    int
		A, B time.Time
	}{
		{0, to("2019-04-03T01:02:03Z"), to("2019-04-03T01:02:03Z")},
		{-1, to("2019-04-03T01:02:03Z"), to("2019-04-03T01:02:04Z")},
		{1, to("2019-04-03T01:02:04Z"), to("2019-04-03T01:02:03Z")},
	}
	for _, c := range cases {
		assert.Equalf(t, c.V, Time(c.A, c.B), "%v <=> %v", c.A, c.B)
	}
}

func TestTimePtr(t *testing.T) {
	to := func(s string) *time.Time {
		v, err := time.Parse(time.RFC3339, s)
		if err != nil {
			panic(err)
		}
		return &v
	}
	cases := []struct {
		V    int
		A, B *time.Time
	}{
		{0, nil, nil},
		{0, to("2019-04-03T01:02:03Z"), to("2019-04-03T01:02:03Z")},
		{-1, to("2019-04-03T01:02:03Z"), to("2019-04-03T01:02:04Z")},
		{1, to("2019-04-03T01:02:04Z"), to("2019-04-03T01:02:03Z")},
	}
	for _, c := range cases {
		assert.NotPanicsf(t, func() {
			assert.Equalf(t, c.V, TimePtr(NilFirst)(c.A, c.B), "%v <=> %v", c.A, c.B)
		}, "%v <=> %v", c.A, c.B)
	}
}
