package task0011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	assert.Equal(t, 0, maxArea([]int{1}))
	assert.Equal(t, 1, maxArea([]int{1, 1}))
	assert.Equal(t, 16, maxArea([]int{4, 3, 2, 1, 4}))
	assert.Equal(t, 4, maxArea([]int{4, 3, 2}))
	assert.Equal(t, 2, maxArea([]int{1, 2, 1}))
}
