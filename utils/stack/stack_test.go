package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := StackFactory()
	assert.Equal(t, 0, stack.Len())
	_, ok1 := stack.Pop()
	assert.False(t, ok1)
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Len())
	x, ok := stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, x)
	x, ok = stack.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, x)
	assert.Equal(t, 0, stack.Len())
}
