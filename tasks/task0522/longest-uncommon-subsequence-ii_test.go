package task0522

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLUSlength(t *testing.T) {
	assert.Equal(t, 3, findLUSlength([]string{"aba", "cdc", "eae"}))
}

func Test_findLUSlength2(t *testing.T) {
	assert.Equal(t, 15, findLUSlength([]string{"abcdef", "YYYbYYYdYYY", "XXXaXXXbXXXdXXX"}))
}

func Test_findLUSlength3(t *testing.T) {
	assert.Equal(t, -1, findLUSlength([]string{"aaa", "aaa", "aa"}))
}

func Test_isSubSequence(t *testing.T) {
	assert.True(t, isSubSequence("aaa", "aa"))
	assert.True(t, isSubSequence("aab", "aa"))
	assert.True(t, isSubSequence("baa", "aa"))
	assert.True(t, isSubSequence("aba", "aa"))
	assert.False(t, isSubSequence("xxxx", "zzzz"))
	assert.False(t, isSubSequence("abc", "abd"))
}
