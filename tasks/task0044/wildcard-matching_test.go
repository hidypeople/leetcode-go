package task0040

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isMatch(t *testing.T) {
	assert.True(t, isMatch("aa", "*"))
}

func Test_isMatch2(t *testing.T) {
	assert.False(t, isMatch("cb", "?a"))
}

func Test_isMatch3(t *testing.T) {
	assert.True(t, isMatch("adceb", "*a*b"))
}

func Test_isMatch4(t *testing.T) {
	assert.False(t, isMatch("acdcb", "a*c?b"))
}

func Test_isMatch5(t *testing.T) {
	assert.False(t, isMatch("aa", "a"))
}

func Test_isMatch6(t *testing.T) {
	assert.True(t, isMatch("abcabczzzde", "*abc???de*"))
}

func Test_isMatch7(t *testing.T) {
	assert.True(t, isMatch("ho", "**ho"))
}

func Test_isMatch8(t *testing.T) {
	assert.False(t, isMatch("abbbabaaabbabbabbabaabbbaabaaaabbbabaaabbbbbaaababbb", "**a*b*aa***b***bbb*ba*a"))
}

func Test_isMatch9(t *testing.T) {
	assert.False(t, isMatch(
		"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
		"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))
}

func Test_isMatch10(t *testing.T) {
	assert.True(t, isMatch("aXXXXb", "aX*b"))
}
