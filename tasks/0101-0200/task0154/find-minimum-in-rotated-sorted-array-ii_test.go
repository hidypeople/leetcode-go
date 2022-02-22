package task0154

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMin(t *testing.T) {
	assert.Equal(t, 1, findMin([]int{1, 3, 5}))
}

func Test_findMin2(t *testing.T) {
	assert.Equal(t, 0, findMin([]int{2, 2, 2, 0, 1}))
}
