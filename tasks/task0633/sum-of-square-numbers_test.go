package task0633

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_judgeSquareSum(t *testing.T) {
	assert.False(t, judgeSquareSum(3))
}

func Test_judgeSquareSum2(t *testing.T) {
	assert.True(t, judgeSquareSum(4))
}

func Test_judgeSquareSum3(t *testing.T) {
	assert.True(t, judgeSquareSum(2))
}

func Test_judgeSquareSum4(t *testing.T) {
	assert.True(t, judgeSquareSum(5))
}

func Test_judgeSquareSum5(t *testing.T) {
	assert.True(t, judgeSquareSum(14645)) // 121^2 + 2^2
}

func Test_judgeSquareSum6(t *testing.T) {
	assert.False(t, judgeSquareSum(14644)) // 121^2 + 2^2 - 1
}

func Test_judgeSquareSum7(t *testing.T) {
	assert.False(t, judgeSquareSum(6))
}

func Test_judgeSquareSum8(t *testing.T) {
	assert.False(t, judgeSquareSum(14))
}

func Test_judgeSquareSum9(t *testing.T) {
	assert.False(t, judgeSquareSum(2147483642))
}
