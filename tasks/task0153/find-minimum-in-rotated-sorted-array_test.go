package task0153

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMin(t *testing.T) {
	assert.Equal(t, 1, findMin([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 1, findMin([]int{2, 3, 4, 5, 1}))
	assert.Equal(t, 1, findMin([]int{3, 4, 5, 1, 2}))
	assert.Equal(t, 1, findMin([]int{4, 5, 1, 2, 3}))
	assert.Equal(t, 1, findMin([]int{5, 1, 2, 3, 4}))
}

func Test_findMin2(t *testing.T) {
	assert.Equal(t, 0, findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func Test_findMin3(t *testing.T) {
	assert.Equal(t, 11, findMin([]int{11, 13, 15, 17}))
}
