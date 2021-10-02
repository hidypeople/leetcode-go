package task0725

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_splitListToParts(t *testing.T) {
	head := LinkedListFromArray([]int{1, 2, 3})
	assert.Equal(t,
		[][]int{{1}, {2}, {3}, {}, {}},
		LinkedListsToArray(splitListToParts(head, 5)))
}

func Test_splitListToParts2(t *testing.T) {
	head := LinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t,
		[][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}},
		LinkedListsToArray(splitListToParts(head, 3)))
}
