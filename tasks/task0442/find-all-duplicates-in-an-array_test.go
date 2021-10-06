package task0442

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDuplicates(t *testing.T) {
	assert.Equal(t, []int{2, 3}, findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func Test_findDuplicates2(t *testing.T) {
	assert.Equal(t, []int{1}, findDuplicates([]int{1, 1, 2}))
}

func Test_findDuplicates3(t *testing.T) {
	assert.Equal(t, []int{}, findDuplicates([]int{1}))
}
