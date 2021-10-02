package task0174

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateMinimumHP(t *testing.T) {
	maze := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5}}
	assert.Equal(t, 7, calculateMinimumHP(maze))
}

func Test_calculateMinimumHP2(t *testing.T) {
	maze := [][]int{{0}}
	assert.Equal(t, 1, calculateMinimumHP(maze))
}

func Test_calculateMinimumHP3(t *testing.T) {
	maze := [][]int{
		{0, 0},
		{-5, -4}}
	assert.Equal(t, 5, calculateMinimumHP(maze))
}

func Test_calculateMinimumHP4(t *testing.T) {
	maze := [][]int{
		{0, 0}}
	assert.Equal(t, 1, calculateMinimumHP(maze))
}
