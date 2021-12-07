package task1290

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getDecimalValue(t *testing.T) {
	assert.Equal(t, 5, getDecimalValue(LinkedListFromArray([]int{1, 0, 1})))
	assert.Equal(t, 0, getDecimalValue(LinkedListFromArray([]int{0})))
	assert.Equal(t, 1, getDecimalValue(LinkedListFromArray([]int{1})))
	assert.Equal(t, 0, getDecimalValue(LinkedListFromArray([]int{0, 0, 0})))
	assert.Equal(t, 18880, getDecimalValue(LinkedListFromArray([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})))
}
