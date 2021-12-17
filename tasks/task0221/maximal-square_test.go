package task0221

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximalSquare(t *testing.T) {
	assert.Equal(t, 4, maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
}

func Test_maximalSquare2(t *testing.T) {
	assert.Equal(t, 1, maximalSquare([][]byte{{'0', '1'}, {'1', '0'}}))
}

func Test_maximalSquare3(t *testing.T) {
	assert.Equal(t, 0, maximalSquare([][]byte{{'0'}}))
}
