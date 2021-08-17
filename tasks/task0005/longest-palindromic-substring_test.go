package task0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, "aba", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
	assert.Equal(t, "a", longestPalindrome("a"))
	assert.Equal(t, "c", longestPalindrome("ac"))
}
