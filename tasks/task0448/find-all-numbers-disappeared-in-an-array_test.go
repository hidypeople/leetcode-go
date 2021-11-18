package task0448

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDisappearedNumbers(t *testing.T) {
	assert.Equal(t, []int{5, 6}, findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func Test_findDisappearedNumbers2(t *testing.T) {
	assert.Equal(t, []int{2}, findDisappearedNumbers([]int{1, 1}))
}
