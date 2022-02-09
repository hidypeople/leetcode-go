package task0532

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPairs(t *testing.T) {
	assert.Equal(t, 2, findPairs([]int{3, 1, 4, 1, 5}, 2))
}

func Test_findPairs2(t *testing.T) {
	assert.Equal(t, 4, findPairs([]int{1, 2, 3, 4, 5}, 1))
}

func Test_findPairs3(t *testing.T) {
	assert.Equal(t, 1, findPairs([]int{1, 3, 1, 5, 4}, 0))
}

func Test_findPairs4(t *testing.T) {
	assert.Equal(t, 2, findPairs([]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, 3))
}
