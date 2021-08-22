package task0037

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveSudoku(t *testing.T) {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(input)
	assert.Equal(t, [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}, input)
}

func Test_solveSudoku2(t *testing.T) {
	sudoku := [][]byte{
		{'.', '.', '3', '.', '1', '.', '.', '.', '6'},
		{'.', '1', '.', '.', '5', '8', '.', '.', '.'},
		{'.', '.', '.', '.', '9', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '4', '1'},
		{'.', '.', '.', '.', '.', '7', '5', '.', '.'},
		{'2', '.', '4', '.', '.', '.', '8', '.', '.'},
		{'7', '.', '.', '.', '4', '.', '.', '1', '3'},
		{'.', '9', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '.', '.', '.', '8', '.', '9', '7', '.'},
	}
	solveSudoku(sudoku)
	assert.Equal(t, [][]byte{
		{'9', '5', '3', '2', '1', '4', '7', '8', '6'},
		{'6', '1', '2', '7', '5', '8', '4', '3', '9'},
		{'8', '4', '7', '6', '9', '3', '1', '2', '5'},
		{'5', '7', '6', '8', '2', '9', '3', '4', '1'},
		{'1', '8', '9', '4', '3', '7', '5', '6', '2'},
		{'2', '3', '4', '5', '6', '1', '8', '9', '7'},
		{'7', '6', '8', '9', '4', '5', '2', '1', '3'},
		{'4', '9', '1', '3', '7', '2', '6', '5', '8'},
		{'3', '2', '5', '1', '8', '6', '9', '7', '4'},
	}, sudoku)
}