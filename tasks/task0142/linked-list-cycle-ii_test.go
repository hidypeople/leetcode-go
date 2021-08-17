package task0142

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T) {
	l := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	l.Next.Next.Next = l.Next.Next
	assert.Equal(t, detectCycle(l), l.Next.Next)
}
