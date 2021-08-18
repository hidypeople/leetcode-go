package task0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	assert.True(t, isValid("()"))
	assert.True(t, isValid("()[]{}"))
	assert.True(t, isValid("{[]}"))
	assert.False(t, isValid("(]"))
	assert.False(t, isValid("()]]"))
	assert.False(t, isValid("(])"))
	assert.False(t, isValid("([)]"))
}
