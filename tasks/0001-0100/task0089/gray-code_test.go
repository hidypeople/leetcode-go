package task0089

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_grayCode(t *testing.T) {
	assert.Equal(t, []int{0, 1}, grayCode(1))
}

func Test_grayCode2(t *testing.T) {
	assert.Equal(t, []int{0, 1, 3, 2}, grayCode(2))
}

func Test_grayCode3(t *testing.T) {
	assert.Equal(t, []int{
		0b0000_0000_0000_0000,
		0b0000_0000_0000_0001,
		0b0000_0000_0000_0011,
		0b0000_0000_0000_0010,
		0b0000_0000_0000_0110,
		0b0000_0000_0000_0111,
		0b0000_0000_0000_0101,
		0b0000_0000_0000_0100,
		}, grayCode(3))
}