package task0149

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPoints(t *testing.T) {
	assert.Equal(t, 3, maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
}

func Test_maxPoints2(t *testing.T) {
	assert.Equal(t, 4, maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
}
