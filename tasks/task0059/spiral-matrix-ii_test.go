package task0059

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateMatrix(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, generateMatrix(3))
}

func Test_generateMatrix2(t *testing.T) {
	assert.Equal(t, [][]int{{1}}, generateMatrix(1))
}
