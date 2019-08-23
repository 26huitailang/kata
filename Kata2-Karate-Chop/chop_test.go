package chop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, -1, Chop(3, []int{}))
	assert.Equal(t, -1, Chop(3, []int{1}))
	assert.Equal(t, 0, Chop(1, []int{1}))

	assert.Equal(t, 0, Chop(1, []int{1, 3, 5}))
	assert.Equal(t, 1, Chop(3, []int{1, 3, 5}))
	assert.Equal(t, 2, Chop(5, []int{1, 3, 5}))
	assert.Equal(t, -1, Chop(0, []int{1, 3, 5}))
	assert.Equal(t, -1, Chop(2, []int{1, 3, 5}))
	assert.Equal(t, -1, Chop(4, []int{1, 3, 5}))
	assert.Equal(t, -1, Chop(6, []int{1, 3, 5}))

	assert.Equal(t, 0, Chop(1, []int{1, 3, 5, 7}))
	assert.Equal(t, 1, Chop(3, []int{1, 3, 5, 7}))
	assert.Equal(t, 2, Chop(5, []int{1, 3, 5, 7}))
	assert.Equal(t, 3, Chop(7, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, Chop(0, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, Chop(2, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, Chop(4, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, Chop(6, []int{1, 3, 5, 7}))
	assert.Equal(t, -1, Chop(8, []int{1, 3, 5, 7}))
}
