package task0039

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	result := combinationSum([]int{2, 3, 5}, 8)
	assert.Equal(t, [][]int{{2, 2, 2, 2}, {3, 3, 2}, {5, 3}}, result)
}

func Test_combinationSum2(t *testing.T) {
	result := combinationSum([]int{2}, 1)
	assert.Equal(t, [][]int{}, result)
}

func Test_combinationSum3(t *testing.T) {
	result := combinationSum([]int{1}, 1)
	assert.Equal(t, [][]int{{1}}, result)
}

func Test_combinationSum4(t *testing.T) {
	result := combinationSum([]int{2, 3, 6, 7}, 7)
	assert.Equal(t, [][]int{{3, 2, 2}, {7}}, result)
}

func Test_combinationSum5(t *testing.T) {
	result := combinationSum([]int{1}, 2)
	assert.Equal(t, [][]int{{1, 1}}, result)
}
