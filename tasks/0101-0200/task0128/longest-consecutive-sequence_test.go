package task0128

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestConsecutive(t *testing.T) {
	assert.Equal(t, 4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func Test_longestConsecutive2(t *testing.T) {
	assert.Equal(t, 9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

func Test_longestConsecutive3(t *testing.T) {
	assert.Equal(t, 5, longestConsecutive([]int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3}))
}

func Test_longestConsecutive4(t *testing.T) {
	assert.Equal(t, 5, longestConsecutive([]int{-8, -4, -3, -2, -1, 0, 2, 4, 5, 6, 7}))
	assert.Equal(t, 5, longestConsecutive([]int{7, 6, 5, 4, 2, 0, -1, -2, -3, -4, -8}))
}
