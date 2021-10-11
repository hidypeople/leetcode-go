package task0010

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isMatch(t *testing.T) {
	assert.False(t, isMatch("aa", "a"))
	assert.False(t, isMatch("aa", "aaaa"))
	assert.False(t, isMatch("aa", "abab"))
	assert.True(t, isMatch("aa", "a*"))
	assert.True(t, isMatch("aab", "c*a*b"))
	assert.False(t, isMatch("aab", "c*a*b*c*d"))
	assert.False(t, isMatch("aab", "a*b*bb"))
	assert.False(t, isMatch("mississippi", "mis*is*p*."))
}
