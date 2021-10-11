package task0048

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	input := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	rotate(input)
	assert.Equal(t, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}, input)
}

func Test_rotate2(t *testing.T) {
	input := [][]int{
		{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16},
	}
	rotate(input)
	assert.Equal(t, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}, input)
}

func Test_rotate3(t *testing.T) {
	input := [][]int{
		{1},
	}
	rotate(input)
	assert.Equal(t, [][]int{{1}}, input)
}

func Test_rotate4(t *testing.T) {
	input := [][]int{
		{1, 2}, {3, 4},
	}
	rotate(input)
	assert.Equal(t, [][]int{{3, 1}, {4, 2}}, input)
}
