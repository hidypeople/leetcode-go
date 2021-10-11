package task0035

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert(t *testing.T) {
	assert.Equal(t, 1, searchInsert([]int{1, 3, 5, 6}, 2))
}

func Test_searchInsert2(t *testing.T) {
	assert.Equal(t, 4, searchInsert([]int{1, 3, 5, 6}, 7))
}

func Test_searchInsert3(t *testing.T) {
	assert.Equal(t, 0, searchInsert([]int{1, 3, 5, 6}, 0))
}

func Test_searchInsert4(t *testing.T) {
	assert.Equal(t, 0, searchInsert([]int{1}, 0))
}

func Test_searchInsert5(t *testing.T) {
	assert.Equal(t, 0, searchInsert([]int{}, 1))
}

func Test_searchInsert6(t *testing.T) {
	assert.Equal(t, 2, searchInsert([]int{1, 2, 3, 4, 5}, 3))
}

func Test_searchInsert7(t *testing.T) {
	assert.Equal(t, 3, searchInsert([]int{3, 5, 7, 9, 10}, 8))
}
