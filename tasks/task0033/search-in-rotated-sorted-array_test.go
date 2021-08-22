package task0033

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search0(t *testing.T) {
	assert.Equal(t, 8, search([]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 2))
}

func Test_search(t *testing.T) {
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func Test_search2(t *testing.T) {
	assert.Equal(t, -1, search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}

func Test_search3(t *testing.T) {
	assert.Equal(t, -1, search([]int{1}, 0))
}

func Test_search4(t *testing.T) {
	assert.Equal(t, 0, search([]int{1}, 1))
}

func Test_search5(t *testing.T) {
	assert.Equal(t, 1, search([]int{3, 1}, 1))
}

func Test_search6(t *testing.T) {
	assert.Equal(t, 1, search([]int{1, 3}, 3))
}

func Test_search7(t *testing.T) {
	assert.Equal(t, 2, search([]int{1, 3, 5}, 5))
}

func Test_search8(t *testing.T) {
	assert.Equal(t, 0, search([]int{5, 1, 3}, 5))
}

func Test_search9(t *testing.T) {
	assert.Equal(t, 4, search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8))
}

func Test_search10(t *testing.T) {
	assert.Equal(t, -1, search([]int{1, 2, 3, 6, 7, 9}, 0))
}

// []

func Test_binarySearch(t *testing.T) {
	assert.Equal(t, -1, binarySearch([]int{}, 1))

	assert.Equal(t, 0, binarySearch([]int{1}, 1))
	assert.Equal(t, -1, binarySearch([]int{1}, 3))

	assert.Equal(t, -1, binarySearch([]int{1, 2}, 3))
	assert.Equal(t, 0, binarySearch([]int{1, 2}, 1))
	assert.Equal(t, 1, binarySearch([]int{1, 2}, 2))

	assert.Equal(t, -1, binarySearch([]int{1, 2, 4, 5}, 3))
	assert.Equal(t, -1, binarySearch([]int{1, 2, 4, 5}, 10))
	assert.Equal(t, -1, binarySearch([]int{1, 2, 4, 5}, -1))
	assert.Equal(t, 2, binarySearch([]int{1, 2, 4, 5}, 4))
	assert.Equal(t, 3, binarySearch([]int{1, 2, 4, 5}, 5))
	assert.Equal(t, 0, binarySearch([]int{1, 2, 4, 5}, 1))
}
