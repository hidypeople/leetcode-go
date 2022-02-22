package task0137

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	assert.Equal(t, 3, singleNumber([]int{2, 2, 3, 2}))
}

func Test_singleNumber2(t *testing.T) {
	assert.Equal(t, 99, singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
}

func Test_singleNumber3(t *testing.T) {
	assert.Equal(t, -4, singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}

func Test_singleNumber2_1(t *testing.T) {
	assert.Equal(t, 99, singleNumber2([]int{0, 1, 0, 1, 0, 1, 99}))
}
