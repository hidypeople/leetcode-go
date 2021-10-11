package task0054

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_spiralOrder(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func Test_spiralOrder2(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4, 6, 8, 7, 5, 3},
		spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
}

func Test_spiralOrder3(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4, 3},
		spiralOrder([][]int{{1, 2}, {3, 4}}))
}
