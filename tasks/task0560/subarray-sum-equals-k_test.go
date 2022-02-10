package task0560

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subarraySum(t *testing.T) {
	assert.Equal(t, 2, subarraySum([]int{1, 1, 1}, 2))
}

func Test_subarraySum2(t *testing.T) {
	assert.Equal(t, 2, subarraySum([]int{1, 2, 3}, 3))
}

func Test_subarraySum3(t *testing.T) {
	assert.Equal(t, 3, subarraySum([]int{10, -5, -3, -2, 1, -1}, 0))
}

func Test_subarraySum4(t *testing.T) {
	assert.Equal(t, 3, subarraySum([]int{1, 1, 1, 1}, 2))
}
