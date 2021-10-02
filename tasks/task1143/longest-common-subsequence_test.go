package task1143

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonSubsequence(t *testing.T) {
	text1 := "abcde"
	text2 := "ace"
	assert.Equal(t, 3, longestCommonSubsequence(text1, text2))
}

func Test_longestCommonSubsequence2(t *testing.T) {
	text1 := "abc"
	text2 := "abc"
	assert.Equal(t, 3, longestCommonSubsequence(text1, text2))
}

func Test_longestCommonSubsequence3(t *testing.T) {
	text1 := "abc"
	text2 := "def"
	assert.Equal(t, 0, longestCommonSubsequence(text1, text2))
}

func Test_longestCommonSubsequence4(t *testing.T) {
	text1 := "oxcpqrsvwf"
	text2 := "shmtulqrypy"
	assert.Equal(t, 2, longestCommonSubsequence(text1, text2))
}

func Test_longestCommonSubsequence5(t *testing.T) {
	text1 := "aaa"
	text2 := "bbbbabbbabbb"
	assert.Equal(t, 2, longestCommonSubsequence(text1, text2))
	assert.Equal(t, 2, longestCommonSubsequence(text2, text1))
}

func Test_longestCommonSubsequence6(t *testing.T) {
	text1 := "xabcx"
	text2 := "yyybcyyyayyy"
	assert.Equal(t, 2, longestCommonSubsequence(text1, text2))
	assert.Equal(t, 2, longestCommonSubsequence(text2, text1))
}

func Test_longestCommonSubsequence7(t *testing.T) {
	text1 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	text2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	assert.Equal(t, 210, longestCommonSubsequence(text1, text2))
	assert.Equal(t, 210, longestCommonSubsequence(text2, text1))
}

func Test_longestCommonSubsequence8(t *testing.T) {
	text1 := "tufgfqlspqpidwrmjexifslkzobcjqvwlevghglynojchkvufqnwxixqnercbabm"
	text2 := "xuhadmbsbabqzirgrcxazcxypmjebgxyzmlepcdezwbsjkjalgdotcjavojehsvax"
	assert.Equal(t, 20, longestCommonSubsequence(text1, text2))
	assert.Equal(t, 20, longestCommonSubsequence(text2, text1))
}
