package task0328

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_oddEvenList(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5})
	result := oddEvenList(list)
	assert.Equal(t, []int{1, 3, 5, 2, 4}, LinkedListToArray(result))
}

func Test_oddEvenList2(t *testing.T) {
	list := LinkedListFromArray([]int{2, 1, 3, 5, 6, 4, 7})
	result := oddEvenList(list)
	assert.Equal(t, []int{2, 3, 6, 7, 1, 5, 4}, LinkedListToArray(result))
}

func Test_oddEvenList3(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 5, 6, 7, 8})
	result := oddEvenList(list)
	assert.Equal(t, []int{1, 3, 5, 7, 2, 4, 6, 8}, LinkedListToArray(result))
}
