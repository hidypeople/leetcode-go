package task0073

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setZeroes(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	assert.Equal(t, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, matrix)
}

func Test_setZeroes2(t *testing.T) {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	assert.Equal(t, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}, matrix)
}

func Test_setZeroes3(t *testing.T) {
	matrix := [][]int{{0, 0, 0, 5}, {4, 3, 1, 4}, {0, 1, 1, 4}, {1, 2, 1, 3}, {0, 0, 1, 1}}
	setZeroes(matrix)
	assert.Equal(t, [][]int{{0, 0, 0, 0}, {0, 0, 0, 4}, {0, 0, 0, 0}, {0, 0, 0, 3}, {0, 0, 0, 0}}, matrix)
}

func Test_setZeroes4(t *testing.T) {
	matrix := [][]int{{1, 0, 3}}
	setZeroes(matrix)
	assert.Equal(t, [][]int{{0, 0, 0}}, matrix)
}
