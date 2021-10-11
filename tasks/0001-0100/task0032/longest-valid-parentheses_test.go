package task0032

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestValidParentheses(t *testing.T) {
	assert.Equal(t, 2, longestValidParentheses("(()"))
}

func Test_longestValidParentheses2(t *testing.T) {
	assert.Equal(t, 4, longestValidParentheses(")()())"))
}

func Test_longestValidParentheses3(t *testing.T) {
	assert.Equal(t, 2, longestValidParentheses("()(()"))
}

func Test_longestValidParentheses4(t *testing.T) {
	assert.Equal(t, 2, longestValidParentheses("(((((()"))
}

func Test_longestValidParentheses5(t *testing.T) {
	assert.Equal(t, 2, longestValidParentheses(")))))((((((()"))
}

func Test_longestValidParentheses6(t *testing.T) {
	assert.Equal(t, 10, longestValidParentheses(")()(((())))("))
}

func Test_longestValidParentheses7(t *testing.T) {
	assert.Equal(t, 6, longestValidParentheses("(())()(()(("))
}
