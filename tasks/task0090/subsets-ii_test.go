package task0090

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	combination "leetcode/utils/combination"
)

func Test_subsetsWithDup(t *testing.T) {
	assert.Equal(t, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}, subsetsWithDup([]int{1, 2, 2}))
}

func Test_subsetsWithDup2(t *testing.T) {
	assert.Equal(t, [][]int{{}, {1}}, subsetsWithDup([]int{1}))
}

func Test_subsetsWithDup3(t *testing.T) {
	assert.True(t, combination.IsEqualCombinations(
		[][]int{{}, {-10}, {-10, 10}, {10}},
		subsetsWithDup([]int{-10, 10}),
		true,
	))

}

func Test_subsetsWithDup4(t *testing.T) {
	fmt.Printf("%v\n", subsetsWithDup([]int{4, 4, 4, 1, 4}))
	assert.True(t, combination.IsEqualCombinations(
		[][]int{{}, {1}, {1, 4}, {1, 4, 4}, {1, 4, 4, 4}, {1, 4, 4, 4, 4}, {4}, {4, 4}, {4, 4, 4}, {4, 4, 4, 4}},
		subsetsWithDup([]int{4, 4, 4, 1, 4}),
		true,
	))
}

func Test_subsetsWithDup5(t *testing.T) {
	assert.Equal(t, [][]int{{}, {1}, {1, 1}, {1, 1, 2}, {1, 2}, {2}}, subsetsWithDup([]int{1, 2, 1}))
}
