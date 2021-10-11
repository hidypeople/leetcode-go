package task0066

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	assert.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	assert.Equal(t, []int{1}, plusOne([]int{0}))
	assert.Equal(t, []int{1, 0, 0, 0}, plusOne([]int{9, 9, 9}))
}
