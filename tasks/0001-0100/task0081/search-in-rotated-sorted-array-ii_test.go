package task0081

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	assert.Equal(t, true, search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
}

func Test_search2(t *testing.T) {
	assert.Equal(t, false, search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}

func Test_searchStraight(t *testing.T) {
	assert.Equal(t, true, searchStraight([]int{1}, 0, 0, 1))
	assert.Equal(t, false, searchStraight([]int{2}, 0, 0, 1))
	assert.Equal(t, true, searchStraight([]int{1, 2, 3, 4}, 0, 3, 3))
	assert.Equal(t, true, searchStraight([]int{1, 2, 3, 4, 5, 6}, 0, 5, 1))
	assert.Equal(t, true, searchStraight([]int{1, 2, 3, 4, 5, 6}, 0, 5, 2))
	assert.Equal(t, true, searchStraight([]int{1, 2, 3, 4, 5, 6}, 0, 5, 3))
	assert.Equal(t, true, searchStraight([]int{1, 2, 3, 4, 5, 6}, 0, 5, 4))
	assert.Equal(t, true, searchStraight([]int{1, 2, 3, 4, 5, 6}, 0, 5, 5))
	assert.Equal(t, true, searchStraight([]int{1, 2, 3, 4, 5, 6}, 0, 5, 6))
	assert.Equal(t, false, searchStraight([]int{1, 2, 3, 4, 5, 6}, 0, 5, -1))
	assert.Equal(t, false, searchStraight([]int{1, 2, 3, 4, 5, 6}, 0, 5, 100))
	assert.Equal(t, false, searchStraight([]int{1, 2, 3, 4, 5}, 0, 4, -1))
	assert.Equal(t, false, searchStraight([]int{1, 2, 3, 4, 5}, 0, 4, 100))
	assert.Equal(t, false, searchStraight([]int{11, 22, 33, 44, 55}, 0, 4, 12))
	assert.Equal(t, false, searchStraight([]int{11, 22, 33, 44, 55}, 0, 4, 23))
	assert.Equal(t, false, searchStraight([]int{11, 22, 33, 44, 55}, 0, 4, 34))
	assert.Equal(t, false, searchStraight([]int{11, 22, 33, 44, 55}, 0, 4, 45))
}

func Test_search3(t *testing.T) {
	assert.Equal(t, true, search([]int{11, 22, 33, 44, 55}, 22))
	assert.Equal(t, true, search([]int{55, 11, 22, 33, 44}, 22))
	assert.Equal(t, true, search([]int{44, 55, 11, 22, 33}, 22))
	assert.Equal(t, true, search([]int{33, 44, 55, 11, 22}, 22))
	assert.Equal(t, true, search([]int{22, 33, 44, 55, 11}, 22))

	assert.Equal(t, false, search([]int{11, 22, 33, 44, 55}, 23))
	assert.Equal(t, false, search([]int{55, 11, 22, 33, 44}, 23))
	assert.Equal(t, false, search([]int{44, 55, 11, 22, 33}, 23))
	assert.Equal(t, false, search([]int{33, 44, 55, 11, 22}, 23))
	assert.Equal(t, false, search([]int{22, 33, 44, 55, 11}, 23))
}

func Test_search4(t *testing.T) {
	assert.Equal(t, true, search([]int{1, 0, 1, 1, 1}, 0))
}

func Test_search5(t *testing.T) {
	assert.Equal(t, false, search([]int{3, 1}, 0))
}
