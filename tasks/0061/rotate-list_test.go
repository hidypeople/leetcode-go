package tasks

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	var l1 *ListNode = nil
	assert.Nil(t, RotateRight(l1, 1))
}

func TestRotateRight2(t *testing.T) {
	var l1 *ListNode = LinkedListFromArray([]int{1})
	assert.Equal(t, l1, RotateRight(l1, 100))
}

func TestRotateRight3(t *testing.T) {
	var l1 *ListNode = LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(RotateRight(l1, 1))
	assert.Equal(t, []int{5, 1, 2, 3, 4}, result)
}

func TestRotateRight4(t *testing.T) {
	var l1 *ListNode = LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(RotateRight(l1, 0))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

func TestRotateRight5(t *testing.T) {
	var l1 *ListNode = LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(RotateRight(l1, 5))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}

func TestRotateRight6(t *testing.T) {
	var l1 *ListNode = LinkedListFromArray([]int{1, 2})
	result := LinkedListToArray(RotateRight(l1, 1))
	assert.Equal(t, []int{2, 1}, result)
}

func TestRotateRight7(t *testing.T) {
	var l1 *ListNode = LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := LinkedListToArray(RotateRight(l1, 4))
	assert.Equal(t, []int{2, 3, 4, 5, 1}, result)
}
