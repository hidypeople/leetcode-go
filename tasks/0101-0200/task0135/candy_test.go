package task0135

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_candy(t *testing.T) {
	assert.Equal(t, 5, candy([]int{1, 0, 2}))
}

func Test_candy2(t *testing.T) {
	assert.Equal(t, 4, candy([]int{1, 2, 2}))
}

func Test_candy3(t *testing.T) {
	assert.Equal(t, 11, candy([]int{10, 9, 10, 2, 1, 2}))
}

func Test_candy4(t *testing.T) {
	assert.Equal(t, 10, candy([]int{1, 2, 3, 4}))
	assert.Equal(t, 10, candy([]int{4, 3, 2, 1}))
}

func Test_candy5(t *testing.T) {
	assert.Equal(t, 6, candy([]int{4, 1, 1, 4}))
}

func Test_candy6(t *testing.T) {
	assert.Equal(t, 18, candy([]int{1, 6, 10, 8, 7, 3, 2}))
}

func Test_candy7(t *testing.T) {
	assert.Equal(t, 13, candy([]int{0, 1, 2, 3, 2, 1}))
}
