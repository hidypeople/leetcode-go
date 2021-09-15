package task0070

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	assert.Equal(t, 2, climbStairs(2))
}

func Test_climbStairs2(t *testing.T) {
	assert.Equal(t, 3, climbStairs(3))
}

func Test_climbStairs3(t *testing.T) {
	assert.Equal(t, 8, climbStairs(5))
}

func Test_climbStairs4(t *testing.T) {
	assert.Equal(t, 1836311903, climbStairs(45))
}

func Test_climbStairs5(t *testing.T) {
	assert.Equal(t, 89, climbStairs(10))
}

func Test_climbStairs6(t *testing.T) {
	assert.Equal(t, 1346269, climbStairs(30))
}
