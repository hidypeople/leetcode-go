package task0463

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_islandPerimeter(t *testing.T) {
	assert.Equal(t, 16, islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
}

func Test_islandPerimeter2(t *testing.T) {
	assert.Equal(t, 4, islandPerimeter([][]int{{1}}))
}

func Test_islandPerimeter3(t *testing.T) {
	assert.Equal(t, 8, islandPerimeter([][]int{{1, 1}, {1, 1}}))
}
