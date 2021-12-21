package task0231

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPowerOfTwo(t *testing.T) {
	assert.Equal(t, true, isPowerOfTwo(1))
	assert.Equal(t, true, isPowerOfTwo(16))
	assert.Equal(t, false, isPowerOfTwo(3))
}
