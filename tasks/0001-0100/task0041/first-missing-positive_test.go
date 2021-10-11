package task0041

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstMissingPositive(t *testing.T) {
	assert.Equal(t, 3, firstMissingPositive([]int{1, 2, 0}))
}

func Test_firstMissingPositive2(t *testing.T) {
	assert.Equal(t, 2, firstMissingPositive([]int{3, 4, -1, 1}))
}

func Test_firstMissingPositive3(t *testing.T) {
	assert.Equal(t, 1, firstMissingPositive([]int{7, 8, 9, 11, 12}))
}

func Test_firstMissingPositive4(t *testing.T) {
	assert.Equal(t, 2, firstMissingPositive([]int{1, 1000}))
}
