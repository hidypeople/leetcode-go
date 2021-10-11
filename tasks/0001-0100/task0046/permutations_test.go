package task0046

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permute(t *testing.T) {
	assert.Equal(t, [][]int{{1}}, permute([]int{1}))
}

func Test_permute2(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {2, 1}}, permute([]int{1, 2}))
}

func Test_permute3(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute([]int{1, 2, 3}))
}
