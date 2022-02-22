package task0125

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("A man, a plan, a canal: Panama"))
}

func Test_isPalindrome2(t *testing.T) {
	assert.False(t, isPalindrome("race a car"))
}

func Test_isPalindrome3(t *testing.T) {
	assert.False(t, isPalindrome("0P"))
}
