package task0051

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveNQueens(t *testing.T) {
	assert.Equal(t, [][]string{{"Q"}}, solveNQueens(1))
}

func Test_solveNQueens2(t *testing.T) {
	assert.Equal(t,
		[][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		solveNQueens(4),
	)
}

func Test_solveNQueens3(t *testing.T) {
	assert.Equal(t,
		[][]string{{"Q....", "..Q..", "....Q", ".Q...", "...Q."}, {"Q....", "...Q.", ".Q...", "....Q", "..Q.."}, {".Q...", "...Q.", "Q....", "..Q..", "....Q"}, {".Q...", "....Q", "..Q..", "Q....", "...Q."}, {"..Q..", "Q....", "...Q.", ".Q...", "....Q"}, {"..Q..", "....Q", ".Q...", "...Q.", "Q...."}, {"...Q.", "Q....", "..Q..", "....Q", ".Q..."}, {"...Q.", ".Q...", "....Q", "..Q..", "Q...."}, {"....Q", ".Q...", "...Q.", "Q....", "..Q.."}, {"....Q", "..Q..", "Q....", "...Q.", ".Q..."}},
		solveNQueens(5),
	)
}

func Test_solveNQueens4(t *testing.T) {
	assert.Equal(t,
		[][]string{},
		solveNQueens(2),
	)
}
