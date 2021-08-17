package task0206

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	var list *ListNode = nil
	assert.Nil(t, reverseList(list))
}

func TestReverseList2(t *testing.T) {
	var list *ListNode = LinkedListFromArray([]int{1})
	result := LinkedListToArray(reverseList(list))
	assert.Equal(t, result, []int{1})
}

func TestReverseList3(t *testing.T) {
	var list *ListNode = LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(reverseList(list))
	assert.Equal(t, result, []int{5, 4, 3, 2, 1})
}

func TestReverseList4(t *testing.T) {
	var list *ListNode = LinkedListFromArray([]int{1, 2, 3, 4, 5, 6})
	result := LinkedListToArray(reverseList(list))
	assert.Equal(t, result, []int{6, 5, 4, 3, 2, 1})
}
