package task0007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, -321, reverse(-123))
	assert.Equal(t, 21, reverse(120))
	assert.Equal(t, 0, reverse(0))
	assert.Equal(t, 0, reverse(-2147483648))
	assert.Equal(t, 0, reverse(2147483647))

}
