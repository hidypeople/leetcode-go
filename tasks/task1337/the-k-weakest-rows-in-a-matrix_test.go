package task1337

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kWeakestRows(t *testing.T) {
	assert.Equal(t, []int{2, 0, 3}, kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}},
		3))
}

func Test_kWeakestRows2(t *testing.T) {
	assert.Equal(t, []int{0, 2}, kWeakestRows([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0}},
		2))
}

func Test_kWeakestRows3(t *testing.T) {
	assert.Equal(t, []int{0}, kWeakestRows([][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1}},
		1))
}
