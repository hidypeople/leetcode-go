package task1510

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_winnerSquareGame(t *testing.T) {
	assert.True(t, winnerSquareGame(1))
}

func Test_winnerSquareGame2(t *testing.T) {
	assert.False(t, winnerSquareGame(2))
}

func Test_winnerSquareGame3(t *testing.T) {
	assert.True(t, winnerSquareGame(4))
}

func Test_winnerSquareGame4(t *testing.T) {
	assert.False(t, winnerSquareGame(17))
}
