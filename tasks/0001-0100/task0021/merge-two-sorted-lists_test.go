package task0021

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := LinkedListFromArray([]int{1, 2, 3, 4})
	l2 := LinkedListFromArray([]int{2, 3, 4, 5})
	result := LinkedListToArray(mergeTwoLists(l1, l2))
	assert.Equal(t, result, []int{1, 2, 2, 3, 3, 4, 4, 5})
}
