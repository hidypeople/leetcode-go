package task0782

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_movesToChessboard(t *testing.T) {
	assert.Equal(t, 2, movesToChessboard([][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {1, 0, 0, 1}, {1, 0, 0, 1}}))
}

func Test_movesToChessboard2(t *testing.T) {
	assert.Equal(t, 0, movesToChessboard([][]int{{0, 1}, {1, 0}}))
}

func Test_movesToChessboard3(t *testing.T) {
	assert.Equal(t, -1, movesToChessboard([][]int{{1, 0}, {1, 0}}))
}
