package task0210

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findOrder(t *testing.T) {
	assert.Equal(t, []int{0, 1}, findOrder(2, [][]int{{1, 0}}))
}

func Test_findOrder2(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3}, findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

func Test_findOrder3(t *testing.T) {
	assert.Equal(t, []int{0}, findOrder(1, [][]int{}))
}

func Test_findOrder4(t *testing.T) {
	assert.Equal(t, []int{0, 1}, findOrder(2, [][]int{}))
}

func Test_findOrder5(t *testing.T) {
	assert.Equal(t, []int{1, 0}, findOrder(2, [][]int{{0, 1}}))
}

func Test_findOrder6(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2}, findOrder(3, [][]int{{2, 0}, {2, 1}}))
}
