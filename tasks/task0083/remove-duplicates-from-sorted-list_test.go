package task0083

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	assert.Nil(t, deleteDuplicates(nil))
}

func TestDeleteDuplicates3(t *testing.T) {
	list := LinkedListFromArray([]int{1, 1, 1, 1, 1, 3, 444, 4141, 4141, 5000})
	result := LinkedListToArray(deleteDuplicates(list))
	assert.Equal(t, []int{1, 3, 444, 4141, 5000}, result)
}
