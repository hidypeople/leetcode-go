package task1015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestRepunitDivByK(t *testing.T) {
	assert.Equal(t, 1, smallestRepunitDivByK(1))
}

func Test_smallestRepunitDivByK2(t *testing.T) {
	assert.Equal(t, -1, smallestRepunitDivByK(2))
}

func Test_smallestRepunitDivByK3(t *testing.T) {
	assert.Equal(t, 3, smallestRepunitDivByK(3))
}

func Test_smallestRepunitDivByK4(t *testing.T) {
	assert.Equal(t, 22, smallestRepunitDivByK(23))
}
