package task0152

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProduct(t *testing.T) {
	assert.Equal(t, 6, maxProduct([]int{2, 3, -2, 4}))
}

func Test_maxProduct2(t *testing.T) {
	assert.Equal(t, 0, maxProduct([]int{-2, 0, -1}))
}

func Test_maxProduct3(t *testing.T) {
	assert.Equal(t, 24, maxProduct([]int{2, 3, -1, 2, -1, 2, -2, 3}))
}

func Test_maxProduct4(t *testing.T) {
	assert.Equal(t, -2, maxProduct([]int{-2}))
}

func Test_maxProduct5(t *testing.T) {
	assert.Equal(t, 2, maxProduct([]int{0, 2}))
}

func Test_maxProduct6(t *testing.T) {
	assert.Equal(t, 4, maxProduct([]int{2, 2, -1, 2, 0, 0, -2, -2, -2, 0}))
}
