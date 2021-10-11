package task0015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	assert.Nil(t, threeSum([]int{}))
	assert.Equal(t, [][]int{}, threeSum([]int{1, 2, 3}))
	assert.Equal(t, [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
