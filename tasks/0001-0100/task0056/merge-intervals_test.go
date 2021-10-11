package task0056

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	assert.Equal(t,
		[][]int{{1, 6}, {8, 10}, {15, 18}},
		merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}),
	)
}

func Test_merge2(t *testing.T) {
	assert.Equal(t,
		[][]int{{1, 5}},
		merge([][]int{{1, 4}, {4, 5}}),
	)
}

func Test_merge3(t *testing.T) {
	assert.Equal(t,
		[][]int{{0, 5}},
		merge([][]int{{1, 4}, {0, 5}}),
	)
}

func Test_merge4(t *testing.T) {
	assert.Equal(t,
		[][]int{{1, 10}},
		merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}),
	)
}
