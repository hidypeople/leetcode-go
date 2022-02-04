package task0525

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxLength(t *testing.T) {
	assert.Equal(t, 2, findMaxLength([]int{0, 1}))
}

func Test_findMaxLength2(t *testing.T) {
	assert.Equal(t, 2, findMaxLength([]int{0, 1, 0}))
}

func Test_findMaxLength3(t *testing.T) {
	assert.Equal(t, 8, findMaxLength([]int{0, 0, 1, 0, 1, 0, 1, 1, 1}))
}
