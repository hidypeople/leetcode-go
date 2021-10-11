package task0077

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combine(t *testing.T) {
	assert.Equal(t, [][]int{{1}}, combine(1, 1))
}

func Test_combine2(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}, combine(4, 2))
}

func Test_combine3(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}, combine(4, 3))
}
