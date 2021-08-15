package task0234

import (
	. "leetcode/linkedList"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	list := LinkedListFromArray([]int{1})
	assert.True(t, IsPalindrome(list))
}

func TestIsPalindrome2(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2})
	assert.False(t, IsPalindrome(list))
}

func TestIsPalindrome3(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 2, 1})
	assert.True(t, IsPalindrome(list))
}

func TestIsPalindrome4(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 2, 1})
	assert.True(t, IsPalindrome(list))
}

func TestIsPalindrome5(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 4, 1})
	assert.False(t, IsPalindrome(list))
}

func TestIsPalindrome6(t *testing.T) {
	list := LinkedListFromArray([]int{1, 2, 3, 3, 2, 1})
	assert.True(t, IsPalindrome(list))
}

func TestIsPalindrome7(t *testing.T) {
	list := LinkedListFromArray([]int{0, 0})
	assert.True(t, IsPalindrome(list))
}

func TestIsPalindrome8(t *testing.T) {
	list := LinkedListFromArray([]int{1, 1, 2, 1})
	assert.False(t, IsPalindrome(list))
}
