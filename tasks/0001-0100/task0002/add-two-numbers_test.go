package task0002

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	v1 := LinkedListFromArray([]int{1, 2, 3}) // 321
	v2 := LinkedListFromArray([]int{2, 3, 4}) // 432
	result := LinkedListToArray(addTwoNumbers(v1, v2))
	assert.Equal(t, result, []int{3, 5, 7})
}
