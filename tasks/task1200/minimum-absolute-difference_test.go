package task1200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumAbsDifference(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {2, 3}, {3, 4}}, minimumAbsDifference([]int{4, 2, 1, 3}))
}

func Test_minimumAbsDifference2(t *testing.T) {
	assert.Equal(t, [][]int{{1, 3}}, minimumAbsDifference([]int{1, 3, 6, 10, 15}))
}

func Test_minimumAbsDifference3(t *testing.T) {
	assert.Equal(t, [][]int{{-14, -10}, {19, 23}, {23, 27}}, minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}

func Test_minimumAbsDifference4(t *testing.T) {
	assert.Equal(t, [][]int{{26, 27}}, minimumAbsDifference([]int{40, 11, 26, 27, -20}))
}
