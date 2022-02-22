package task0120

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumTotal(t *testing.T) {
	assert.Equal(t, 11, minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}

func Test_minimumTotal2(t *testing.T) {
	assert.Equal(t, -10, minimumTotal([][]int{{-10}}))
}

func Test_minimumTotal3(t *testing.T) {
	assert.Equal(t, -4, minimumTotal([][]int{{-1}, {-2, -3}}))
}

func Test_minimumTotal4(t *testing.T) {
	assert.Equal(t, -1, minimumTotal([][]int{{-1}, {3, 2}, {-3, 1, -1}}))
}
