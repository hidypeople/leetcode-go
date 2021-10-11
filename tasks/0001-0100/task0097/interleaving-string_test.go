package task0097

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isInterleave(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	assert.True(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave1(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbacc"
	assert.True(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave2(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	assert.False(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave3(t *testing.T) {
	s1 := ""
	s2 := ""
	s3 := ""
	assert.True(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave4(t *testing.T) {
	s1 := "a"
	s2 := ""
	s3 := "c"
	assert.False(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave5(t *testing.T) {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	assert.True(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave6(t *testing.T) {
	s1 := "aaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	s3 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	assert.False(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave7(t *testing.T) {
	s1 := "aa"
	s2 := "ab"
	s3 := "abaa"
	assert.True(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave8(t *testing.T) {
	s1 := "aabaac"
	s2 := "aadaaeaaf"
	s3 := "aadaaeaabaafaac"
	assert.True(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave9(t *testing.T) {
	s1 := "abababababababababababababababababababababababababababababababababababababababababababababababababbb"
	s2 := "babababababababababababababababababababababababababababababababababababababababababababababababaaaba"
	s3 := "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb"
	assert.False(t, isInterleave(s1, s2, s3))
}

func Test_isInterleave10(t *testing.T) {
	s1 := "bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa"
	s2 := "babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab"
	s3 := "babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab"
	assert.False(t, isInterleave(s1, s2, s3))
}
