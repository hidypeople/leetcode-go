package task0063

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	assert.Equal(t, 2, uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}

func Test_uniquePathsWithObstacles2(t *testing.T) {
	assert.Equal(t, 1, uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}

func Test_uniquePathsWithObstacles3(t *testing.T) {
	assert.Equal(t, 0, uniquePathsWithObstacles([][]int{{0, 0}, {0, 1}}))
}

func Test_uniquePathsWithObstacles4(t *testing.T) {
	assert.Equal(t, 0, uniquePathsWithObstacles([][]int{{1, 0}, {0, 0}}))
}
