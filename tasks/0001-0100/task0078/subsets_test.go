package task0078

import (
	"testing"

	"github.com/stretchr/testify/assert"

	combination "leetcode/utils/combination"
)

func Test_subsets(t *testing.T) {
	assert.True(t, combination.IsEqualCombinations(
		[][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		subsets([]int{1, 2, 3}),
		true))
}

func Test_subsets2(t *testing.T) {
	assert.True(t, combination.IsEqualCombinations(
		[][]int{{}, {1}},
		subsets([]int{1}),
		true))
}
