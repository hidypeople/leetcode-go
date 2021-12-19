package task0394

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decodeString(t *testing.T) {
	assert.Equal(t, "aaabcbc", decodeString("3[a]2[bc]"))
}

func Test_decodeString2(t *testing.T) {
	assert.Equal(t, "accaccacc", decodeString("3[a2[c]]"))
}

func Test_decodeString3(t *testing.T) {
	assert.Equal(t, "abcabccdcdcdef", decodeString("2[abc]3[cd]ef"))
}

func Test_decodeString4(t *testing.T) {
	assert.Equal(t, "abccdcdcdxyz", decodeString("abc3[cd]xyz"))
}

func Test_decodeString5(t *testing.T) {
	assert.Equal(t, "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode", decodeString("10[leetcode]"))
}

func Test_decodeString6(t *testing.T) {
	assert.Equal(t, "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode",
		decodeString("100[leetcode]"))
}
