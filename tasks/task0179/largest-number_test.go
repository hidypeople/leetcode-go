package task0179

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestNumber(t *testing.T) {
	assert.Equal(t, "210", largestNumber([]int{10, 2}))
}

func Test_largestNumber2(t *testing.T) {
	assert.Equal(t, "9534330", largestNumber([]int{3, 30, 34, 5, 9}))
}

func Test_largestNumber3(t *testing.T) {
	assert.Equal(t, "990900", largestNumber([]int{9, 900, 90}))
}

func Test_largestNumber4(t *testing.T) {
	assert.Equal(t, "343234323", largestNumber([]int{34323, 3432}))
}

func Test_largestNumber5(t *testing.T) {
	assert.Equal(t, "9876543210", largestNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
