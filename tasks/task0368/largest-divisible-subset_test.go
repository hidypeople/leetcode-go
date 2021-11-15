package task0368

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestDivisibleSubset(t *testing.T) {
	assert.Equal(t, []int{2, 1}, largestDivisibleSubset([]int{1, 2, 3}))
}

func Test_largestDivisibleSubset2(t *testing.T) {
	assert.Equal(t, []int{8, 4, 2, 1}, largestDivisibleSubset([]int{1, 2, 4, 8}))
}

func Test_largestDivisibleSubset3(t *testing.T) {
	assert.Equal(t, []int{4, 2, 1}, largestDivisibleSubset([]int{1, 2, 3, 4, 5}))
}

func Test_largestDivisibleSubset4(t *testing.T) {
	assert.Equal(t, []int{720, 360, 180, 90, 18, 9}, largestDivisibleSubset(
		[]int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}))
}
