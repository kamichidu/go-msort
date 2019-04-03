package compare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilOrder(t *testing.T) {
	i := 1
	var n *int
	e := &i

	assert.Equal(t, -1, NilFirst.compare(n, n))
	assert.Equal(t, 1, NilLast.compare(n, n))

	assert.Equal(t, 0, NilFirst.compare(e, e))
	assert.Equal(t, 0, NilLast.compare(e, e))

	assert.Equal(t, -1, NilFirst.compare(n, e))
	assert.Equal(t, -1, NilLast.compare(e, n))

	assert.Equal(t, 1, NilFirst.compare(e, n))
	assert.Equal(t, 1, NilLast.compare(n, e))
}
