package task0279

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSquares(t *testing.T) {
	assert.Equal(t, 3, numSquares(12))
}

func Test_numSquares2(t *testing.T) {
	assert.Equal(t, 2, numSquares(13))
}

func Test_numSquares3(t *testing.T) {
	assert.Equal(t, 1, numSquares(1))
}

func Test_numSquares4(t *testing.T) {
	assert.Equal(t, 1, numSquares(9))
}
