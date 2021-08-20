package task0028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr_Fastest(t *testing.T) {
	assert.Equal(t, 0, strStr_Fastest("", ""))
}

func Test_strStr(t *testing.T) {
	assert.Equal(t, 2, strStr("hello", "ll"))
	assert.Equal(t, -1, strStr("aaaaa", "bba"))
	assert.Equal(t, 0, strStr("", ""))
	assert.Equal(t, -1, strStr("aaa", "aaaa"))
	assert.Equal(t, 4, strStr("mississippi", "issip"))
	assert.Equal(t, -1, strStr("mississippi", "issipi"))
	assert.Equal(t, 0, strStr("a", "a"))
}
