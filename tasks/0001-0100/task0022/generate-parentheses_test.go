package task0022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	assert.Equal(t, []string{"()"}, generateParenthesis(1))
}

func Test_generateParenthesis2(t *testing.T) {
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
}
