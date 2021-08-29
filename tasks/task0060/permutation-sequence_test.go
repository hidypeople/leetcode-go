package task0060

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getPermutation(t *testing.T) {
	assert.Equal(t, "213", getPermutation(3, 3))
}

func Test_getPermutation2(t *testing.T) {
	assert.Equal(t, "2314", getPermutation(4, 9))
}

func Test_getPermutation3(t *testing.T) {
	assert.Equal(t, "123", getPermutation(3, 1))
}
