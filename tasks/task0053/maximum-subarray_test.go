package task0053

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func Test_maxSubArray2(t *testing.T) {
	assert.Equal(t, 23, maxSubArray([]int{5, 4, -1, 7, 8}))
}

func Test_maxSubArray3(t *testing.T) {
	assert.Equal(t, 1, maxSubArray([]int{1}))
}

func Test_maxSubArray4(t *testing.T) {
	assert.Equal(t, -2, maxSubArray([]int{-100, -2, -2, -100}))
}
