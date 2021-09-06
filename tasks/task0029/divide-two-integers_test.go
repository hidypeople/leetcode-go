package task0029

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divide(t *testing.T) {
	assert.Equal(t, 3, divide(10, 3))
}

func Test_divide2(t *testing.T) {
	assert.Equal(t, -3, divide(10, -3))
}

func Test_divide3(t *testing.T) {
	assert.Equal(t, -3, divide(-10, 3))
}

func Test_divide4(t *testing.T) {
	assert.Equal(t, 3, divide(-10, -3))
}

func Test_divide5(t *testing.T) {
	assert.Equal(t, 1, divide(1, 1))
}

func Test_divide6(t *testing.T) {
	assert.Equal(t, 0, divide(0, 110202))
}

func Test_divide7(t *testing.T) {
	MaxUint := ^uint32(0)
	MaxInt := int(MaxUint >> 1)
	assert.Equal(t, 1, divide(MaxInt, MaxInt))
	assert.Equal(t, 1, divide(MaxInt, MaxInt-1))
}

func Test_divide8(t *testing.T) {
	assert.Equal(t, 617, divide(12345, 20))
}

func Test_divide9(t *testing.T) {
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
}
