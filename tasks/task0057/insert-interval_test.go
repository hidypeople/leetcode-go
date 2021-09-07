package task0057

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert(t *testing.T) {
	assert.Equal(t, [][]int{{1, 5}, {6, 9}}, insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}

func Test_insert2(t *testing.T) {
	assert.Equal(t, [][]int{{1, 5}}, insert([][]int{{1, 5}}, []int{2, 3}))
}

func Test_insert3(t *testing.T) {
	assert.Equal(t, [][]int{{2, 3}}, insert([][]int{}, []int{2, 3}))
}

func Test_insert4(t *testing.T) {
	assert.Equal(t, [][]int{{1, 7}}, insert([][]int{{1, 5}}, []int{2, 7}))
}

func Test_insert5(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {10, 11}}, insert([][]int{{1, 2}, {3, 4}}, []int{10, 11}))
}

func Test_insert6(t *testing.T) {
	assert.Equal(t, [][]int{{-11, -10}, {1, 2}, {3, 4}}, insert([][]int{{1, 2}, {3, 4}}, []int{-11, -10}))
}

func Test_insert7(t *testing.T) {
	assert.Equal(t, [][]int{{0, 50}}, insert([][]int{{1, 3}, {6, 9}}, []int{0, 50}))
}

func Test_insert8(t *testing.T) {
	assert.Equal(t, [][]int{{3, 5}, {6, 6}, {12, 15}}, insert([][]int{{3, 5}, {12, 15}}, []int{6, 6}))
}
