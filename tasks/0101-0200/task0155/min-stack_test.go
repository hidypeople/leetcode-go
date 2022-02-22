package task0155

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack_Push(t *testing.T) {
	stack := Constructor()
	stack.Push(10)
	assert.Equal(t, 10, stack.GetMin())
	assert.Equal(t, 10, stack.Top())
	stack.Push(9)
	assert.Equal(t, 9, stack.GetMin())
	assert.Equal(t, 9, stack.Top())
	stack.Push(1)
	assert.Equal(t, 1, stack.GetMin())
	assert.Equal(t, 1, stack.Top())
	stack.Push(3)
	assert.Equal(t, 1, stack.GetMin())
	assert.Equal(t, 3, stack.Top())
	stack.Pop()
	assert.Equal(t, 1, stack.GetMin())
	assert.Equal(t, 1, stack.Top())
}
