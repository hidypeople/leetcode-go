package task0941

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validMountainArray(t *testing.T) {
	assert.False(t, validMountainArray([]int{2, 1}))
}

func Test_validMountainArray2(t *testing.T) {
	assert.False(t, validMountainArray([]int{3, 5, 5}))
}

func Test_validMountainArray3(t *testing.T) {
	assert.True(t, validMountainArray([]int{0, 3, 2, 1}))
}

func Test_validMountainArray4(t *testing.T) {
	assert.False(t, validMountainArray([]int{0, 3, 3, 4, 6, 2, 1}))
}

func Test_validMountainArray5(t *testing.T) {
	assert.False(t, validMountainArray([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
}

func Test_validMountainArray6(t *testing.T) {
	assert.False(t, validMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7}))
}

func Test_validMountainArray7(t *testing.T) {
	assert.False(t, validMountainArray([]int{1, 1, 1, 1, 1, 1, 1, 2, 1}))
}
