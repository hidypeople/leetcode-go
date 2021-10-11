package task0080

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	assert.Equal(t, 5, removeDuplicates(nums))
	assert.Equal(t, []int{1, 1, 2, 2, 3}, nums[:len(nums)-1])
}

func Test_removeDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	assert.Equal(t, 7, removeDuplicates(nums))
	assert.Equal(t, []int{0, 0, 1, 1, 2, 3, 3}, nums[:len(nums)-2])
}

func Test_removeDuplicates3(t *testing.T) {
	nums := []int{1}
	assert.Equal(t, 1, removeDuplicates(nums))
	assert.Equal(t, []int{1}, nums)
}

func Test_removeDuplicates4(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	assert.Equal(t, 2, removeDuplicates(nums))
	assert.Equal(t, []int{1, 1}, nums[:len(nums)-3])
}
