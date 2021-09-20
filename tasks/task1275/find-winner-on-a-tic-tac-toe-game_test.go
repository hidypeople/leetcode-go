package task1275

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tictactoe(t *testing.T) {
	assert.Equal(t, "A", tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}))
}

func Test_tictactoe2(t *testing.T) {
	assert.Equal(t, "B", tictactoe([][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}))
}

func Test_tictactoe3(t *testing.T) {
	assert.Equal(t, "Draw", tictactoe([][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}))
}

func Test_tictactoe4(t *testing.T) {
	assert.Equal(t, "Pending", tictactoe([][]int{{0, 0}, {1, 1}}))
}

func Test_tictactoe5(t *testing.T) {
	assert.Equal(t, "Pending", tictactoe([][]int{{0, 2}, {2, 0}, {2, 1}, {0, 1}, {1, 2}}))
}
