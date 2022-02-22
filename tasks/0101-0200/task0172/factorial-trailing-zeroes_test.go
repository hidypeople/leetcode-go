package task0172

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trailingZeroes(t *testing.T) {
	assert.Equal(t, 0, trailingZeroes(3))
}

func Test_trailingZeroes2(t *testing.T) {
	assert.Equal(t, 1, trailingZeroes(5))
}

func Test_trailingZeroes3(t *testing.T) {
	assert.Equal(t, 0, trailingZeroes(0))
}

func Test_trailingZeroes4(t *testing.T) {
	assert.Equal(t, 4, trailingZeroes(20))
}

func Test_trailingZeroes5(t *testing.T) {
	assert.Equal(t, 6, trailingZeroes(25))
}
