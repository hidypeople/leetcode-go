package task0040

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum2(t *testing.T) {
	assert.Equal(t, [][]int{
		{6, 1, 1},
		{5, 2, 1},
		{7, 1},
		{6, 2},
	}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func Test_combinationSum22(t *testing.T) {
	assert.Equal(t, [][]int{
		{2, 2, 1},
		{5},
	}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
