package task1734

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decode(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, decode([]int{3, 1}))
}

func Test_decode2(t *testing.T) {
	assert.Equal(t, []int{2, 4, 1, 5, 3}, decode([]int{6, 5, 4, 6}))
}
