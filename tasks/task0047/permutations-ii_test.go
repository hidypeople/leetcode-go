package task0047

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permuteUnique(t *testing.T) {
	assert.Equal(t, [][]int{{1}}, permuteUnique([]int{1}))
}

func Test_permuteUnique2(t *testing.T) {
	assert.Equal(t, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}, permuteUnique([]int{1, 1, 2}))
}

func Test_permuteUnique3(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permuteUnique([]int{1, 2, 3}))
}
