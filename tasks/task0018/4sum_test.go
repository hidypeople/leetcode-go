package task0018

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fourSum(t *testing.T) {
	assert.Equal(t, [][]int{}, fourSum([]int{1, 0, -1}, 0))
}

func Test_fourSum2(t *testing.T) {
	assert.Equal(t, [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}, fourSum([]int{-2, -1, 0, 0, 1, 2}, 0))
}

// func Test_fourSum3(t *testing.T) {
// 	assert.Equal(t, [][]int{
// 		{-2, -2, 1, 2},
// 		{-2, -1, 1, 2},
// 		{-2, 0, 1, 1},
// 	}, fourSum([]int{-2, -2, -1, 0, 1, 1, 2, 2}, 0))
// }

func Test_fourSum4(t *testing.T) {
	assert.Equal(t, [][]int{{2, 2, 2, 2}}, fourSum([]int{2, 2, 2, 2}, 8))
}

func Test_fourSum7(t *testing.T) {
	assert.Equal(t, [][]int{
		{2, 2, 2, 2},
	}, fourSum([]int{2, 2, 2, 2, 2}, 8))
}

func Test_fourSum5(t *testing.T) {
	assert.Equal(t, [][]int{}, fourSum([]int{-2, 2, 2, 2, 2, 3}, 0))
}

func Test_fourSum6(t *testing.T) {
	assert.Equal(t, [][]int{
		{-2, -1, 1, 2},
		{-1, -1, 1, 1},
	}, fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
}

func Test_fourSum8(t *testing.T) {
	assert.Equal(t, [][]int{
		{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1},
	}, fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
