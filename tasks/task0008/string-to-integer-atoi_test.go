package task0008

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, -42, myAtoi("   -42"))
	assert.Equal(t, 4193, myAtoi("4193 with words"))
	assert.Equal(t, 0, myAtoi("words and 987"))
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))
	assert.Equal(t, 0, myAtoi("  --91283472332"))
	assert.Equal(t, -123, myAtoi("  -000123"))
	assert.Equal(t, 0, myAtoi("  -000a123"))
	assert.Equal(t, 2147483647, myAtoi("91283472332"))
}
