package tasks

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	l := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	assert.False(t, HasCycle(l))
}

func TestHasCycle2(t *testing.T) {
	l := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	l.Next.Next.Next = l.Next.Next
	assert.True(t, HasCycle(l))
}
