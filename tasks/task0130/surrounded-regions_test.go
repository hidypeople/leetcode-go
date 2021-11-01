package task0130

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	input := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(input)
	assert.Equal(t, [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}, input)
}

func Test_solve2(t *testing.T) {
	input := [][]byte{{'X'}}
	solve(input)
	assert.Equal(t, [][]byte{{'X'}}, input)
}
