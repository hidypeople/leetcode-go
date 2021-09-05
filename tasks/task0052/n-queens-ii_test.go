package task0052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_totalNQueens(t *testing.T) {
	assert.Equal(t, 2, totalNQueens(4))
}

func Test_totalNQueens2(t *testing.T) {
	assert.Equal(t, 1, totalNQueens(1))
}

func Test_totalNQueens3(t *testing.T) {
	assert.Equal(t, 0, totalNQueens(0))
}
