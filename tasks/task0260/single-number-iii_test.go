package task0260

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	assert.Equal(t, []int{3, 5}, singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

func Test_singleNumber2(t *testing.T) {
	assert.Equal(t, []int{-1, 0}, singleNumber([]int{-1, 0}))
}

func Test_singleNumber3(t *testing.T) {
	assert.Equal(t, []int{999, 300}, singleNumber([]int{1, 1, 2, 2, 1000, 1000, 999, 300, 800, 800}))
}
