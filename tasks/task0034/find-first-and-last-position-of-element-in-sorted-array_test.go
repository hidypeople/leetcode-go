package task0034

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	assert.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func Test_searchRange2(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func Test_searchRange3(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, searchRange([]int{}, 0))
}

func Test_searchRange4(t *testing.T) {
	assert.Equal(t, []int{3, 8}, searchRange([]int{5, 7, 7, 8, 8, 8, 8, 8, 8, 10}, 8))
}

func Test_searchRange5(t *testing.T) {
	assert.Equal(t, []int{4, 4}, searchRange([]int{1, 1, 1, 1, 8, 9, 9, 9, 9}, 8))
}

func Test_searchRange6(t *testing.T) {
	assert.Equal(t, []int{4, 5}, searchRange([]int{1, 1, 1, 1, 8, 8, 9, 9, 9}, 8))
}

func Test_searchRange7(t *testing.T) {
	assert.Equal(t, []int{3, 4}, searchRange([]int{1, 1, 1, 8, 8, 9, 9, 9, 9}, 8))
}

func Test_searchRange8(t *testing.T) {
	assert.Equal(t, []int{4, 7}, searchRange([]int{1, 1, 1, 1, 8, 8, 8, 8, 9}, 8))
}

func Test_searchRange9(t *testing.T) {
	assert.Equal(t, []int{1, 4}, searchRange([]int{1, 8, 8, 8, 8, 9, 9, 9, 9}, 8))
}

func Test_searchRange10(t *testing.T) {
	assert.Equal(t, []int{4, 8}, searchRange([]int{1, 1, 1, 1, 8, 8, 8, 8, 8}, 8))
}

func Test_searchRange11(t *testing.T) {
	assert.Equal(t, []int{0, 4}, searchRange([]int{8, 8, 8, 8, 8, 9, 9, 9, 9}, 8))
}

func Test_searchRange12(t *testing.T) {
	assert.Equal(t, []int{0, 8}, searchRange([]int{8, 8, 8, 8, 8, 8, 8, 8, 8}, 8))
}

func Test_searchRange13(t *testing.T) {
	assert.Equal(t, []int{0, 0}, searchRange([]int{1}, 1))
}

func Test_searchRange14(t *testing.T) {
	assert.Equal(t, []int{1, 1}, searchRange([]int{1, 4}, 4))
}

func Test_searchRange15(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, searchRange([]int{2, 2}, 3))
}

func Test_searchRange16(t *testing.T) {
	assert.Equal(t, []int{-1, -1}, searchRange([]int{2, 2}, 1))
}
