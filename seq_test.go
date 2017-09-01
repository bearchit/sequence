package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ds = NewZbSeq(0, 2)
)

func TestZbSeq_Contains(t *testing.T) {
	assert.True(t, ds.Contains(4))
	assert.False(t, ds.Contains(3))
	assert.True(t, ds.Contains(-2))
	assert.True(t, ds.Contains(0))
}

func TestZbSeq_Value(t *testing.T) {
	assert.Equal(t, 4, ds.Value(2))
}

func TestZbSeq_ToArray(t *testing.T) {
	assert.Equal(t, []int{0, 2, 4, 6, 8}, ds.ToArray(5))

	s := NewZbSeq(-2, 2)
	assert.Equal(t, []int{0, 2, 4, 6, 8}, s.ToArray(5))
}
