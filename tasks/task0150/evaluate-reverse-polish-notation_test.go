package task0150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evalRPN(t *testing.T) {
	assert.Equal(t, 9, evalRPN([]string{"2", "1", "+", "3", "*"}))
}

func Test_evalRPN2(t *testing.T) {
	assert.Equal(t, 6, evalRPN([]string{"4", "13", "5", "/", "+"}))
}

func Test_evalRPN3(t *testing.T) {
	assert.Equal(t, 22, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
