package task0030

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermutation0(t *testing.T) {
	array := []int{1, 2}
	nextPermutation(array)
	assert.Equal(t, []int{2, 1}, array)
}

func Test_nextPermutation(t *testing.T) {
	array := []int{1, 2, 3}
	nextPermutation(array)
	assert.Equal(t, []int{1, 3, 2}, array)
}

func Test_nextPermutation2(t *testing.T) {
	array := []int{1, 1, 5}
	nextPermutation(array)
	assert.Equal(t, []int{1, 5, 1}, array)
}

func Test_nextPermutation3(t *testing.T) {
	array := []int{1}
	nextPermutation(array)
	assert.Equal(t, []int{1}, array)
}

func Test_nextPermutation4(t *testing.T) {
	array := []int{0, 1, 2, 5, 3, 3, 0}
	nextPermutation(array)
	assert.Equal(t, []int{0, 1, 3, 0, 2, 3, 5}, array)
}

func Test_nextPermutation5(t *testing.T) {
	array := []int{3, 2, 1, 0}
	nextPermutation(array)
	assert.Equal(t, []int{0, 1, 2, 3}, array)
}
