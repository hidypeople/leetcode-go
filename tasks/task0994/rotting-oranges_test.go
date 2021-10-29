package task0994

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_orangesRotting(t *testing.T) {
	input := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	assert.Equal(t, 4, orangesRotting(input))
}

func Test_orangesRotting2(t *testing.T) {
	input := [][]int{{0, 2}}
	assert.Equal(t, 0, orangesRotting(input))
}

func Test_orangesRotting3(t *testing.T) {
	input := [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1}}
	assert.Equal(t, -1, orangesRotting(input))
}
