package task0013

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	assert.Equal(t, 1, romanToInt("I"))
	assert.Equal(t, 3, romanToInt("III"))
	assert.Equal(t, 4, romanToInt("IV"))
	assert.Equal(t, 8, romanToInt("VIII"))
	assert.Equal(t, 9, romanToInt("IX"))
	assert.Equal(t, 11, romanToInt("XI"))
	assert.Equal(t, 22, romanToInt("XXII"))
	assert.Equal(t, 49, romanToInt("XLIX"))
	assert.Equal(t, 58, romanToInt("LVIII"))
	assert.Equal(t, 502, romanToInt("DII"))
	assert.Equal(t, 402, romanToInt("CDII"))
	assert.Equal(t, 202, romanToInt("CCII"))
	assert.Equal(t, 1994, romanToInt("MCMXCIV"))
}
