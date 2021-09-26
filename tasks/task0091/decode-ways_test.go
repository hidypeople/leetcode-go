package task0091

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numDecodings(t *testing.T) {
	assert.Equal(t, 2, numDecodings("12"))
}

func Test_numDecodings2(t *testing.T) {
	assert.Equal(t, 3, numDecodings("226"))
}

func Test_numDecodings3(t *testing.T) {
	assert.Equal(t, 0, numDecodings("0"))
}

func Test_numDecodings4(t *testing.T) {
	assert.Equal(t, 0, numDecodings("06"))
}

func Test_numDecodings5(t *testing.T) {
	assert.Equal(t, 2, numDecodings("11106"))
}

func Test_numDecodings6(t *testing.T) {
	assert.Equal(t, 2584, numDecodings("12121212121212121"))
}

func Test_numDecodings7(t *testing.T) {
	assert.Equal(t, 3, numDecodings("1201234"))
}

func Test_numDecodings8(t *testing.T) {
	assert.Equal(t, 9, numDecodings("123123"))
}

func Test_numDecodings9(t *testing.T) {
	assert.Equal(t, 1, numDecodings("2839"))
}
