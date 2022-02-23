package task0201

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rangeBitwiseAnd(t *testing.T) {
	assert.Equal(t, 4, rangeBitwiseAnd(5, 7))
}

func Test_rangeBitwiseAnd2(t *testing.T) {
	assert.Equal(t, 0, rangeBitwiseAnd(0, 0))
}

func Test_rangeBitwiseAnd3(t *testing.T) {
	assert.Equal(t, 0, rangeBitwiseAnd(1, 2147483647))
}

func Test_rangeBitwiseAnd4(t *testing.T) {
	assert.Equal(t, 22, rangeBitwiseAnd(22, 22))
}

func Test_rangeBitwiseAnd5(t *testing.T) {
	assert.Equal(t, 6, rangeBitwiseAnd(6, 7))
}
