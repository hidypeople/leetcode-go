package bitLifehacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Remove the smallest '1' bit (x >= 0):
// x       = 0x0011_1010_0001_0000
// x-1     = 0x0011_1010_0000_1111
// x&(x-1) = 0x0011_1010_0000_0000
//
// For x < 0: -((-x) & ((-x) - 1))
func Test_removeTheLastNonZeroBit(t *testing.T) {
	var x int

	x = 0b11010_1_0000
	assert.Equal(t, 0b11010_0_0000, x&(x-1))

	x = 0b0000001000
	assert.Equal(t, 0b0000000000, x&(x-1))

	x = 0b0000000000
	assert.Equal(t, 0b0000000000, x&(x-1))

	x = 0b0011_1010_0001_0000
	assert.Equal(t, 0b0011_1010_0000_0000, x&(x-1))

	x = -0b1111_0110
	assert.Equal(t, -0b1111_0100, -((-x) & ((-x) - 1)))
}

// Get the smallest '1' bit:
// x       = 0x0011_1010_0001_0000
// -x      = 0x1100_0101_1111_1111 (invert(0x0011_1010_0001_0000) + 1 and set the first bit to 1)
// x&(-x)  = 0x0000_0000_0001_0000
func Test_getTheLastNonZeroBit(t *testing.T) {
	var x int

	x = 0b11010_1_0000
	assert.Equal(t, 0b00000_1_0000, x&(-x))

	x = 0b0000001000
	assert.Equal(t, 0b0000001000, x&(-x))

	x = 0b0000000000
	assert.Equal(t, 0b0000000000, x&(-x))

	x = 0b0011_1010_0001_0000
	assert.Equal(t, 0b0000_0000_0001_0000, x&(-x))

	x = 0b1011_1010_0001_0000
	assert.Equal(t, 0b0000_0000_0001_0000, x&(-x))
}
