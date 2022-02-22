package task0190

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBits(t *testing.T) {
	assert.Equal(t, uint32(0b00111001011110000010100101000000),
		reverseBits(0b00000010100101000001111010011100))
}

func Test_reverseBits2(t *testing.T) {
	assert.Equal(t, uint32(0b00111001011110000010100101000000),
		reverseBits(0b00000010100101000001111010011100))
}

func Test_reverseBits3(t *testing.T) {
	assert.Equal(t, uint32(0b10111111111111111111111111111111),
		reverseBits(0b11111111111111111111111111111101))
}
