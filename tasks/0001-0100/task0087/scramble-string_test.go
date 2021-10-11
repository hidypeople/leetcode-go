package task0087

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isScramble(t *testing.T) {
	assert.True(t, isScramble("great", "rgeat"))
}

func Test_isScramble2(t *testing.T) {
	assert.False(t, isScramble("abcde", "caebd"))
}

func Test_isScramble3(t *testing.T) {
	assert.True(t, isScramble("a", "a"))
}

func Test_isScramble4(t *testing.T) {
	assert.True(t, isScramble("abc", "cba"))
}

func Test_isScramble5(t *testing.T) {
	assert.False(t, isScramble("eebaacbcbcadaaedceaaacadccd", "eadcaacabaddaceacbceaabeccd"))
}

func Test_isScramble6(t *testing.T) {
	assert.True(t, isScramble("aaabcdefga", "aagfedcbaa"))
}
