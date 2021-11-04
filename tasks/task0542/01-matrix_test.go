package task0542

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_updateMatrix(t *testing.T) {
	assert.Equal(t, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}

func Test_updateMatrix2(t *testing.T) {
	assert.Equal(t, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}, updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}

func Test_updateMatrix3(t *testing.T) {
	input := [][]int{
		{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
		{1, 1, 1, 1, 0, 1, 0, 0, 1, 1}}
	expect := [][]int{
		{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 2, 1, 1, 0, 1},
		{2, 1, 1, 1, 1, 2, 1, 0, 1, 0},
		{3, 2, 2, 1, 0, 1, 0, 0, 1, 1}}
	/*actual := [][]int{
	{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
	{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
	{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
	{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
	{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
	{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
	{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
	{1, 0, 0, 0, 1, 2, 1, 1, 0, 1},
	{2, 1, 1, 1, 1, 2, 1, 0, 1, 0},
	{2, 2, 2, 1, 0, 1, 0, 0, 1, 1}}*/
	assert.Equal(t, expect, updateMatrix(input))
}

func Test_updateMatrix4(t *testing.T) {
	assert.Equal(t, [][]int{{1, 0, 0}, {2, 1, 1}, {3, 2, 2}}, updateMatrix([][]int{{1, 0, 0}, {1, 1, 1}, {1, 1, 1}}))
}
